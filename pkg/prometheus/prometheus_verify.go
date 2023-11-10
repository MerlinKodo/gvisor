// Copyright 2022 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
	"unicode"

	pb "github.com/MerlinKodo/gvisor/pkg/metric/metric_go_proto"
)

const (
	// maxExportStaleness is the maximum allowed age of a snapshot when it is verified.
	// Used to avoid exporting snapshots from bogus times from ages past.
	maxExportStaleness = 10 * time.Second

	// MetaMetricPrefix is a prefix used for metrics defined by the metric server,
	// as opposed to metrics generated by each sandbox.
	// For this reason, this prefix is not allowed to be used in sandbox metrics.
	MetaMetricPrefix = "meta_"
)

// Prometheus process-level metric names and definitions.
// These are not necessarily exported, but we enforce that sandboxes may not
// export metrics sharing the same names.
// https://prometheus.io/docs/instrumenting/writing_clientlibs/#process-metrics
var (
	ProcessCPUSecondsTotal = Metric{
		Name: "process_cpu_seconds_total",
		Type: TypeGauge,
		Help: "Total user and system CPU time spent in seconds.",
	}
	ProcessOpenFDs = Metric{
		Name: "process_open_fds",
		Type: TypeGauge,
		Help: "Number of open file descriptors.",
	}
	ProcessMaxFDs = Metric{
		Name: "process_max_fds",
		Type: TypeGauge,
		Help: "Maximum number of open file descriptors.",
	}
	ProcessVirtualMemoryBytes = Metric{
		Name: "process_virtual_memory_bytes",
		Type: TypeGauge,
		Help: "Virtual memory size in bytes.",
	}
	ProcessVirtualMemoryMaxBytes = Metric{
		Name: "process_virtual_memory_max_bytes",
		Type: TypeGauge,
		Help: "Maximum amount of virtual memory available in bytes.",
	}
	ProcessResidentMemoryBytes = Metric{
		Name: "process_resident_memory_bytes",
		Type: TypeGauge,
		Help: "Resident memory size in bytes.",
	}
	ProcessHeapBytes = Metric{
		Name: "process_heap_bytes",
		Type: TypeGauge,
		Help: "Process heap size in bytes.",
	}
	ProcessStartTimeSeconds = Metric{
		Name: "process_start_time_seconds",
		Type: TypeGauge,
		Help: "Start time of the process since unix epoch in seconds.",
	}
	ProcessThreads = Metric{
		Name: "process_threads",
		Type: TypeGauge,
		Help: "Number of OS threads in the process.",
	}
)

// processMetrics is the set of process-level metrics.
var processMetrics = [9]*Metric{
	&ProcessCPUSecondsTotal,
	&ProcessOpenFDs,
	&ProcessMaxFDs,
	&ProcessVirtualMemoryBytes,
	&ProcessVirtualMemoryMaxBytes,
	&ProcessResidentMemoryBytes,
	&ProcessHeapBytes,
	&ProcessStartTimeSeconds,
	&ProcessThreads,
}

// internedStringMap allows for interning strings.
type internedStringMap map[string]*string

// Intern returns the interned version of the given string.
// If it is not already interned in the map, this function interns it.
func (m internedStringMap) Intern(s string) string {
	if existing, found := m[s]; found {
		return *existing
	}
	m[s] = &s
	return s
}

// globalInternMap is a string intern map used for globally-relevant data that repeats across
// verifiers, such as metric names and field names, but not field values or combinations of field
// values.
var (
	globalInternMu  sync.Mutex
	verifierCount   uint64
	globalInternMap = make(internedStringMap)
)

// globalIntern returns the interned version of the given string.
// If it is not already interned in the map, this function interns it.
func globalIntern(s string) string {
	globalInternMu.Lock()
	defer globalInternMu.Unlock()
	return globalInternMap.Intern(s)
}

func globalInternVerifierCreated() {
	globalInternMu.Lock()
	defer globalInternMu.Unlock()
	verifierCount++
}

func globalInternVerifierReleased() {
	globalInternMu.Lock()
	defer globalInternMu.Unlock()
	verifierCount--
	if verifierCount <= 0 {
		verifierCount = 0
		// No more verifiers active, so release the global map to not keep consuming needless resources.
		globalInternMap = make(internedStringMap)
	}
}

// numberPacker holds packedNumber data. It is useful to store large amounts of Number structs in a
// small memory footprint.
type numberPacker struct {
	// `data` *must* be pre-allocated if there is any number to be stored in it.
	// Attempts to pack a number that cannot fit into the existing space
	// allocated for this slice will cause a panic.
	// Callers may use `needsIndirection` to determine whether a number needs
	// space in this slice or not ahead of packing it.
	data []uint64
}

