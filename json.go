package primitive

import (
	"encoding/json"
	"errors"
	"fmt"
)

type myjson struct {
	jmap map[String]interface{}
}

func NewJson() myjson {
	j := new(myjson)
	j.jmap = make(map[String]interface{})
	return *j
}

func (self *myjson) Parse(data []byte) error {
	var v = make(map[string]interface{})
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	self.jmap = self.parse(v)
	return nil
}

func (self *myjson) AddFloat(key String, element Float) {
	self.jmap[key] = element
}

func (self *myjson) AddString(key String, element String) {
	self.jmap[key] = element
}

func (self *myjson) AddBool(key String, element bool) {
	self.jmap[key] = element
}

func (self *myjson) AddNil(key String) {
	self.jmap[key] = nil
}

func (self *myjson) AddArray(key String, element jsonarray) {
	self.jmap[key] = element.jarray
}

func (self *myjson) AddObject(key String, element myjson) {
	self.jmap[key] = element.jmap
}

func (self *myjson) GetString(key String) (String, error) {
	s, ok := self.jmap[key].(String)
	if !ok {
		return "", errors.New(fmt.Sprintf("got data of type '%T' but wanted 'String'", self.jmap[key]))
	}
	return s, nil
}

func (self *myjson) GetFloat(key String) (Float, error) {
	f, ok := self.jmap[key].(Float)
	if !ok {

		return 0.0, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'Float'", self.jmap[key]))
	}
	return f, nil
}

func (self *myjson) GetBool(key String) (bool, error) {
	b, ok := self.jmap[key].(bool)
	if !ok {
		return false, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'bool'", self.jmap[key]))
	}
	return b, nil
}

func (self *myjson) GetArray(key String) (*jsonarray, error) {
	m, ok := self.jmap[key].([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'Json Array'", self.jmap[key]))
	}
	ja := new(jsonarray)
	ja.jarray = m
	return ja, nil
}

func (self *myjson) GetObject(key String) (*myjson, error) {
	m, ok := self.jmap[key].(map[String]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("got data of type '%T' but wanted 'Json Object'", self.jmap[key]))
	}
	j := new(myjson)
	j.jmap = m
	return j, nil
}

func (self *myjson) IsNil(key String) bool {
	if self.HasKey(key) {
		if self.jmap[key] == nil {
			return true
		}
		return false
	} else {
		panic("key '" + string(key) + "' is not defined!")
	}
}

func (self *myjson) HasKey(key String) bool {
	_, ok := self.jmap[key]
	if !ok {
		return false
	}
	return true
}

func (self *myjson) Delete(key String) {
	if self.HasKey(key) {
		delete(self.jmap, key)
	}
}

func (self *myjson) parse(json map[string]interface{}) map[String]interface{} {
	newj := make(map[String]interface{})
	for k, v := range json {
		switch v.(type) {
		case float64:
			tutk := String(k)
			tutv := Float(v.(float64))
			newj[tutk] = tutv
			break;
		case string:
			tutk := String(k)
			tutv := String(v.(string))
			newj[tutk] = tutv
			break;
		case bool:
			tutk := String(k)
			newj[tutk] = v
			break;
		case map[string]interface{}:
			tutk := String(k)
			newj[tutk] = self.parse(v.(map[string]interface{}))
			break;
		case []interface{}:
			tutk := String(k)
			newj[tutk] = self.parsearray(v.([]interface{}))
			break;
		case nil:
			tutk := String(k)
			newj[tutk] = nil
			break;
		}
	}
	return newj
}

func (self *myjson) parsearray(jsonarray []interface{}) []interface{} {
	for i, v := range jsonarray {
		switch v.(type) {
		case float64:
			tutv := Float(v.(float64))
			jsonarray[i] = tutv
			break;
		case string:
			tutv := String(v.(string))
			jsonarray[i] = tutv
			break;
		case bool:
			break;
		case map[string]interface{}:
			jsonarray[i] = self.parse(v.(map[string]interface{}))
			break;
		case []interface{}:
			jsonarray[i] = self.parsearray(v.([]interface{}))
			break;
		case nil:
			break;
		}
	}
	return jsonarray
}
