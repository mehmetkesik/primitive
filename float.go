package primitive

import (
	"math"
	"strconv"
)

type Float float64

func (self *Float) ToString() string {
	return strconv.FormatFloat(float64(*self), 'E', -1, 64)
}

func (self *Float) ToInt() Int {
	return Int(*self)
}

func (self *Float) Abs() Float {
	return Float(math.Abs(float64(*self)))
}

func (self *Float) Pow(y Float) Float {
	return Float(math.Pow(float64(*self), float64(y)))
}

type FloatSlice []Float

func (self *FloatSlice) Get(index Int) Float {
	return (*self)[index]
}

func (self *FloatSlice) Len() Int {
	return Int(len(*self))
}

func (self *FloatSlice) Append(i Float) {
	*self = append(*self, i)
}

func (self *FloatSlice) Delete(index Int) {
	*self = append((*self)[:index], (*self)[index+1:]...)
}

func (self *FloatSlice) Map(mapping func(Float) Float) {
	for i, k := range *self {
		(*self)[i] = mapping(k)
	}
}

func (self *FloatSlice) Filter(mapping func(Float) bool) {
	var v FloatSlice
	for _, k := range *self {
		if mapping(k) {
			v = append(v, k)
		}
	}
	*self = v
}

func (self *FloatSlice) Reduce(mapping func(Float, Float) Float) Float {
	if len(*self) == 0 {
		return 0.0
	}
	if len(*self) == 1 {
		return (*self)[0]
	}
	tut := (*self)[0]
	for i := 1; i < len(*self); i++ {
		tut = mapping(tut, (*self)[i])
	}
	return tut
}

func (self *FloatSlice) Any(mapping func(Float) bool) (Float, bool) {
	for _, k := range *self {
		if (mapping(k)) {
			return k, true
		}
	}
	return 0, false
}

func (self *FloatSlice) Max() Float {
	return self.Reduce(func(prev Float, next Float) Float {
		if prev > next {
			return prev
		}
		return next
	})
}

func (self *FloatSlice) Min() Float {
	return self.Reduce(func(prev Float, next Float) Float {
		if prev < next {
			return prev
		}
		return next
	})
}
