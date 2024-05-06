package main

import (
    "fmt"
    "github.com/AlexNos/mymath"
)

func Sqrt(x float64) float64 {
    return mymath.Sqrt(x)
}

func Abs(x float64) float64 {
	return mymath.Abs(x)
}

func Ceil(x float64) float64 {
    return mymath.Ceil(x)
}

func Floor(x float64) float64 {
    return mymath.Floor(x)
}

func Pow(x float64, y float64) float64 {
    return mymath.Pow(x, y)
}

func Max(x float64, y float64) float64 {
    return mymath.Max(x, y)
}

func Min(x float64, y float64) float64 {
    return mymath.Min(x, y)
}

func Yn(n int, x float64) float64 {
    return mymath.Yn(n, x)
}



func main() {
    fmt.Println(mymath.Sqrt(4)) // 2
}