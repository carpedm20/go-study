package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    ans := make([][]uint8, dy)
    for i:= 0; i<dy; i++ {
      tmp := make([]uint8, dx)
        for idx := range tmp {
            tmp[idx] = uint8(0+idx)
        }
        ans[i] = tmp
    }
    return ans
}

func main() {
    pic.Show(Pic)
}
