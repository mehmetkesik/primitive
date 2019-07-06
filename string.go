package primitive

import (
	"fmt"
	"strconv"
	"strings"
)

type String string

func (self *String) Get(index Int) String {
	return String(fmt.Sprintf("%c", (*self)[index]))
}

func (self *String) Set(index Int, c String) {
	if len(c) != 1 {
		return
	}
	var tut = []rune(c)
	var t = []rune(string(*self))
	t[index] = tut[0]
	*self = String(string(t))
}

func (self *String) ToInt() (Int, error) {
	x, err := strconv.ParseInt(string(*self), 10, 64)
	if err != nil {
		return *new(Int), err
	}
	return Int(x), nil
}

func (self *String) ToFloat() (Float, error) {
	x, err := strconv.ParseFloat(string(*self), 10)
	if err != nil {
		return *new(Float), err
	}
	return Float(x), nil
}

func (self *String) Map(mapping func(String) String) {
	for i := 0; i < len(*self); i++ {
		tut := mapping((*self).Get(Int(i)))
		if len(tut) != 1 {
			continue
		}
		(*self).Set(Int(i), tut)
	}
}

func (self *String) Split(c String) []String {
	var v []String
	tut := strings.Split(string(*self), string(c))
	for _, k := range tut {
		v = append(v, String(k))
	}
	return v
}

func (self *String) ToUpper() String {
	return String(strings.ToUpper(string(*self)))
}

func (self *String) ToLower() String {
	return String(strings.ToLower(string(*self)))
}

func (self *String) Replace(old String, new String, n Int) String {
	return String(strings.Replace(string(*self), string(old), string(new), int(n)))
}

func (self *String) ReplaceAll(old String, new String) String {
	return String(strings.ReplaceAll(string(*self), string(old), string(new)))
}

func (self *String) Contains(subs String) bool {
	return strings.Contains(string(*self), string(subs))
}

type StringSlice []String

func (self *StringSlice) Len() Int {
	return Int(len(*self))
}

func (self *StringSlice) Append(s String) {
	*self = append(*self, s)
}

func (self *StringSlice) Delete(index Int) {
	*self = append((*self)[:index], (*self)[index+1:]...)
}

func (self *StringSlice) Map(mapping func(String) String) {
	for i, k := range *self {
		(*self)[i] = mapping(k)
	}
}

func (self *StringSlice) Filter(mapping func(String) bool) {
	var v StringSlice
	for _, k := range *self {
		if mapping(k) {
			v = append(v, k)
		}
	}
	*self = v
}

func (self *StringSlice) Reduce(mapping func(String, String) String) String {
	if len(*self) == 0 {
		return ""
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

func (self *StringSlice) Any(mapping func(String) bool) (String, bool) {
	for _, k := range *self {
		if (mapping(k)) {
			return k, true
		}
	}
	return "", false
}
