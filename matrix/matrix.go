package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	rNum := len(rows)
	matrix := make(Matrix, rNum)

	//
	cNum := len(strings.Fields(rows[0]))
	for i, rowStr := range rows {
		elems := strings.Fields(rowStr)
		if len(elems) != cNum {
			return Matrix{}, fmt.Errorf("inconsistent row length")
		}
		row := make([]int, cNum)
		for j, s := range elems {
			if x, err := strconv.Atoi(s); err == nil {
				row[j] = x
			} else {
				return Matrix{}, fmt.Errorf("string %s cannot be converted to integer: %v", s, err)
			}
		}
		matrix[i] = row
	}
	return matrix, nil
}

// auxiliar func
func (m Matrix) rowNum() int {
	return len(m)
}
func (m Matrix) colNum() int {
	return len(m[0])
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	rNum := m.rowNum()
	cNum := m.colNum()
	out := make([][]int, cNum)
	for j := 0; j < cNum; j++ {
		colOut := make([]int, rNum)
		for i, row := range m {
			colOut[i] = row[j]
		}
		out[j] = colOut
	}
	return out
}

// func (m Matrix) Cols() [][]int {
// 	out := m.Transpose().Rows()
// 	return out
// }

func (m Matrix) Transpose() Matrix {
	rNum := m.rowNum()
	cNum := m.colNum()
	out := make(Matrix, cNum)
	for j := 0; j < cNum; j++ {
		out[j] = make([]int, rNum)
		for i := 0; i < rNum; i++ {
			out[j][i] = m[i][j]
		}
	}
	return out
}

func (m Matrix) Rows() [][]int {
	rNum := m.rowNum()
	cNum := m.colNum()
	out := make([][]int, rNum)
	for i, row := range m {
		rowOut := make([]int, cNum)
		copy(rowOut, row)
		out[i] = rowOut
	}
	return out
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= m.rowNum() || col >= m.colNum() {
		return false
	}
	mRow := m[row]
	mRow[col] = val
	return true
}