// packedNumber is a non-serializable but smaller-memory-footprint container for a numerical value.
// It can be unpacked out to a Number struct.
// This contains 4 bytes where we try to pack as much as possible.
// For the overhwelmingly-common case of integers that fit in 30 bits (i.e. 30 bits where the first
// 2 bits are zero, we store them directly here. Otherwise, we store the offset of a 64-bit number
// within numberPacker.
// Layout, going from highest to lowest bit:
// Bit 0 is the type: 0 for integer, 1 for float.
// Bit 1 is 0 if the number's value is stored within the next 30 bits, or 1 if the next 30 bits
// refer to an offset within numberPacker instead.
// In the case of a float, the next two bits (bits 2 and 3) may be used to encode a special value:
//   - 00 means not a special value
//   - 01 means NaN
//   - 10 means -infinity
//   - 11 means +infinity
//
// When not using a special value, the 32-bit exponent must fit in 5 bits, and is encoded using a
// bias of 2^4, meaning it ranges from -15 (encoded as 0b00000) to 16 (encoded as 0b11111), and an
// exponent of 0 is encoded as 0b01111.
// Floats that do not fit within this range must be encoded indirectly as float64s, similar to
// integers that don't fit in 30 bits.
type packedNumber uint32

// Useful masks and other bit-twiddling stuff for packedNumber.
const (
	typeField                = uint32(1 << 31)
	typeFieldInteger         = uint32(0)
	typeFieldFloat           = uint32(typeField)
	storageField             = uint32(1 << 30)
	storageFieldDirect       = uint32(0)
	storageFieldIndirect     = uint32(storageField)
	valueField               = uint32(1<<30 - 1)
	maxDirectUint            = uint64(valueField)
	float32ExponentField     = uint32(0x7f800000)
	float32ExponentShift     = uint32(23)
	float32ExponentBias      = uint32(127)
	float32FractionField     = uint32(0x7fffff)
	packedFloatExponentField = uint32(0x0f800000)
	packedFloatExponentBias  = uint32(15)
	packedFloatNaN           = packedNumber(typeFieldFloat | storageFieldDirect | 0x10000000)
	packedFloatNegInf        = packedNumber(typeFieldFloat | storageFieldDirect | 0x20000000)
	packedFloatInf           = packedNumber(typeFieldFloat | storageFieldDirect | 0x30000000)
)

// needsPackerStorage returns 0 for numbers that can be
// stored directly into the 32 bits of a packedNumber, or 1 for numbers that
// need more bits and would need to be stored into a numberPacker's `data`
// field.
//
//go:nosplit
func needsPackerStorage(n *Number) uint64 {
	if n.Float == 0.0 {
		v := n.Int
		if v >= 0 && v <= int64(valueField) {
			return 0
		}
		return 1
	}
	// n is a float.
	v := n.Float
	if math.IsNaN(v) || v == math.Inf(-1) || v == math.Inf(1) {
		return 0
	}
	if v >= 0.0 && float64(float32(v)) == v {
		float32Bits := math.Float32bits(float32(v))
		exponent := (float32Bits&float32ExponentField)>>float32ExponentShift - float32ExponentBias
		packedExponent := (exponent + packedFloatExponentBias) << float32ExponentShift
		if packedExponent&packedFloatExponentField == packedExponent {
			return 0
		}
	}
	return 1
}

// isIndirect returns 1 iff this packedNumber needs storage in a numberPacker.
//
//go:nosplit
func (n packedNumber) isIndirect() uint64 {
	if uint32(n)&storageField == storageFieldIndirect {
		return 1
	}
	return 0
}

// errOutOfPackerMemory is emitted when the number cannot be packed into a numberPacker.
var errOutOfPackerMemory = errors.New("out of numberPacker memory")

