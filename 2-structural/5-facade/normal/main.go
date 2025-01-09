package main

import (
	"fmt"
	"log"
)

// Không sử dụng Facade Pattern

func main() {
	fmt.Println("Starting operations without Facade")

	// Khởi tạo từng thành phần riêng lẻ
	account := newAccount("abc")
	securityCode := newSecurityCode(1234)
	wallet := newWallet()
	notification := &Notification{}
	ledger := &Ledger{}

	fmt.Println("All components initialized")

	// Add money to wallet
	fmt.Println("Starting add money to wallet")
	err := account.checkAccount("abc")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	err = securityCode.checkCode(1234)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	wallet.creditBalance(10)
	notification.sendWalletCreditNotification()
	ledger.makeEntry("abc", "credit", 10)

	fmt.Println()

	// Deduct money from wallet
	fmt.Println("Starting debit money from wallet")

	err = account.checkAccount("abc")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = securityCode.checkCode(1234)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = wallet.debitBalance(5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	notification.sendWalletDebitNotification()
	ledger.makeEntry("abc", "debit", 5)
}

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

// Nhược điểm và khó khăn khi không sử dụng Facade Pattern:
// 1. Phức tạp hóa mã nguồn:
//    - Client phải xử lý trực tiếp với nhiều thành phần hệ thống con, làm tăng độ phức tạp và khó đọc mã.
//    - Mỗi lần thêm chức năng mới, client phải biết rõ từng thành phần liên quan và cách giao tiếp với chúng.
//
// 2. Dễ xảy ra lỗi:
//    - Các bước xử lý cần được thực hiện theo đúng thứ tự, nếu không sẽ gây lỗi. Không có lớp trung gian để đảm bảo thứ tự logic này.
//
// 3. Khó bảo trì:
//    - Khi thay đổi một phần của hệ thống con, tất cả các đoạn mã client liên quan cần được cập nhật để phù hợp.
//
// 4. Tái sử dụng hạn chế:
//    - Không có lớp Facade để đóng gói logic chung, dẫn đến việc viết lại mã lặp lại ở nhiều nơi nếu cần sử dụng cùng chức năng.
//
// 5. Đọc hiểu và quản lý kém:
//    - Đối với các hệ thống phức tạp hơn, việc phải làm việc trực tiếp với nhiều hệ thống con khiến mã trở nên khó hiểu và khó quản lý.
//
// Việc sử dụng Facade Pattern giúp khắc phục các vấn đề này bằng cách cung cấp một giao diện đơn giản và nhất quán để giao tiếp với toàn bộ hệ thống.
