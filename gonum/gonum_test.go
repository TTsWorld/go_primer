package gonum

import (
	"fmt"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"

	"gonum.org/v1/gonum/floats"
)

func Test01(t *testing.T) {
	// If one wants to store the result in a
	// new container, just make a new slice
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{5, 6, 7, 8}
	s3 := []float64{1, 1, 1, 1}
	dst := make([]float64, len(s1))

	floats.AddTo(dst, s1, s2)
	floats.Add(dst, s3)

	fmt.Println("dst =", dst)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)
}

func Test02(t *testing.T) {
	a := mat.NewDense(4, 4, nil) // 生成一个空的 4x4 矩阵

	// 对矩阵 a 进行随机赋值
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a.Set(i, j, rand.Float64())
		}
	}

	// Create a matrix formatting value with a prefix using Python format...
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	// and then print with and without layout formatting.
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}