// pack packs a Number into a packedNumber.
//
//go:nosplit
func (p *numberPacker) pack(n *Number) packedNumber {
	if n.Float == 0.0 {
		v := n.Int
		if v >= 0 && v <= int64(maxDirectUint) {
			// We can store the integer value directly.
			return packedNumber(typeFieldInteger | storageFieldDirect | uint32(v))
		}
		if len(p.data) == cap(p.data) {
			panic(errOutOfPackerMemory)
		}
		p.data = append(p.data, uint64(v))
		return packedNumber(typeFieldInteger | storageFieldIndirect | uint32(len(p.data)-1))
	}
	// n is a float.
	v := n.Float
	if math.IsNaN(v) {
		return packedFloatNaN
	}
	if v == math.Inf(-1) {
		return packedFloatNegInf
	}
	if v == math.Inf(1) {
		return packedFloatInf
	}
	if v >= 0.0 && float64(float32(v)) == v {
		float32Bits := math.Float32bits(float32(v))
		exponent := (float32Bits&float32ExponentField)>>float32ExponentShift - float32ExponentBias
		packedExponent := (exponent + packedFloatExponentBias) << float32ExponentShift
		if packedExponent&packedFloatExponentField == packedExponent {
			float32Fraction := float32Bits & float32FractionField
			return packedNumber(typeFieldFloat | storageFieldDirect | packedExponent | float32Fraction)
		}
	}
	if len(p.data) == cap(p.data) {
		panic(errOutOfPackerMemory)
	}
	p.data = append(p.data, math.Float64bits(v))
	return packedNumber(typeFieldFloat | storageFieldIndirect | uint32(len(p.data)-1))
}

// packInt packs an integer.
//
//go:nosplit
func (p *numberPacker) packInt(val int64) packedNumber {
	n := Number{Int: val}
	return p.pack(&n)
}

// packFloat packs a floating-point number.
//
//go:nosplit
func (p *numberPacker) packFloat(val float64) packedNumber {
	n := Number{Float: val}
	return p.pack(&n)
}

// unpack unpacks a packedNumber back into a Number.
func (p *numberPacker) unpack(n packedNumber) *Number {
	switch uint32(n) & typeField {
	case typeFieldInteger:
		switch uint32(n) & storageField {
		case storageFieldDirect:
			return NewInt(int64(uint32(n) & valueField))
		case storageFieldIndirect:
			return NewInt(int64(p.data[uint32(n)&valueField]))
		}
	case typeFieldFloat:
		switch uint32(n) & storageField {
		case storageFieldDirect:
			switch n {
			case packedFloatNaN:
				return NewFloat(math.NaN())
			case packedFloatNegInf:
				return NewFloat(math.Inf(-1))
			case packedFloatInf:
				return NewFloat(math.Inf(1))
			default:
				exponent := ((uint32(n) & packedFloatExponentField) >> float32ExponentShift) - packedFloatExponentBias
				float32Bits := ((exponent + float32ExponentBias) << float32ExponentShift) | (uint32(n) & float32FractionField)
				return NewFloat(float64(math.Float32frombits(float32Bits)))
			}
		case storageFieldIndirect:
			return NewFloat(math.Float64frombits(p.data[uint32(n)&valueField]))
		}
	}
	panic("unreachable")
}

// mustUnpackInt unpacks an integer.
// It panics if the packedNumber is not an integer.
func (p *numberPacker) mustUnpackInt(n packedNumber) int64 {
	num := p.unpack(n)
	if !num.IsInteger() {
		panic("not an integer")
	}
	return num.Int
}

// mustUnpackFloat unpacks a floating-point number.
// It panics if the packedNumber is not an floating-point number.
func (p *numberPacker) mustUnpackFloat(n packedNumber) float64 {
	num := p.unpack(n)
	if *num == zero {
		return 0.0
	}
	if num.IsInteger() {
		panic("not a float")
	}
	return num.Float
}

// portTo ports over a packedNumber from this numberPacker to a new one.
// It is equivalent to `p.pack(other.unpack(n))` but avoids
// allocations in the overwhelmingly-common case where the number is direct.
func (p *numberPacker) portTo(other *numberPacker, n packedNumber) packedNumber {
	if uint32(n)&storageField == storageFieldDirect {
		// `n` is self-contained, just return as-is.
		return n
	}
	if len(other.data) == cap(other.data) {
		panic(errOutOfPackerMemory)
	}
	other.data = append(other.data, p.data[uint32(n)&valueField])
	return packedNumber(uint32(n)&(typeField|storageField) | uint32(len(other.data)-1))
}

// distributionSnapshot contains the data for a single field combination of a
// distribution ("histogram") metric.
type distributionSnapshot struct {
	// sum is the sum of all samples across all buckets.
	sum packedNumber

	// count is the number of samples across all buckets.
	count packedNumber

	// min is the lowest-recorded sample in the distribution.
	// It is only meaningful when count >= 1.
	min packedNumber

	// max is the highest-recorded sample in the distribution.
	// It is only meaningful when count >= 1.
	max packedNumber

	// ssd is the sum-of-squared-deviations computation of the distribution.
	// If non-zero, it is always a floating-point number.
	// It is only meaningful when count >= 2.
	ssd packedNumber

	// numSamples is the number of samples in each bucket.
	numSamples []packedNumber
}

