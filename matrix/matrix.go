package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix []int

func New(s string) (Matrix, error) {
	rStrs := strings.Split(s, "\n")
	rNum := len(rStrs)
	cNum := len(strings.Fields(rStrs[0]))
	eNum := rNum * cNum
	vals := make([]int, 2+eNum)
	vals[0] = rNum
	vals[1] = cNum

	for i, rowStr := range rStrs {
		elems := strings.Fields(rowStr)
		if len(elems) != cNum {
			return Matrix{}, fmt.Errorf("inconsistent row length")
		}
		for j, s := range elems {
			if x, err := strconv.Atoi(s); err == nil {
				vals[2+i*cNum+j] = x
			} else {
				return Matrix{}, fmt.Errorf("string %s cannot be converted to integer: %v", s, err)
			}
		}
	}
	out := vals
	return out, nil
}
func (m Matrix) String() string {
	return fmt.Sprintf("%d %d %v", m[0], m[1], m[2:])
}

// auxiliar func
func (m Matrix) rowNum() int {
	return m[0]
}
func (m Matrix) colNum() int {
	return m[1]
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	rNum := m.rowNum()
	cNum := m.colNum()
	out := make([][]int, cNum)
	for j := 0; j < cNum; j++ {
		colOut := make([]int, rNum)
		for i := 0; i < rNum; i++ {
			colOut[i] = m[2+i*cNum+j]
		}
		out[j] = colOut
	}
	return out
}

// func (m Matrix) Cols() [][]int {
// 	out := m.Transpose().Rows()
// 	return out
// }

// func (m Matrix) Transpose() Matrix {
// 	rNum := m.rowNum()
// 	cNum := m.colNum()
// 	out := make(Matrix, cNum)
// 	for j := 0; j < cNum; j++ {
// 		out[j] = make([]int, rNum)
// 		for i := 0; i < rNum; i++ {
// 			out[j][i] = m[i][j]
// 		}
// 	}
// 	return out
// }

func (m Matrix) Rows() [][]int {
	rNum := m.rowNum()
	cNum := m.colNum()
	out := make([][]int, rNum)
	for i := 0; i < rNum; i++ {
		rowOut := make([]int, cNum)
		copy(rowOut, m[2+i*cNum:2+(i+1)*cNum])
		out[i] = rowOut
	}
	return out
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= m.rowNum() || col >= m.colNum() {
		return false
	}
	cNum := m.colNum()
	m[2+row*cNum+col] = val
	return true
}
