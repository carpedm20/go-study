package main

import (
    "fmt"
    "math"
)

type Matrix [][]float64
type Vector []float64

func Rand(a float64, b float64) float64 {
    return ((b-a)*math.rand.Float64() + a)
}

func MakeMatrix(I int, J int, fill float64) Matrix {
    matrix := make(Matrix, J)
    for i := range matrix {
        matrix[i] = make(Vector, I)
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
    ni, nh, no float64
    // activations for nodes
    ai, ah, ao Vector
    // weights for nodes
    wi, wo Matrix
    regression bool
}

func New(ni, nh, no float64, regression bool) *NN {
    nn = new(NN)
    nn.ni = ni + 1 // +1 for bias node
    nn.nh = nh + 1 // +1 for bias node
    nn.no = no
    nn.regression = regression

    nn.ai = make(Vector, ni)
    nn.ah = make(Vector, no)
    nn.ao = make(Vector, nh)

    nn.wi = MakeMatrix(nn.ni, nn.nh)
    nn.wo = MakeMatrix(nn.nh, nn.no)

    for i:= range nn.nh {
    }

    return &nn
}

func main() {
    rand.Seed(0)

    fmt.Println("123");
    mat := MakeMatrix(3, 4, 0.0)
    fmt.Println(mat[0][0])
}