// verifiableMetric verifies a single metric within a Verifier.
type verifiableMetric struct {
	metadata              *pb.MetricMetadata
	wantMetric            Metric
	numFields             uint32
	verifier              *Verifier
	allowedFieldValues    map[string]map[string]struct{}
	wantBucketUpperBounds []Number

	// The following fields are used to verify that values are actually increasing monotonically.
	// They are only read and modified when the parent Verifier.mu is held.
	// They are mapped by their combination of field values.

	// lastCounterValue is used for counter metrics.
	lastCounterValue map[string]packedNumber

	// lastDistributionSnapshot is used for distribution ("histogram") metrics.
	lastDistributionSnapshot map[string]*distributionSnapshot
}

// newVerifiableMetric creates a new verifiableMetric that can verify the
// values of a metric with the given metadata.
func newVerifiableMetric(metadata *pb.MetricMetadata, verifier *Verifier) (*verifiableMetric, error) {
	promName := metadata.GetPrometheusName()
	if metadata.GetName() == "" || promName == "" {
		return nil, errors.New("metric has no name")
	}
	for _, processMetric := range processMetrics {
		if promName == processMetric.Name {
			return nil, fmt.Errorf("metric name %q is reserved by Prometheus for process-level metrics", promName)
		}
	}
	if strings.HasPrefix(promName, MetaMetricPrefix) {
		return nil, fmt.Errorf("metric name %q starts with %q which is a reserved prefix", promName, "meta_")
	}
	if !unicode.IsLower(rune(promName[0])) {
		return nil, fmt.Errorf("invalid initial character in prometheus metric name: %q", promName)
	}
	for _, r := range promName {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) && r != '_' {
			return nil, fmt.Errorf("invalid character %c in prometheus metric name %q", r, promName)
		}
	}
	numFields := uint32(len(metadata.GetFields()))
	var allowedFieldValues map[string]map[string]struct{}
	if numFields > 0 {
		seenFields := make(map[string]struct{}, numFields)
		allowedFieldValues = make(map[string]map[string]struct{}, numFields)
		for _, field := range metadata.GetFields() {
			fieldName := field.GetFieldName()
			if _, alreadyExists := seenFields[fieldName]; alreadyExists {
				return nil, fmt.Errorf("field %s is defined twice", fieldName)
			}
			seenFields[fieldName] = struct{}{}
			if len(field.GetAllowedValues()) == 0 {
				return nil, fmt.Errorf("field %s has no allowed values", fieldName)
			}
			fieldValues := make(map[string]struct{}, len(field.GetAllowedValues()))
			for _, value := range field.GetAllowedValues() {
				if _, alreadyExists := fieldValues[value]; alreadyExists {
					return nil, fmt.Errorf("field %s has duplicate allowed value %q", fieldName, value)
				}
				fieldValues[globalIntern(value)] = struct{}{}
			}
			allowedFieldValues[globalIntern(fieldName)] = fieldValues
		}
	}
	v := &verifiableMetric{
		metadata: metadata,
		verifier: verifier,
		wantMetric: Metric{
			Name: globalIntern(promName),
			Help: globalIntern(metadata.GetDescription()),
		},
		numFields:          numFields,
		allowedFieldValues: allowedFieldValues,
	}
	numFieldCombinations := len(allowedFieldValues)
	switch metadata.GetType() {
	case pb.MetricMetadata_TYPE_UINT64:
		v.wantMetric.Type = TypeGauge
		if metadata.GetCumulative() {
			v.wantMetric.Type = TypeCounter
			v.lastCounterValue = make(map[string]packedNumber, numFieldCombinations)
		}
	case pb.MetricMetadata_TYPE_DISTRIBUTION:
		v.wantMetric.Type = TypeHistogram
		numBuckets := len(metadata.GetDistributionBucketLowerBounds()) + 1
		if numBuckets <= 1 || numBuckets > 256 {
			return nil, fmt.Errorf("unsupported number of buckets: %d", numBuckets)
		}
		v.wantBucketUpperBounds = make([]Number, numBuckets)
		for i, boundary := range metadata.GetDistributionBucketLowerBounds() {
			v.wantBucketUpperBounds[i] = Number{Int: boundary}
		}
		v.wantBucketUpperBounds[numBuckets-1] = Number{Float: math.Inf(1)}
		v.lastDistributionSnapshot = make(map[string]*distributionSnapshot, numFieldCombinations)
	default:
		return nil, fmt.Errorf("invalid type: %v", metadata.GetType())
	}
	return v, nil
}

func (v *verifiableMetric) numFieldCombinations() int {
	return len(v.allowedFieldValues)
}

