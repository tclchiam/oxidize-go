package blockchain

import (
	"fmt"
	"github.com/tclchiam/block_n_go/tx"
	"github.com/tclchiam/block_n_go/wallet"
)

func (bc *Blockchain) buildExpenseTransaction(sender, receiver *wallet.Wallet, expense uint) (*tx.Transaction, error) {
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
	takeMinimumToMeetExpense := func(_ *tx.Transaction, output *tx.Output) bool {
		take := liquidBalance < expense
		if take {
			liquidBalance += output.Value
		}
		return take
	}

	buildInputs := func(res interface{}, transaction *tx.Transaction, output *tx.Output) interface{} {
		input := tx.NewUnsignedInput(transaction.ID, output, sender.PublicKey)
		return res.(tx.UnsignedInputs).Add(input)
	}

	inputs := unspentOutputs.
		Filter(takeMinimumToMeetExpense).
		Reduce(tx.EmptyUnsignedInputs(nil), buildInputs).(tx.UnsignedInputs)

	outputs := tx.EmptyOutputs().
		Add(tx.NewOutput(expense, receiver.GetAddress()))

	if liquidBalance-expense > 0 {
		outputs = outputs.Add(tx.NewOutput(liquidBalance-expense, senderAddress))
	}

	return tx.NewTx(&inputs, &outputs, sender.PrivateKey), nil
}
