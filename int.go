package primitive

import (
	"math"
	"strconv"
)

type Int int64

func (self *Int) ToString() String {
	return String(strconv.FormatInt(int64(*self), 10))
}

func (self *Float) ToFloat() Float {
	return Float(*self)
}

func (self *Int) Abs() Int {
	return Int(math.Abs(float64(*self)))
}

func (self *Int) Pow(y Float) Int {
	return Int(math.Pow(float64(*self), float64(y)))
}

type IntSlice []Int

func (self *IntSlice) Get(index Int) Int {
	return (*self)[index]
}

func (self *IntSlice) Len() Int {
	return Int(len(*self))
}

func (self *IntSlice) Append(i Int) {
	*self = append(*self, i)
}

func (self *IntSlice) Delete(index Int) {
	*self = append((*self)[:index], (*self)[index+1:]...)
}

func (self *IntSlice) Map(mapping func(Int) Int) {
	for i, k := range *self {
		(*self)[i] = mapping(k)
	}
}

func (self *IntSlice) Filter(mapping func(Int) bool) {
	var v IntSlice
	for _, k := range *self {
		if mapping(k) {
			v = append(v, k)
		}
	}
	*self = v
}

func (self *IntSlice) Reduce(mapping func(Int, Int) Int) Int {
	if len(*self) == 0 {
		return 0
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

func (self *IntSlice) Any(mapping func(Int) bool) (Int, bool) {
	for _, k := range *self {
		if (mapping(k)) {
			return k, true
		}
	}
	return 0, false
}

func (self *IntSlice) Max() Int {
	return self.Reduce(func(prev Int, next Int) Int {
		if prev > next {
			return prev
		}
		return next
	})
}

func (self *IntSlice) Min() Int {
	return self.Reduce(func(prev Int, next Int) Int {
		if prev < next {
			return prev
		}
		return next
	})
}