// verify does read-only checks on `data`.
// `metricFieldsSeen` is passed across calls to `verify`. It is used to track the set of metric
// field values that have already been seen. `verify` should populate this.
// `dataToFieldsSeen` is passed across calls to `verify` and other methods of `verifiableMetric`.
// It is used to store the canonical representation of the field values seen for each *Data.
//
// Precondition: `Verifier.mu` is held.
func (v *verifiableMetric) verify(data *Data, metricFieldsSeen map[string]struct{}, dataToFieldsSeen map[*Data]string) error {
	if *data.Metric != v.wantMetric {
		return fmt.Errorf("invalid metric definition: got %+v want %+v", data.Metric, v.wantMetric)
	}

	// Verify fields.
	if uint32(len(data.Labels)) != v.numFields {
		return fmt.Errorf("invalid number of fields: got %d want %d", len(data.Labels), v.numFields)
	}
	var fieldValues strings.Builder
	firstField := true
	for _, field := range v.metadata.GetFields() {
		fieldName := field.GetFieldName()
		value, found := data.Labels[fieldName]
		if !found {
			return fmt.Errorf("did not specify field %q", fieldName)
		}
		if _, allowed := v.allowedFieldValues[fieldName][value]; !allowed {
			return fmt.Errorf("value %q is not allowed for field %s", value, fieldName)
		}
		if !firstField {
			fieldValues.WriteRune(',')
		}
		fieldValues.WriteString(value)
		firstField = false
	}
	fieldValuesStr := fieldValues.String()
	if _, alreadySeen := metricFieldsSeen[fieldValuesStr]; alreadySeen {
		return fmt.Errorf("combination of field values %q was already seen", fieldValuesStr)
	}

	// Verify value.
	gotNumber := data.Number != nil
	gotHistogram := data.HistogramValue != nil
	numSpecified := 0
	if gotNumber {
		numSpecified++
	}
	if gotHistogram {
		numSpecified++
	}
	if numSpecified != 1 {
		return fmt.Errorf("invalid number of value fields specified: %d", numSpecified)
	}
	switch v.metadata.GetType() {
	case pb.MetricMetadata_TYPE_UINT64:
		if !gotNumber {
			return errors.New("expected number value for gauge or counter")
		}
		if !data.Number.IsInteger() {
			return fmt.Errorf("integer metric got non-integer value: %v", data.Number)
		}
	case pb.MetricMetadata_TYPE_DISTRIBUTION:
		if !gotHistogram {
			return errors.New("expected histogram value for histogram")
		}
		if len(data.HistogramValue.Buckets) != len(v.wantBucketUpperBounds) {
			return fmt.Errorf("invalid number of buckets: got %d want %d", len(data.HistogramValue.Buckets), len(v.wantBucketUpperBounds))
		}
		if data.HistogramValue.SumOfSquaredDeviations.IsInteger() && data.HistogramValue.SumOfSquaredDeviations.Int != 0 {
			return fmt.Errorf("sum of squared deviations must be a floating-point value, got %v", data.HistogramValue.SumOfSquaredDeviations)
		}
		for i, b := range data.HistogramValue.Buckets {
			if want := v.wantBucketUpperBounds[i]; b.UpperBound != want {
				return fmt.Errorf("invalid upper bound for bucket %d (0-based): got %v want %v", i, b.UpperBound, want)
			}
		}
	default:
		return fmt.Errorf("invalid metric type: %v", v.wantMetric.Type)
	}

	// All passed. Update the maps that are shared across calls.
	fieldValuesStr = v.verifier.internMap.Intern(fieldValuesStr)
	dataToFieldsSeen[data] = fieldValuesStr
	metricFieldsSeen[fieldValuesStr] = struct{}{}
	return nil
}

