package utxo

import (
	"github.com/tclchiam/oxidize-go/blockchain/engine/iter"
	"github.com/tclchiam/oxidize-go/blockchain/entity"
	"github.com/tclchiam/oxidize-go/identity"
)

type utxoCrawlerEngine struct {
	repository entity.BlockRepository
}

func NewCrawlerEngine(repository entity.BlockRepository) Engine {
	return &utxoCrawlerEngine{repository: repository}
}

func (engine *utxoCrawlerEngine) FindUnspentOutputs(spender *identity.Identity) (*TransactionOutputSet, error) {
	inputs, err := findInputs(engine.repository)
	if err != nil {
		return nil, err
	}

	outputsByTx, err := findOutputsByTransaction(engine.repository)
	if err != nil {
		return nil, err
	}

	spentOutputs := inputs.
		Reduce(make(map[*entity.Hash][]*entity.Output), addInputToMap).(map[*entity.Hash][]*entity.Output)

	return outputsByTx.
		Filter(func(_ *entity.Transaction, output *entity.Output) bool { return output.ReceivedBy(spender) }).
		Filter(isUnspent(spentOutputs)), nil
}

func findInputs(repository entity.BlockRepository) (entity.SignedInputs, error) {
	var gatherInputs = func(res interface{}, tx *entity.Transaction) interface{} {
		return res.(entity.SignedInputs).Append(entity.NewSignedInputs(tx.Inputs))
	}

	inputs := entity.EmptySingedInputs()

	err := iter.ForEachBlock(repository, func(block *entity.Block) {
		inputs = block.Transactions().Reduce(inputs, gatherInputs).(entity.SignedInputs)
	})

	return inputs, err
}

func findOutputsByTransaction(repository entity.BlockRepository) (*TransactionOutputSet, error) {
	outputsForAddress := NewTransactionSet()

	err := iter.ForEachBlock(repository, func(block *entity.Block) {
		for _, transaction := range block.Transactions() {
			addToTxSet := func(res interface{}, output *entity.Output) interface{} {
				return res.(*TransactionOutputSet).Add(transaction, output)
			}

			outputsForAddress = entity.NewOutputs(transaction.Outputs).
				Reduce(outputsForAddress, addToTxSet).(*TransactionOutputSet)
		}
	})

	return outputsForAddress, err
}

var isUnspent = func(spentOutputs map[*entity.Hash][]*entity.Output) func(*entity.Transaction, *entity.Output) bool {
	return func(transaction *entity.Transaction, output *entity.Output) bool {
		if outputs, ok := spentOutputs[transaction.ID]; ok {
			for _, spentOutput := range outputs {
				if spentOutput.IsEqual(output) {
					return false
				}
			}
		}
		return true
	}
}

var addInputToMap = func(res interface{}, input *entity.SignedInput) interface{} {
	outputs := res.(map[*entity.Hash][]*entity.Output)
	transactionId := input.OutputReference.ID
	outputs[transactionId] = append(outputs[transactionId], input.OutputReference.Output)

	return res
}
