package matrix

import (
	// Temporary until appropriate finite field arithmetic is written
	"math/big"
)

// Matrix: A struct that represents an nRows x nCols matrix
type Matrix struct {
	nRows    uint
	nCols    uint
	contents [][]big.Int
}

func (a *Matrix) AddElement(row, col uint, x big.Int) error {
	return nil
}

func (a *Matrix) RemoveElement(row, col uint, x big.Int) error {
	return nil
}

func (a Matrix) GetElement(row, col uint) (big.Int, error) {
	return big.Int{}, nil
}

func (a Matrix) Add(b Matrix) (Matrix, error) {
	return Matrix{}, nil
}

func (a Matrix) Substract(b Matrix) (Matrix, error) {
	return Matrix{}, nil
}

func (a Matrix) Multiply(b Matrix) (Matrix, error) {
	return Matrix{}, nil
}

func (a Matrix) Inverse(b Matrix) (Matrix, error) {
	return Matrix{}, nil
}