// verifyIncrement verifies that incremental metrics are monotonically increasing.
//
// Preconditions: `verify` has succeeded on the given `data`, and `Verifier.mu` is held.
func (v *verifiableMetric) verifyIncrement(data *Data, fieldValues string, packer *numberPacker) error {
	switch v.wantMetric.Type {
	case TypeCounter:
		last := packer.unpack(v.lastCounterValue[v.verifier.internMap.Intern(fieldValues)])
		if !last.SameType(data.Number) {
			return fmt.Errorf("counter number type changed: %v vs %v", last, data.Number)
		}
		if last.GreaterThan(data.Number) {
			return fmt.Errorf("counter value decreased from %v to %v", last, data.Number)
		}
	case TypeHistogram:
		lastDistributionSnapshot := v.lastDistributionSnapshot[v.verifier.internMap.Intern(fieldValues)]
		if lastDistributionSnapshot == nil {
			lastDistributionSnapshot = &distributionSnapshot{
				numSamples: make([]packedNumber, len(v.wantBucketUpperBounds)),
			}
			v.lastDistributionSnapshot[v.verifier.internMap.Intern(fieldValues)] = lastDistributionSnapshot
		}
		lastCount := packer.mustUnpackInt(lastDistributionSnapshot.count)
		if lastCount >= 1 {
			lastMin := packer.unpack(lastDistributionSnapshot.min)
			if !lastMin.SameType(&data.HistogramValue.Min) {
				return fmt.Errorf("minimum value type changed: %v vs %v", lastMin, data.HistogramValue.Min)
			}
			if data.HistogramValue.Min.GreaterThan(lastMin) {
				return fmt.Errorf("minimum value strictly increased: from %v to %v", lastMin, data.HistogramValue.Min)
			}
			lastMax := packer.unpack(lastDistributionSnapshot.max)
			if !lastMax.SameType(&data.HistogramValue.Max) {
				return fmt.Errorf("maximum value type changed: %v vs %v", lastMax, data.HistogramValue.Max)
			}
			if lastMax.GreaterThan(&data.HistogramValue.Max) {
				return fmt.Errorf("maximum value strictly decreased: from %v to %v", lastMax, data.HistogramValue.Max)
			}
		}
		if lastCount >= 2 {
			// We already verified that the new data is a floating-point number
			// earlier, no need to double-check here.
			lastSSD := packer.mustUnpackFloat(lastDistributionSnapshot.ssd)
			if data.HistogramValue.SumOfSquaredDeviations.Float < lastSSD {
				return fmt.Errorf("sum of squared deviations decreased from %v to %v", lastSSD, data.HistogramValue.SumOfSquaredDeviations.Float)
			}
		}
		numSamples := lastDistributionSnapshot.numSamples
		for i, b := range data.HistogramValue.Buckets {
			if uint64(packer.mustUnpackInt(numSamples[i])) > b.Samples {
				return fmt.Errorf("number of samples in bucket %d (0-based) decreased from %d to %d", i, packer.mustUnpackInt(numSamples[i]), b.Samples)
			}
		}
	}
	return nil
}

// packerCapacityNeeded returns the `numberPacker` capacity to store `Data`.
func (v *verifiableMetric) packerCapacityNeededForData(data *Data, fieldValues string) uint64 {
	switch v.wantMetric.Type {
	case TypeCounter:
		return needsPackerStorage(data.Number)
	case TypeHistogram:
		var toPack uint64
		var totalSamples uint64
		var buf Number
		for _, b := range data.HistogramValue.Buckets {
			buf = Number{Int: int64(b.Samples)}
			toPack += needsPackerStorage(&buf)
			totalSamples += b.Samples
		}
		toPack += needsPackerStorage(&data.HistogramValue.Total)
		toPack += needsPackerStorage(&data.HistogramValue.Min)
		toPack += needsPackerStorage(&data.HistogramValue.Max)
		toPack += needsPackerStorage(&data.HistogramValue.SumOfSquaredDeviations)
		buf = Number{Int: int64(totalSamples)}
		toPack += needsPackerStorage(&buf)
		return toPack
	default:
		return 0
	}
}

// packerCapacityNeededForLast returns the `numberPacker` capacity needed to
// store the last snapshot's data that was not seen in the current snapshot
// (aka not in metricFieldsSeen).
func (v *verifiableMetric) packerCapacityNeededForLast(metricFieldsSeen map[string]struct{}) uint64 {
	var capacity uint64
	switch v.wantMetric.Type {
	case TypeCounter:
		for fieldValues, lastCounterValue := range v.lastCounterValue {
			if _, found := metricFieldsSeen[fieldValues]; found {
				continue
			}
			capacity += lastCounterValue.isIndirect()
		}
	case TypeHistogram:
		for fieldValues, distributionSnapshot := range v.lastDistributionSnapshot {
			if _, found := metricFieldsSeen[fieldValues]; found {
				continue
			}
			for _, b := range distributionSnapshot.numSamples {
				capacity += b.isIndirect()
			}
			capacity += distributionSnapshot.sum.isIndirect()
			capacity += distributionSnapshot.count.isIndirect()
			capacity += distributionSnapshot.min.isIndirect()
			capacity += distributionSnapshot.max.isIndirect()
			capacity += distributionSnapshot.ssd.isIndirect()
		}
	}
	return capacity
}

