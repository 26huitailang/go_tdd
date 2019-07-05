package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, 10)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoint(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoint(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoint) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Error("wanted an error but didnt get one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
