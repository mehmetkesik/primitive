# primitive
A simple and functional json type that uses them with a simplified integer, float, and string for the Go programming language.

# Install
The project does not have any dependencies other than the standard library. Just run `go get github.com/mehmetkesik/primitive` to install the project.

# Example

## Int
```
var i Int
i = 13
fmt.Println(i.Abs(), i.Pow(3), i.ToString())
```

## Float
```
var f Float
f = 13.7
fmt.Println(f.ToInt(), f.ToString(), f.Pow(3), f.Abs())
```

## String
```
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
```