// update updates incremental metrics' "last seen" data.
//
// Preconditions: `verifyIncrement` has succeeded on the given `data`, `Verifier.mu` is held,
// and `packer` is guaranteed to have enough room to store all numbers.
func (v *verifiableMetric) update(data *Data, fieldValues string, packer *numberPacker) {
	switch v.wantMetric.Type {
	case TypeCounter:
		v.lastCounterValue[v.verifier.internMap.Intern(fieldValues)] = packer.pack(data.Number)
	case TypeHistogram:
		lastDistributionSnapshot := v.lastDistributionSnapshot[v.verifier.internMap.Intern(fieldValues)]
		lastBucketSamples := lastDistributionSnapshot.numSamples
		var count uint64
		for i, b := range data.HistogramValue.Buckets {
			lastBucketSamples[i] = packer.packInt(int64(b.Samples))
			count += b.Samples
		}
		lastDistributionSnapshot.sum = packer.pack(&data.HistogramValue.Total)
		lastDistributionSnapshot.count = packer.packInt(int64(count))
		lastDistributionSnapshot.min = packer.pack(&data.HistogramValue.Min)
		lastDistributionSnapshot.max = packer.pack(&data.HistogramValue.Max)
		lastDistributionSnapshot.ssd = packer.pack(&data.HistogramValue.SumOfSquaredDeviations)
	}
}

// repackUnseen packs all numbers that must be carried over from snapshot to snapshot and which were
// not seen in the latest snapshot's data.
// This function should carry over all numbers typically packed in `v.update` but for all metric
// field combinations that are not in `metricFieldsSeen`.
//
// Preconditions: `verifyIncrement` has succeeded on the given `data`,
// and `newPacker` is guaranteed to have enough room to store all numbers.
func (v *verifiableMetric) repackUnseen(metricFieldsSeen map[string]struct{}, oldPacker, newPacker *numberPacker) {
	switch v.wantMetric.Type {
	case TypeCounter:
		for fieldValues, lastCounterValue := range v.lastCounterValue {
			if _, found := metricFieldsSeen[fieldValues]; found {
				continue
			}
			v.lastCounterValue[fieldValues] = oldPacker.portTo(newPacker, lastCounterValue)
		}
	case TypeHistogram:
		for fieldValues, lastDistributionSnapshot := range v.lastDistributionSnapshot {
			if _, found := metricFieldsSeen[fieldValues]; found {
				continue
			}
			lastBucketSamples := lastDistributionSnapshot.numSamples
			for i, b := range lastBucketSamples {
				lastBucketSamples[i] = oldPacker.portTo(newPacker, b)
			}
			lastDistributionSnapshot.sum = oldPacker.portTo(newPacker, lastDistributionSnapshot.sum)
			lastDistributionSnapshot.count = oldPacker.portTo(newPacker, lastDistributionSnapshot.count)
			lastDistributionSnapshot.min = oldPacker.portTo(newPacker, lastDistributionSnapshot.min)
			lastDistributionSnapshot.max = oldPacker.portTo(newPacker, lastDistributionSnapshot.max)
			lastDistributionSnapshot.ssd = oldPacker.portTo(newPacker, lastDistributionSnapshot.ssd)
		}
	}
}

// Verifier allows verifying metric snapshot against metric registration data.
// The aim is to prevent a compromised Sentry from emitting bogus data or DoS'ing metric ingestion.
// A single Verifier should be used per sandbox. It is expected to be reused across exports such
// that it can enforce the export snapshot timestamp is strictly monotonically increasing.
type Verifier struct {
	knownMetrics map[string]*verifiableMetric

	// mu protects the fields below.
	mu sync.Mutex

	// internMap is used to intern strings relevant to this verifier only.
	// Globally-relevant strings should be interned in globalInternMap.
	internMap internedStringMap

	// lastPacker is a reference to the numberPacker used to pack numbers in the last successful
	// verification round.
	lastPacker *numberPacker

	// lastTimestamp is the snapshot timestamp of the last successfully-verified snapshot.
	lastTimestamp time.Time
}

// NewVerifier returns a new metric verifier that can verify the integrity of snapshots against
// the given metric registration data.
// It returns a cleanup function that must be called when the Verifier is no longer needed.
func NewVerifier(registration *pb.MetricRegistration) (*Verifier, func(), error) {
	globalInternVerifierCreated()
	verifier := &Verifier{
		knownMetrics: make(map[string]*verifiableMetric),
		internMap:    make(internedStringMap),
	}
	for _, metric := range registration.GetMetrics() {
		metricName := metric.GetPrometheusName()
		if _, alreadyExists := verifier.knownMetrics[metricName]; alreadyExists {
			globalInternVerifierReleased()
			return nil, func() {}, fmt.Errorf("metric %q registered twice", metricName)
		}
		verifiableM, err := newVerifiableMetric(metric, verifier)
		if err != nil {
			globalInternVerifierReleased()
			return nil, func() {}, fmt.Errorf("metric %q: %v", metricName, err)
		}
		verifier.knownMetrics[globalIntern(metricName)] = verifiableM
	}
	return verifier, globalInternVerifierReleased, nil
}

