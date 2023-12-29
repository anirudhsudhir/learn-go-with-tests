package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		checkError(t, err)
		checkBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficicent funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(1000))
		checkInvalidBalance(t, err, InvalidFunds)
		checkBalance(t, wallet, Bitcoin(10))
	})
}

func checkBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func checkError(t testing.TB, err error) {
	if err != nil {
		t.Error("expected no error but received one")
	}
}

func checkInvalidBalance(t testing.TB, err, want error) {
	if err == nil {
		t.Fatal("wanted error but did not get one")
	}
	if err != want {
		t.Errorf("got %s, want %s", err.Error(), want)
	}
}
