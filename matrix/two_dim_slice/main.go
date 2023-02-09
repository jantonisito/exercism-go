// demo of making matrix
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) rowNum() int {
	return len(m)
}
func (m Matrix) colNum() int {
	return len(m[0])
}

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

func main() {
	m := Matrix{{1, 2}, {3, 4}}
	fmt.Println(m)
	r := 2
	c := 2
	k := 1
	m1 := make(Matrix, r)
	for i := range m1 {
		row := make([]int, c)
		for j := range row {
			row[j] = k
			k += 1
		}
		m1[i] = row

	}
	fmt.Println(m1)
	m2Str := "1 2\n3 4"
	m2, _ := New(m2Str)
	fmt.Println(m2)
	//fmt.Printf("matrix size %d x %d", m2.rowNum(), m2.colNum())

	m3 := m2.Transpose()
	fmt.Println(m3)
}
