// demo of making matrix
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	rNum int
	cNum int
	vals *[]int
}

func New(s string) (Matrix, error) {
	rStrs := strings.Split(s, "\n")
	rNum := len(rStrs)
	cNum := len(strings.Fields(rStrs[0]))
	eNum := rNum * cNum
	vals := make([]int, eNum)

	for i, rowStr := range rStrs {
		elems := strings.Fields(rowStr)
		if len(elems) != cNum {
			return Matrix{}, fmt.Errorf("inconsistent row length")
		}
		for j, s := range elems {
			if x, err := strconv.Atoi(s); err == nil {
				vals[i*rNum+j] = x
			} else {
				return Matrix{}, fmt.Errorf("string %s cannot be converted to integer: %v", s, err)
			}
		}
	}
	out := Matrix{rNum, cNum, &vals}
	return out, nil
}
func (m Matrix) String() string {
	return fmt.Sprintf("%d %d %v", m.rNum, m.cNum, m.vals)
}

// func (m Matrix) Transpose() Matrix {
// 	rNum := m.rNum
// 	cNum := m.cNum
// 	out := make(Matrix, cNum)
// 	for j := 0; j < cNum; j++ {
// 		out[j] = make([]int, rNum)
// 		for i := 0; i < rNum; i++ {
// 			out[j][i] = m[i][j]
// 		}
// 	}
// 	return out
// }

func main() {
	m := Matrix{2, 2, &[]int{1, 2, 3, 4}}
	fmt.Println(m)

	m2Str := "1 2\n3 4"
	m2, _ := New(m2Str)
	fmt.Println(m2)
	//fmt.Printf("matrix size %d x %d", m2.rowNum(), m2.colNum())
}
