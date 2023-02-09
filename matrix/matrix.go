package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rNum int
	cNum int
	vals []int
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
				vals[i*cNum+j] = x
			} else {
				return Matrix{}, fmt.Errorf("string %s cannot be converted to integer: %v", s, err)
			}
		}
	}
	out := Matrix{rNum, cNum, vals}
	return out, nil
}
func (m Matrix) String() string {
	return fmt.Sprintf("%d %d %v", m.rNum, m.cNum, m.vals)
}

// auxiliar func
func (m Matrix) rowNum() int {
	return m.rNum
}
func (m Matrix) colNum() int {
	return m.cNum
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	rNum := m.rowNum()
	cNum := m.colNum()
	out := make([][]int, cNum)
	for j := 0; j < cNum; j++ {
		colOut := make([]int, rNum)
		for i := 0; i < rNum; i++ {
			colOut[i] = m.vals[i*cNum+j]
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
	rNum := m.rNum
	cNum := m.cNum
	out := make([][]int, rNum)
	for i := 0; i < rNum; i++ {
		rowOut := make([]int, cNum)
		copy(rowOut, m.vals[i*cNum:(i+1)*cNum])
		out[i] = rowOut
	}
	return out
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= m.rNum || col >= m.cNum {
		return false
	}
	m.vals[row*m.rNum+col] = val
	return true
}
