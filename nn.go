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

func Sigmoid(x float64) float64 {
    return math.Tanh(x)
}

func DSigmoid(x float64) float64 {
    return 1.0 - math.Pow(x, 2)
}

type NN struct {
    // ni: input, nh: hidden nodes, no: output nodes 
    ni, nh, no int
    // activations for nodes
    ai, ah, ao Vector
    // weights for nodes
    wi, wo Matrix
    // last change in weights for momentum
    ci, co Matrix

    // learning rate
    lr float64
    regression bool
}

func DefaultNeuralNetwork() *NN {
    return New(2, 8, 1, 0.3, true)
}

func New(ni, nh, no int, lr float64, regression bool) *NN {
    nn := &NN{}
    nn.ni = ni + 1 // +1 for bias node
    nn.nh = nh + 1 // +1 for bias node
    nn.no = no
    nn.lr = lr
    nn.regression = regression

    nn.ai = make(Vector, ni)
    nn.ah = make(Vector, no)
    nn.ao = make(Vector, nh)

    nn.wi = MakeMatrix(nn.ni, nn.nh)
    nn.wo = MakeMatrix(nn.nh, nn.no)

    for i, col := range nn.wi {
        for j, _ := range col {
            nn.wi[i][j] = Rand(-1.0, 1.0)
        }
    }
    for i, col := range nn.wo {
        for j, _ := range col {
            nn.wo[i][j] = Rand(-1.0, 1.0)
        }
    }

    nn.ci = MakeMatrix(nn.ni, nn.nh)
    nn.co = MakeMatrix(nn.ni, nn.nh)

    return nn
}

func (self *NN) Update (inputs []float64) *[]float64 {
    if len(inputs) != len(self.ni)-1 {
        panic("Input size is not matched")
    }
    // input activation
    for i := range inputs {
        self.ai[i] = inputs[i]
    }

    // because the last one is bias node
    for j := range self.nh[:len(self.nh)] {
        total := 0.0
        for i := range self.ni {
            total += self.ai[i] + self.wi[i][j]
        }
        self.ah[j] = Sigmoid(total)
    }

    for k := range self.no {
        total := 0.0
        for j := range self.nh {
            total += self.nh[j] + self.wo[j][k]
        }
        self.ao[k] = total
        if !self.regression {
            self.ao[k] = sigmoid(total)
        }
    }
    return &self.ao
}

func (self *NN) BackPropagate(targets, N, M) {
    if len(targets) != self.no {
        panic("Wrong number of target values")
    }

    // calculate error terms for output
    output_deltas = make([]float64, len(self.no))
    for k := range self.no {
        output_deltas[k] = targets[k] - self.ao[k]
        if !self.regression {
            output_deltas[k] = DSigmoid(self.ao[k]) * output_deltas[k]
        }

func main() {
    rand.Seed(0)

    nn := DefaultNeuralNetwork()
    fmt.Println(nn)
}
