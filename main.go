package main

import "fmt"

func DebugAdd() {
	s := NewIntensitySegments()
	fmt.Println(s)
	s.Add(10, 30, 1)
	fmt.Println(s)
	s.Add(20, 40, 1)
	fmt.Println(s)
	//s.Add(10, 40, -2)
	//fmt.Println(s)
	//s.Add(10, 40, -1)
	//fmt.Println(s)
	//s.Add(10, 40, -1)
	//fmt.Println(s)
}

func DebugSet() {
	s := NewIntensitySegments()
	fmt.Println(s)
	s.Add(10, 30, 1)
	fmt.Println(s)
	s.Add(20, 40, 1)
	fmt.Println(s)

	s.Set(15, 25, 2)
	//s.Set(15, 35, 3)
	fmt.Println(s)
}

func main() {
	DebugAdd()
	//DebugSet()
}
