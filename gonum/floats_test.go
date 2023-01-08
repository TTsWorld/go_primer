package gonum

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"

	"gonum.org/v1/gonum/floats"
)

// 在 Go 语言中可以采用 slice 来保存一组值，并且将其作为数组来使用。
// 为了便于开展数值计算，Gonum 科学计算包对 slice 进行封装和扩展，引入了 VectorDense 、 Dense 、CDense和 BandDense等新的数据对象。
// 其中，VectorDense 和 Dense 可分别用于生成一维数组（向量）和多维数组（矩阵）。
func TestFloatsAdd(t *testing.T) {
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

func TestFloatsDot(t *testing.T) {
	s1 := []float64{0, 1, 2, 4}
	s2 := []float64{0, 2, 3, 6}

	s := floats.Dot(s1, s2)
	fmt.Println("s=", s)
}

//将 slice 转换为 VectorDense 数据类型，然后进行矩阵运算
func TestFloatsDot2(t *testing.T) {
	s1 := []float64{0, 1, 2, 4}
	s2 := []float64{0, 2, 3, 6}
	s := floats.Dot(s1, s2)

	vs1 := mat.NewVecDense(4, s1)
	vs2 := mat.NewVecDense(4, s2)
	vs := mat.Dot(vs1, vs2)

	fmt.Println("s=", s)
	fmt.Println("vs=", vs)
}

func TestMaxMin(t *testing.T) {
	s1 := []float64{0, 1, 2, 4}
	s2 := []float64{0, 2, 3, 6}

	fmt.Println("max(s1)=", floats.Max(s1))
	fmt.Println("min(s1)=", floats.Min(s2))

}

func TestAA(t *testing.T) {

}
