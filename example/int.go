package main

import (
	"fmt"
	. "github.com/mehmetkesik/primitive"
)

func intExample() {
	var i Int

	i = 13

	fmt.Println(i.Abs(), i.Pow(3), i.ToString())

}
