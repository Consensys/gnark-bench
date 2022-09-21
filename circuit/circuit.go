package circuit

import (
	"github.com/consensys/gnark-crypto/ecc"
	bls12377fr "github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	bls12381fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	bls24315fr "github.com/consensys/gnark-crypto/ecc/bls24-315/fr"
	bn254fr "github.com/consensys/gnark-crypto/ecc/bn254/fr"
	bw6633fr "github.com/consensys/gnark-crypto/ecc/bw6-633/fr"
	bw6761fr "github.com/consensys/gnark-crypto/ecc/bw6-761/fr"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
)

var BenchCircuits map[string]BenchCircuit

type BenchCircuit interface {
	Circuit(size int) frontend.Circuit
	Witness(size int, curveID ecc.ID) *witness.Witness
}

func init() {
	BenchCircuits = make(map[string]BenchCircuit)
	BenchCircuits["expo"] = &defaultCircuit{}
}

type defaultCircuit struct {
}

func (d *defaultCircuit) Circuit(size int) frontend.Circuit {
	return &benchCircuit{n: size}
}

func (d *defaultCircuit) Witness(size int, curveID ecc.ID) *witness.Witness {
	witness := benchCircuit{n: size}

	witness.X = (2)

	switch curveID {
	case ecc.BN254:
		// compute expected Y
		var expectedY bn254fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y = (expectedY)
	case ecc.BLS12_381:
		// compute expected Y
		var expectedY bls12381fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y = (expectedY)
	case ecc.BLS12_377:
		// compute expected Y
		var expectedY bls12377fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y = (expectedY)
	case ecc.BLS24_315:
		// compute expected Y
		var expectedY bls24315fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y = (expectedY)
	case ecc.BW6_761:
		// compute expected Y
		var expectedY bw6761fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y = (expectedY)
	case ecc.BW6_633:
		// compute expected Y
		var expectedY bw6633fr.Element
		expectedY.SetInterface(2)
		for i := 0; i < size; i++ {
			expectedY.Mul(&expectedY, &expectedY)
		}

		witness.Y = (expectedY)
	default:
		panic("not implemented")
	}

	witness.Y = (2)
	w, err := frontend.NewWitness(&witness, curveID.ScalarField())
	if err != nil {
		panic(err)
	}
	return w
}

// benchCircuit is a simple circuit that checks X*X*X*X*X... == Y
type benchCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
	n int
}

func (circuit *benchCircuit) Define(api frontend.API) error {
	for i := 0; i < circuit.n; i++ {
		_ = api.Mul(circuit.X, circuit.Y)
	}
	api.AssertIsEqual(circuit.X, circuit.Y)
	return nil
}
