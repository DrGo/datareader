package datareader

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
)

// A Series is a homogeneously-typed one-dimensional sequence of data
// values, with an optional mask for missing values.
type Series struct {

	// A name describing what is in this series.
	Name string

	// The length of the series.
	length int

	// The data, must be a homogeneous array, e.g. []float64.
	data interface{}

	// Indicators that data values are missing.  If nil, there are
	// no missing values.
	missing []bool
}

// NewSeries returns a new Series object with the given name and data
// contents.  The data parameter must be an array of floats, ints, or
// strings.
func NewSeries(name string, data interface{}, missing []bool) (*Series, error) {

	var length int

	switch data.(type) {
	default:
		return nil, fmt.Errorf("Unknown data type in NewSeries")
	case []float64:
		length = len(data.([]float64))
	case []string:
		length = len(data.([]string))
	case []int64:
		length = len(data.([]int64))
	case []int32:
		length = len(data.([]int32))
	case []float32:
		length = len(data.([]float32))
	case []int16:
		length = len(data.([]int16))
	case []int8:
		length = len(data.([]int8))
	case []uint64:
		length = len(data.([]uint64))
	case []time.Time:
		length = len(data.([]time.Time))
	}

	ser := Series{
		Name:    name,
		length:  length,
		data:    data,
		missing: missing}

	return &ser, nil
}

// Write writes the entire Series to the given writer.
func (ser *Series) Write(w io.Writer) {
	ser.WriteRange(w, 0, ser.length)
}

// WriteRange writes the given subinterval of the Series to the given writer.
func (ser *Series) WriteRange(w io.Writer, first, last int) {

	io.WriteString(w, fmt.Sprintf("Name: %s\n", ser.Name))
	ty := fmt.Sprintf("%T", ser.data)
	io.WriteString(w, fmt.Sprintf("Type: %s\n", ty[2:len(ty)]))

	switch ser.data.(type) {
	default:
		panic("Unknown type in WriteRange")
	case []float64:
		data := ser.data.([]float64)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []float32:
		data := ser.data.([]float32)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []int64:
		data := ser.data.([]int64)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []int32:
		data := ser.data.([]int32)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []int16:
		data := ser.data.([]int16)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []int8:
		data := ser.data.([]int8)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []uint64:
		data := ser.data.([]uint64)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []string:
		data := ser.data.([]string)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	case []time.Time:
		data := ser.data.([]time.Time)
		for j := first; j < last; j++ {
			if (ser.missing == nil) || !ser.missing[j] {
				io.WriteString(w, fmt.Sprintf("%d:  %v\n", j, data[j]))
			} else {
				io.WriteString(w, fmt.Sprintf("%d:\n", j))
			}
		}
	}
}

// Print prints the entire Series to the standard output.
func (ser *Series) Print() {
	ser.Write(os.Stdout)
}

// PrintRange printes a slice of the Series to the standard output.
func (ser *Series) PrintRange(first, last int) {
	ser.WriteRange(os.Stdout, first, last)
}

// Data returns the data component of the Series.
func (ser *Series) Data() interface{} {
	return ser.data
}

// Missing returns the array of missing value indicators.
func (ser *Series) Missing() []bool {
	return ser.missing
}

// Length returns the number of elements in a Series.
func (ser *Series) Length() int {
	return ser.length
}

