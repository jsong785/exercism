package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(matrixString string) (Matrix, error) {
	rows := strings.Split(matrixString, "\n")
	if len(rows) <= 0 {
		return nil, errors.New("Failed to create matrix from string; no rows could be parsed.")
	}

	columns := strings.Split(rows[0], " ")
	if len(columns) <= 0 {
		return nil, errors.New("Failed to create matrix from string; no columns could be parsed.")
	}

	matrix := CreateEmptyMatrix(len(rows), len(columns))
	for rowIndex, rowString := range rows {
		if err := matrix.SetRowFromString(rowIndex, rowString); err != nil {
			return nil, err
		}
	}
	return matrix, nil
}

func CreateEmptyMatrix(m, n int) Matrix {
	matrix := make([][]int, m)
	for row := range matrix {
		matrix[row] = make([]int, n)
	}
	return Matrix(matrix)
}

func (m Matrix) SetRowFromString(rowIndex int, rowString string) error {
	if rowIndex < 0 || rowIndex >= m.RowCount() {
		return errors.New("Failure to set row from string; invalid row index.")
	}

	rowStringTrimmed := strings.TrimSpace(rowString)
	columns := strings.Split(rowStringTrimmed, " ")
	if len(columns) != m.ColumnCount() {
		return errors.New("Failure to set row from string; columns parsed doesn't match columns in matrix.")
	}

	for colIndex, colString := range columns {
		num, err := strconv.Atoi(colString)
		if err != nil {
			return err
		}
		if !m.Set(rowIndex, colIndex, num) {
			return fmt.Errorf("Error setting value for row: %d col: %d.",
				rowIndex, colIndex)
		}
	}
	return nil
}

func (m Matrix) RowCount() int {
	return len(m)
}

func (m Matrix) ColumnCount() int {
	return len(m[0])
}

func (m Matrix) Set(row, col, val int) bool {
	validIndexes := 0 <= row && row < m.RowCount() &&
		0 <= col && col < m.ColumnCount()
	if validIndexes {
		matrixRow := m[row]
		matrixRow[col] = val
	}
	return validIndexes
}

func (m Matrix) Cols() [][]int {
	transpose := m.GetTranspose()
	return [][]int(transpose)
}

func (m Matrix) Rows() [][]int {
	rowMatrix := m.GetCopy()
	return [][]int(rowMatrix)
}

func (m Matrix) GetCopy() Matrix {
	newMatrix := CreateEmptyMatrix(m.RowCount(), m.ColumnCount())
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[row]); col++ {
			newMatrix.Set(row, col, m[row][col])
		}
	}
	return newMatrix
}

func (m Matrix) GetTranspose() Matrix {
	newMatrix := CreateEmptyMatrix(m.ColumnCount(), m.RowCount())
	for col := 0; col < m.ColumnCount(); col++ {
		for row := 0; row < m.RowCount(); row++ {
			newMatrix.Set(col, row, m[row][col])
		}
	}
	return newMatrix
}
