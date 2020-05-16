package database

import (
	"math/big"

	"github.com/hashcloak/go-baryon/matrix"
)

// PIRDatabase : Interface for a generalized PIR database
type PIRDatabase interface {
	Init(dbSize uint, dbBlockSize uint, initData *matrix.Matrix)

	Add(row, col uint, content big.Int) error

	Remove(row, col uint) error

	Retrieve(vector matrix.Matrix) (big.Int, error)

	// OPTIONAL: Only for IT-PIR protocols that work in the expressive query setting
	ExpressiveRetrieve(vector matrix.Matrix) (matrix.Matrix, error)
}