// Verify verifies the integrity of a snapshot against the metric registration data of the Verifier.
// It assumes that it will be called on snapshots obtained chronologically over time.
func (v *Verifier) Verify(snapshot *Snapshot) error {
	var err error

	// Basic timestamp checks.
	now := timeNow()
	if snapshot.When.After(now) {
		return errors.New("snapshot is from the future")
	}
	if snapshot.When.Before(now.Add(-maxExportStaleness)) {
		return fmt.Errorf("snapshot is too old; it is from %v, expected at least %v (%v from now)", snapshot.When, now.Add(-maxExportStaleness), maxExportStaleness)
	}

	// Start critical section.
	v.mu.Lock()
	defer v.mu.Unlock()

	// Metrics checks.
	fieldsSeen := make(map[string]map[string]struct{}, len(v.knownMetrics))
	dataToFieldsSeen := make(map[*Data]string, len(snapshot.Data))
	for _, data := range snapshot.Data {
		metricName := data.Metric.Name
		verifiableM, found := v.knownMetrics[metricName]
		if !found {
			return fmt.Errorf("snapshot contains unknown metric %q", metricName)
		}
		metricName = globalIntern(metricName)
		metricFieldsSeen, found := fieldsSeen[metricName]
		if !found {
			metricFieldsSeen = make(map[string]struct{}, verifiableM.numFieldCombinations())
			fieldsSeen[metricName] = metricFieldsSeen
		}
		if err = verifiableM.verify(data, metricFieldsSeen, dataToFieldsSeen); err != nil {
			return fmt.Errorf("metric %q: %v", metricName, err)
		}
	}

	if v.lastTimestamp.After(snapshot.When) {
		return fmt.Errorf("consecutive snapshots are not chronologically ordered: last verified snapshot was exported at %v, this one is from %v", v.lastTimestamp, snapshot.When)
	}

	for _, data := range snapshot.Data {
		if err := v.knownMetrics[data.Metric.Name].verifyIncrement(data, dataToFieldsSeen[data], v.lastPacker); err != nil {
			return fmt.Errorf("metric %q: %v", data.Metric.Name, err)
		}
	}
	var neededPackerCapacity uint64
	for _, data := range snapshot.Data {
		neededPackerCapacity += v.knownMetrics[data.Metric.Name].packerCapacityNeededForData(data, dataToFieldsSeen[data])
	}
	for name, metric := range v.knownMetrics {
		neededPackerCapacity += metric.packerCapacityNeededForLast(fieldsSeen[name])
	}
	if neededPackerCapacity > uint64(valueField) {
		return fmt.Errorf("snapshot contains too many large numbers to fit into packer memory (%d numbers needing indirection)", neededPackerCapacity)
	}

	// All checks succeeded, update last-seen data.
	// We need to be guaranteed to not fail past this point in the function.
	newPacker := &numberPacker{}
	if neededPackerCapacity != 0 {
		newPacker.data = make([]uint64, 0, neededPackerCapacity)
	}
	v.lastTimestamp = snapshot.When
	for _, data := range snapshot.Data {
		v.knownMetrics[globalIntern(data.Metric.Name)].update(data, v.internMap.Intern(dataToFieldsSeen[data]), newPacker)
	}
	if uint64(len(newPacker.data)) != neededPackerCapacity {
		for name, metric := range v.knownMetrics {
			metric.repackUnseen(fieldsSeen[name], v.lastPacker, newPacker)
		}
	}
	if uint64(len(newPacker.data)) != neededPackerCapacity {
		// We panic here because this represents an internal logic error,
		// not something the user did wrong.
		panic(fmt.Sprintf("did not pack the expected number of numbers in numberPacker: packed %d, expected %d; this indicates a logic error in verifyIncrement", len(newPacker.data), neededPackerCapacity))
	}
	v.lastPacker = newPacker
	return nil
}

// AllMetrics returns the metadata of all the metrics that were declared as
// part of this Verifier.
func (v *Verifier) AllMetrics() []*pb.MetricMetadata {
	metrics := make([]*pb.MetricMetadata, 0, len(v.knownMetrics))
	for _, m := range v.knownMetrics {
		metrics = append(metrics, m.metadata)
	}
	return metrics
}
