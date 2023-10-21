package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s but wanted %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("expected error but didn't got one")
		}

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{balance: 20}
		w.Withdraw(Bitcoin(5))
		assertBalance(t, w, Bitcoin(15))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingFunds := Bitcoin(20)
		w := Wallet{startingFunds}
		err := w.Withdraw(Bitcoin(100))

		assertError(t, err, ErrNotEnoughFunds)
		assertBalance(t, w, startingFunds)
	})
}

func TestBitcoin(t *testing.T) {
	got := fmt.Sprintf("%s", Bitcoin(10))
	want := "10 BTC"

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}
