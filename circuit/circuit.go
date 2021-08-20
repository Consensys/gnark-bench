package circuit

import (
	"github.com/consensys/gnark-crypto/ecc"
	bls12381fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	bn254fr "github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/frontend"
)

var BenchCircuits map[string]BenchCircuit

type BenchCircuit interface {
	Circuit(size int) frontend.Circuit
	Witness(size int, curveID ecc.ID) frontend.Circuit
}

func init() {
	BenchCircuits = make(map[string]BenchCircuit)
	BenchCircuits["expo"] = &defaultCircuit{}
	BenchCircuits["paru"] = &paruCircuit{}
}

type defaultCircuit struct {
}

func (d *defaultCircuit) Circuit(size int) frontend.Circuit {
	return &benchCircuit{n: size}
}

func (d *defaultCircuit) Witness(size int, curveID ecc.ID) frontend.Circuit {
	witness := benchCircuit{n: size}

	witness.X.Assign(2)

	switch curveID {
	case ecc.BN254:
		// compute expected Y
		var expectedY bn254fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y.Assign(expectedY)
	case ecc.BLS12_381:
		// compute expected Y
		var expectedY bls12381fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y.Assign(expectedY)
	default:
		panic("not implemented")
	}

	return &witness
}

// benchCircuit is a simple circuit that checks X*X*X*X*X... == Y
type benchCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
	n int
}

func (circuit *benchCircuit) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	for i := 0; i < circuit.n; i++ {
		circuit.X = cs.Mul(circuit.X, circuit.X)
	}
	cs.AssertIsEqual(circuit.X, circuit.Y)
	return nil
}
