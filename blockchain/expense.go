package blockchain

import (
	"fmt"

	"github.com/tclchiam/block_n_go/blockchain/entity"
	"github.com/tclchiam/block_n_go/blockchain/tx"
	"github.com/tclchiam/block_n_go/crypto"
	"github.com/tclchiam/block_n_go/wallet"
	"github.com/tclchiam/block_n_go/encoding"
)

func (bc *Blockchain) buildExpenseTransaction(sender, receiver *wallet.Wallet, expense uint) (*entity.Transaction, error) {
	senderAddress := sender.GetAddress()

	unspentOutputs, err := bc.findUnspentOutputs(senderAddress)
	if err != nil {
		return nil, err
	}

	balance := calculateBalance(unspentOutputs)
	if balance < expense {
		return nil, fmt.Errorf("account '%s' does not have enough to send '%d', due to balance '%d'", senderAddress, expense, balance)
	}

	liquidBalance := uint(0)
	takeMinimumToMeetExpense := func(_ *entity.Transaction, output *entity.Output) bool {
		take := liquidBalance < expense
		if take {
			liquidBalance += output.Value
		}
		return take
	}

	buildInputs := func(res interface{}, transaction *entity.Transaction, output *entity.Output) interface{} {
		input := entity.NewUnsignedInput(transaction.ID, output, sender.PublicKey)
		return res.(entity.UnsignedInputs).Add(input)
	}

	inputs := unspentOutputs.
		Filter(takeMinimumToMeetExpense).
		Reduce(entity.EmptyUnsignedInputs(nil), buildInputs).(entity.UnsignedInputs)

	outputs := entity.EmptyOutputs().
		Add(entity.NewOutput(expense, receiver.GetAddress()))

	if liquidBalance-expense > 0 {
		outputs = outputs.Add(entity.NewOutput(liquidBalance-expense, senderAddress))
	}

	finalizedOutputs := outputs.Reduce(make([]*entity.Output, 0), collectOutputs).([]*entity.Output)
	signedInputs := inputs.Reduce(make([]*entity.SignedInput, 0), signInputs(finalizedOutputs, sender.PrivateKey)).([]*entity.SignedInput)

	return entity.NewTx(signedInputs, finalizedOutputs, encoding.NewTransactionGobEncoder()), nil
}

func signInputs(outputs []*entity.Output, privateKey *crypto.PrivateKey) func(res interface{}, input *entity.UnsignedInput) interface{} {
	return func(res interface{}, input *entity.UnsignedInput) interface{} {
		signature := tx.GenerateSignature(input, outputs, privateKey, encoding.NewTransactionGobEncoder())
		return append(res.([]*entity.SignedInput), entity.NewSignedInput(input, signature))
	}
}

func collectOutputs(res interface{}, output *entity.Output) interface{} {
	outputs := res.([]*entity.Output)
	output.Index = uint(len(outputs))
	return append(outputs, output)
}
