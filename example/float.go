package main

import (
	"fmt"
	. "github.com/mehmetkesik/primitive"
)

func floatExample() {
	var f Float

	f = 13.7
	fmt.Println(f.ToInt(), f.ToString(), f.Pow(3), f.Abs())
}