// AllClose returns true, 0 if the Series is within tol of the other
// series.  If the Series have different lengths, AllClose returns
// false, -1.  If the Series have different types, AllClose returns
// false, -2.  If the Series have the same type and the same length
// but are not equal, AllClose returns false, j, where j is the index
// of the first position where the two series differ.
func (ser *Series) AllClose(other *Series, tol float64) (bool, int) {

	if ser.length != other.length {
		return false, -1
	}

	if (ser.missing != nil) && (other.missing != nil) {
		for j := 0; j < ser.length; j++ {
			if ser.missing[j] != other.missing[j] {
				return false, j
			}
		}
	}

	// Utility function for missing mask
	cmiss := func(j int) int {
		f1 := (ser.missing == nil) || (ser.missing[j] == false)
		f2 := (other.missing == nil) || (other.missing[j] == false)
		if f1 != f2 {
			return 0 // inconsistent
		} else if f1 {
			return 1 // both non-missing
		} else {
			return 2 // both missing
		}
	}

	switch ser.data.(type) {
	default:
		panic(fmt.Sprintf("Unknown type %T in Series.AllClose", ser.data))
	case []float64:
		u := ser.data.([]float64)
		v, ok := other.data.([]float64)
		if !ok {
			return false, -2
		}
		for i := 0; i < ser.length; i++ {
			c := cmiss(i)
			if c == 0 {
				return false, i
			}
			if (c == 1) && (math.Abs(u[i]-v[i]) > tol) {
				return false, i
			}
		}
	case []float32:
		u := ser.data.([]float32)
		v, ok := other.data.([]float32)
		if !ok {
			return false, -2
		}
		for i := 0; i < ser.length; i++ {
			c := cmiss(i)
			if c == 0 {
				return false, i
			}
			if (c == 1) && (math.Abs(float64(u[i]-v[i])) > tol) {
				return false, i
			}
		}
	case []int64:
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && (ser.data.([]int64)[j] != other.data.([]int64)[j]) {
				return false, j
			}
		}
	case []int32:
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && (ser.data.([]int32)[j] != other.data.([]int32)[j]) {
				return false, j
			}
		}
	case []int16:
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && (ser.data.([]int16)[j] != other.data.([]int16)[j]) {
				return false, j
			}
		}
	case []int8:
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && (ser.data.([]int8)[j] != other.data.([]int8)[j]) {
				return false, j
			}
		}
	case []uint64:
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && (ser.data.([]uint64)[j] != other.data.([]uint64)[j]) {
				return false, j
			}
		}
	case []string:
		u := ser.data.([]string)
		v, ok := other.data.([]string)
		if !ok {
			return false, -2
		}
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && (u[j] != v[j]) {
				return false, j
			}
		}
	case []time.Time:
		u := ser.data.([]time.Time)
		v, ok := other.data.([]time.Time)
		if !ok {
			return false, -2
		}
		for j := 0; j < ser.length; j++ {
			c := cmiss(j)
			if c == 0 {
				return false, j
			}
			if (c == 1) && !u[j].Equal(v[j]) {
				return false, j
			}
		}
	}
	return true, 0
}

// AllEqual is equivalent to AllClose with tol=0.
func (ser *Series) AllEqual(other *Series) (bool, int) {
	return ser.AllClose(other, 0.0)
}

// UpcastNumeric converts in-place all numeric type variables to
// float64 values.  Non-numeric data is not affected.
func (ser *Series) UpcastNumeric() *Series {

	n := ser.Length()
	cmiss := ser.missing
	if cmiss != nil {
		cmiss = make([]bool, n)
		copy(cmiss, ser.missing)
	}

	switch ser.data.(type) {

	default:
		panic(fmt.Sprintf("unknown data type: %T\n", ser.data))
	case []float64:
		return ser
	case []string:
		return ser
	case []time.Time:
		return ser
	case []float32:
		d := ser.data.([]float32)
		n := len(d)
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = float64(d[i])
		}
		s, _ := NewSeries(ser.Name, a, cmiss)
		return s
	case []int64:
		d := ser.data.([]int64)
		n := len(d)
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = float64(d[i])
		}
		s, _ := NewSeries(ser.Name, a, cmiss)
		return s
	case []int32:
		d := ser.data.([]int32)
		n := len(d)
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = float64(d[i])
		}
		ser.data = a
		s, _ := NewSeries(ser.Name, a, cmiss)
		return s
	case []int16:
		d := ser.data.([]int16)
		n := len(d)
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = float64(d[i])
		}
		ser.data = a
		s, _ := NewSeries(ser.Name, a, cmiss)
		return s
	case []int8:
		d := ser.data.([]int8)
		n := len(d)
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = float64(d[i])
		}
		ser.data = a
		s, _ := NewSeries(ser.Name, a, cmiss)
		return s
	}
}

