// walletFacade.go: Facade
package main

import (
	"fmt"
	"log"
)

// Step 2 + 3: Khai báo và triển khai giao diện đơn giản trong một class Facade
type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	// Facade chịu trách nhiệm khởi tạo các phần hệ thống con
	walletFacade := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "debit", amount)
	return nil
}

// Step 4: Nếu Facade quá lớn, có thể tách logic này ra thành các Facade nhỏ hơn cho từng chức năng.
// Ví dụ: một Facade riêng cho thông báo hoặc ghi sổ cái.

// account.go: Complex subsystem parts
type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{name: accountName}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

// securityCode.go: Complex subsystem parts
type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

// wallet.go: Complex subsystem parts
type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance -= amount
	return nil
}

// ledger.go: Complex subsystem parts
type Ledger struct{}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}

// notification.go: Complex subsystem parts
type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

// main.go: Client code
func main() {
	fmt.Println()
	// Step 3: Khách hàng giao tiếp với hệ thống chỉ thông qua Facade
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}