package main

import (
	"fmt"
	. "github.com/mehmetkesik/primitive"
)

func stringExample() {
	var s String

	s = "hello world"

	s.Set(5, "-")
	fmt.Println(s, s.Get(4))

	s.Map(func(s String) String {
		if s == "o" {
			return s.ToUpper()
		}
		return s
	})

	fmt.Println(s)

	fmt.Println(s.ToUpper(), s.ToLower(), s.Split("-"), s.Contains("he"),
		s.ReplaceAll("l", "L"), s.Replace("l", "L", 1))

	s = "13"

	i, err := s.ToInt()
	if err != nil {
		panic(err)
	}
	f, err := s.ToFloat()
	if err != nil {
		panic(err)
	}
	fmt.Println(i, f)

}