// ForceNumeric converts string values to float64 values, creating
// missing values where the conversion is not possible.  If the data
// is not string type, it is unaffected.
func (ser *Series) ForceNumeric() *Series {

	n := ser.length
	cmiss := make([]bool, n)
	if ser.missing != nil {
		copy(cmiss, ser.missing)
	}

	switch ser.data.(type) {
	default:
		return ser
	case []string:
		x := make([]float64, n)
		y := ser.data.([]string)
		for i := 0; i < n; i++ {
			if !cmiss[i] {
				v, err := strconv.ParseFloat(y[i], 64)
				if err != nil {
					cmiss[i] = true
				} else {
					x[i] = v
				}
			}
		}
		s, _ := NewSeries(ser.Name, x, cmiss)
		return s
	}
}

func (ser *Series) CountMissing() int {

	m := 0
	for i := 0; i < ser.length; i++ {
		if ser.missing[i] {
			m++
		}
	}

	return m
}

func (ser *Series) StringFunc(f func(string) string) *Series {

	n := ser.length
	cmiss := make([]bool, n)
	if ser.missing != nil {
		copy(cmiss, ser.missing)
	}

	switch ser.data.(type) {
	default:
		return ser
	case []string:
		x := ser.data.([]string)
		y := make([]string, n)
		for i, v := range x {
			y[i] = f(v)
		}
		s, _ := NewSeries(ser.Name, y, cmiss)
		return s
	}
}

func (ser *Series) ToString() *Series {

	n := ser.length
	cmiss := make([]bool, n)
	if ser.missing != nil {
		copy(cmiss, ser.missing)
	}

	switch ser.data.(type) {
	default:
		panic(fmt.Sprintf("unknown data type %T in ToString", ser.data))
	case []string:
		return ser
	case []float64:
		x := make([]string, n)
		y := ser.data.([]float64)
		for i := 0; i < n; i++ {
			if !cmiss[i] {
				x[i] = fmt.Sprintf("%v", y[i])
			}
		}
		s, _ := NewSeries(ser.Name, x, cmiss)
		return s
	}
}

func (ser *Series) NullStringMissing() *Series {

	n := ser.length
	cmiss := make([]bool, n)
	if ser.missing != nil {
		copy(cmiss, ser.missing)
	}

	switch ser.data.(type) {
	default:
		return ser
	case []string:
		x := make([]string, n)
		y := ser.data.([]string)
		copy(x, y)
		for i := 0; i < n; i++ {
			if len(x[i]) == 0 {
				cmiss[i] = true
			}
		}
		s, _ := NewSeries(ser.Name, x, cmiss)
		return s
	}
}

// SeriesArray is an array of pointers to Series objects.  It can represent
// a dataset consisting of several variables.
type SeriesArray []*Series

// AllClose returns (true, 0, 0) if all numeric values in
// corresponding columns of the two arrays of Series objects are
// within the given tolerance.  If any corresponding columns are not
// identically equal, returns (false, j, i), where j is the index of a
// column and i is the index of a row where the two Series are not
// identical.  If the two SeriesArray objects have different numbers
// of columns, returns (false, -1, -1).  If column j of the two
// SeriesArray objects have different lengths, returns (false, j, -1).
// If column j of the two SeriesArray objects have different types,
// returns (false, j, -2)
func (ser SeriesArray) AllClose(other []*Series, tol float64) (bool, int, int) {

	if len(ser) != len(other) {
		return false, -1, -1
	}

	for j := 0; j < len(ser); j++ {
		f, i := ser[j].AllClose(other[j], tol)
		if !f {
			return false, j, i
		}
	}

	return true, 0, 0
}

// AllEqual is equivalent to AllClose with tol = 0.
func (ser SeriesArray) AllEqual(other []*Series) (bool, int, int) {
	return ser.AllClose(other, 0.0)
}

func (ser *Series) Date_from_duration(base time.Time, units string) (*Series, error) {

	n := ser.Length()

	var miss []bool
	if ser.missing != nil {
		miss = make([]bool, n)
		copy(miss, ser.missing)
	}

	td, err := upcast_numeric(ser.data)
	if err != nil {
		return nil, err
	}

	newdate := make([]time.Time, n)
	for i := 0; i < n; i++ {
		switch units {
		default:
			return nil, fmt.Errorf("unknown time unit duration")
		case "days":
			if (miss == nil) || !miss[i] {
				newdate[i] = base.Add(time.Hour * time.Duration(24*td[i]))
			}
		}
	}

	rslt, err := NewSeries(ser.Name, newdate, miss)
	if err != nil {
		return nil, err
	}
	return rslt, nil
}
