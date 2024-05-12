package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("invalid withdraw", func(t *testing.T) {
		startingBal := Bitcoin(20)
		wallet := Wallet{startingBal}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBal)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
