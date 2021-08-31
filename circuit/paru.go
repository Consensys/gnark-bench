package circuit

import (
	"github.com/OlivierBBB/paru/circuits"
	"github.com/OlivierBBB/paru/paru"
	"github.com/consensys/gnark-crypto/ecc"
	bn254fr "github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/frontend"
)

type paruCircuit struct {
}

const (
	chunkSize            = 2
	accountTreeDepth     = 32
	balanceTreeDepth     = 16
	moneyOrderTreeDepth  = 32
	moneyOrderBatchDepth = 16
)

func (d *paruCircuit) Circuit(nMoneyOrderCreationRequests int) frontend.Circuit {

	// moneyOrderBatchDepth := common.Log2Ceil(nMocrs)
	c := circuits.AllocateMocReqBatchCircuit(
		nMoneyOrderCreationRequests, accountTreeDepth, balanceTreeDepth, moneyOrderBatchDepth, moneyOrderTreeDepth, chunkSize,
	)
	return &c
}

func (d *paruCircuit) Witness(nMoneyOrderCreationRequests int, curveID ecc.ID) frontend.Circuit {

	// this should be generated from the pb.ProverData
	mocrs, oldAccountRoot, oldMoRoot, moBatchInclusionProof, moBatchIndex := GenerateMocrBatchWitness(
		nMoneyOrderCreationRequests, accountTreeDepth, balanceTreeDepth, moneyOrderBatchDepth, moneyOrderTreeDepth,
	)

	witness := circuits.AllocateMocReqBatchCircuit(
		nMoneyOrderCreationRequests, accountTreeDepth, balanceTreeDepth, moneyOrderBatchDepth, moneyOrderTreeDepth, chunkSize,
	)
	paruio := paru.NewGKRio(2 * chunkSize)
	witness.AssignPreGKR(&paruio, mocrs, oldAccountRoot, oldMoRoot, moBatchInclusionProof, moBatchIndex)

	proof, q, qPrime := circuits.GkrProve(&paruio)
	witness.AssignPostGKR(proof, q, qPrime)
	return &witness

}

func GenerateMocrBatchWitness(
	nMocrs, accountDepth, balanceDepth, moBatchDepth, moDepth int,
) (
	mocrs []paru.MocRequest,
	oldAccountRoot, oldMoRoot bn254fr.Element,
	moBatchInclusionProof paru.MerkleProof,
	moBatchIndex uint64,
) {
	state := paru.NewPARU(42, uint64(nMocrs), uint64(accountDepth), uint64(balanceDepth), uint64(moDepth))
	oldAccountRoot = state.AccountTree.MT.Root()
	oldMoRoot = state.MoneyOrderTree.MT.Root()

	mos := paru.NewMoneyOrderBatch(78, uint64(nMocrs-1), uint64(balanceDepth), uint64(nMocrs), uint64(moBatchDepth))
	mocrs, moBatchInclusionProof, moBatchIndex = state.ProcessMocBatch(mos, uint64(moBatchDepth))
	return mocrs, oldAccountRoot, oldMoRoot, moBatchInclusionProof, moBatchIndex
}
