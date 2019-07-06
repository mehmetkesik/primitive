package main

import (
	"fmt"
	"github.com/mehmetkesik/primitive"
)

func jsonExample() {
	var jo = primitive.NewJson()

	var data = []byte(`
			{
				"key1":"hello",
				"key2":"world",
				"key3":null,
				"key4":13,
				"key5":13.7,
				"key6":true,
				"key7":[
						"a",
						"b",
						1,
						1.3,
						"x",
						"y",
						{
							"key":"value"
						}
					]
			}
		`)

	err := jo.Parse(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(jo)

	fmt.Println(jo.HasKey("key1"), jo.HasKey("key13"))

	//JSON ARRAY EXAMPLE

	ja, err := jo.GetArray("key7")
	if err != nil {
		panic(err)
	}

	s, err := ja.GetString(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(ja, s)

}
