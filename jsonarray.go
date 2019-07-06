package primitive

import (
	"errors"
	"fmt"
)

type jsonarray struct {
	jarray []interface{}
}

func NewJsonArray() jsonarray {
	j := new(jsonarray)
	j.jarray = make([]interface{}, 0)
	return *j
}

func (self *jsonarray) Len() Int {
	return Int(len(self.jarray))
}

func (self *jsonarray) AddFloat(element Float) {
	self.jarray = append(self.jarray, element)
}

func (self *jsonarray) AddString(element String) {
	self.jarray = append(self.jarray, element)
}

func (self *jsonarray) AddBool(element bool) {
	self.jarray = append(self.jarray, element)
}

func (self *jsonarray) AddNil(key String) {
	self.jarray = append(self.jarray, nil)
}

func (self *jsonarray) AddArray(key String, element jsonarray) {
	self.jarray = append(self.jarray, element.jarray)
}

func (self *jsonarray) AddObject(key String, element myjson) {
	self.jarray = append(self.jarray, element.jmap)
}

func (self *jsonarray) Set(index Int, element interface{}) {
	switch v := element.(type) {
	case int:
		element = Int(v)
		break
	case float64:
		element = Float(v)
		break
	case string:
		element = String(v)
		break
	}
	self.jarray[index] = element
}

func (self *jsonarray) Delete(index Int) {
	self.jarray = append(self.jarray[:index], self.jarray[index+1:]...)
}

func (self *jsonarray) GetString(index Int) (String, error) {
	i, ok := self.jarray[index].(String)
	if !ok {
		return "", errors.New(fmt.Sprintf("got data of type '%T' but wanted 'String'", self.jarray[index]))
	}
	return i, nil
}

func (self *jsonarray) GetFloat(index Int) (Float, error) {
	i, ok := self.jarray[index].(Float)
	if !ok {
		return 0.0, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'Float'", self.jarray[index]))
	}
	return i, nil
}

func (self *jsonarray) GetBool(index Int) (bool, error) {
	i, ok := self.jarray[index].(bool)
	if !ok {
		return false, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'bool'", self.jarray[index]))
	}
	return i, nil
}

func (self *jsonarray) GetArray(index Int) (*jsonarray, error) {
	i, ok := self.jarray[index].([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'Json Array'", self.jarray[index]))
	}
	ja := new(jsonarray)
	ja.jarray = i
	return ja, nil
}

func (self *jsonarray) GetObject(index Int) (*myjson, error) {
	i, ok := self.jarray[index].(map[String]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'Json Object'", self.jarray[index]))
	}
	j := new(myjson)
	j.jmap = i
	return j, nil
}

func (self *jsonarray) IsNil(index Int) bool {
	if self.jarray[index] == nil {
		return true
	}
	return false
}

func (self *jsonarray) AppendStringSlice(slice StringSlice) {
	for _, v := range slice {
		self.AddString(v)
	}
}

func (self *jsonarray) AppendFloatSlice(slice FloatSlice) {
	for _, v := range slice {
		self.AddFloat(v)
	}
}

func (self *jsonarray) AppendBoolSlice(slice []bool) {
	for _, v := range slice {
		self.AddBool(v)
	}
}
