package account

import (
    "sync"
)

type Account struct {
    isOpen bool
    balance int64
    accountLock sync.Mutex
}

func Open(initialDeposit int64) *Account {
    if initialDeposit < 0 {
        return nil
    }
    return &Account {
        isOpen: true,
        balance: initialDeposit,
    }
}

func (a *Account) Close() (payout int64, ok bool) {
    a.accountLock.Lock()
    defer a.accountLock.Unlock()

    // explicit
    payout = 0
    ok = false

    if a.isOpen {
        payout = a.balance
        ok = true

        a.balance = 0
        a.isOpen = false
    }

    return payout, ok
}

func (a *Account) Balance() (balance int64, ok bool) {
    a.accountLock.Lock()
    defer a.accountLock.Unlock()

    return a.balance, a.isOpen
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
    a.accountLock.Lock()
    defer a.accountLock.Unlock()

    // explicit
    newBalance = a.balance
    ok = false

    if a.isOpen {
        if tempBalance := a.balance + amount; tempBalance >= 0 {
            a.balance = tempBalance
            newBalance = a.balance
            ok = true
        }
    }

    return newBalance, ok
}
