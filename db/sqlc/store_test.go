package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before: ", account1.Balance, account2.Balance)

	// run n concurrent transfer transaction

	n := 10
	amount := int64(10)

	errs := make(chan error)
	for i := 0; i < n; i++ {
		i := i
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			fromAccountID := account1.ID
			toAccountID := account2.ID

			if i%2 == 1 {
				fromAccountID = account2.ID
				toAccountID = account1.ID
			}
			ctx := context.WithValue(context.Background(), txKey, txName)
			_, err := store.TransferTx(ctx, TransferTxParagrams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Account:       amount,
			})
			errs <- err
		}()
	}

	// check result
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balance
	updateAccount1, err := testqueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	updateAccount2, err := testqueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after: ", updateAccount1.Balance, updateAccount2.Balance)
	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)

}
