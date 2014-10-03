package main

import (
    "fmt"
    "math"
    "math/rand"
)

type Matrix [][]float64
type Vector []float64

func Rand(a float64, b float64) float64 {
    return ((b-a)*rand.Float64() + a)
}

func MakeMatrix(I int, J int, fill_opt ...float64) Matrix {
    fill := 0.0
    if len(fill_opt) > 0 {
        fill = fill_opt[0]
    }
    matrix := make(Matrix, J)
    for i := range matrix {
        matrix[i] = make(Vector, I)

        for j, _ := range matrix[i] {
            matrix[i][j] = fill
        }
    }

    return matrix
}

func Transfer(x float64) float64 {
    return math.Tanh(x)
}

func Dtransfer(x float64) float64 {
    return 1.0 - math.Pow(x, 2)
}

type NN struct {
    // ni: input, nh: hidden nodes, no: output nodes 
    ni, nh, no int
    // activations for nodes
    ai, ah, ao Vector
    // weights for nodes
    wi, wo Matrix
    regression bool
}

func New(ni, nh, no int, regression bool) *NN {
    nn := &NN{}
    nn.ni = ni + 1 // +1 for bias node
    nn.nh = nh + 1 // +1 for bias node
    nn.no = no
    nn.regression = regression

    nn.ai = make(Vector, ni)
    nn.ah = make(Vector, no)
    nn.ao = make(Vector, nh)

    nn.wi = MakeMatrix(nn.ni, nn.nh)
    nn.wo = MakeMatrix(nn.nh, nn.no)

    for i, col := range nn.wi {
        for j, _ := range col {
            nn.wi[i][j] = Rand(-1.0, 1.0)
            fmt.Println(nn.wi[i][j])
        }
    }

    return nn
}

func main() {
    rand.Seed(0)

    fmt.Println("123");
    mat := MakeMatrix(3, 4, 0.0)
    fmt.Println(mat[0][0])

    nn := New(2, 3, 1, false)
    fmt.Println(nn)
}
