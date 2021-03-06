package store

import (
	"encoding/json"
	"log"

	"github.com/adonese/noebs/ebs_fields"
	"github.com/jinzhu/gorm"
)

// User List transactions associated to this user
type User struct {
	gorm.Model
	Transactions []Transaction
	Cards        []Card
	Name         string
	Mobile       string
	Username     string
}

func (u *User) createAllUser(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}, &Card{}).Error; err != nil {
		log.Printf("Error in db.AutoMigrate: %v", err)
		return err
	}
	if err := db.Create(u).Error; err != nil {
		log.Printf("Error in db.AutoMigrate: %v", err)
		return err
	}
	return nil
}

/*
- getAll()
- getFailedCount()
- getSucceededCount()
- getFailedAmount()
- getSucceededAmount()
- getMostUsedService()
- getLeastUsedService()
- getTotalSpending()
- getCards()
- getMobile()
*/
func (u *User) GetProfile(db *gorm.DB) User {
	db.Find(&u)
	return *u
}

func (u *User) New(db *gorm.DB, username, email, mobile string) {
	db.Where("username = ? OR email = ? or mobile = ?", username, email, mobile).Find(u)
}

//GetFailedCount reports the number of failed transactions for user
func (u *User) GetFailedCount(db *gorm.DB) int {
	id := u.ID
	var count int
	db.Raw("select count(*) from transactions where user_id = ? AND successful = 0", &id).Count(&count)
	return count
}

func (u *User) GetSucceededCount(db *gorm.DB) int {
	id := u.ID
	var count int
	db.Exec("select count(*) from transactions where user_id = 1 AND successful = 0", &id).Find(&count)
	return count
}

//GetFailedAmount responds with the failed services
func (u *User) GetFailedAmount(db *gorm.DB) []Summary {
	id := u.ID
	var summary []Summary
	db.Raw(`select tt.name, t.service_id, sum(t.amount) as amount, count(*) as count from 
	transactions t
	inner join users u on u.ID = t.user_id
	JOIN transaction_types tt
	where t.user_id = ? AND t.successful = 0 AND tt.id = service_id
	group by t.service_id
	`, &id).Scan(&summary)
	return summary
}

//GetSpending returns the sum of spending of this user
func (u *User) GetSpending(db *gorm.DB) float32 {
	id := u.ID
	var count float32
	db.Raw("select sum(amount) from transactions where user_id = ? AND successful = 1", &id).Count(&count)
	return count
}

//GetMostUsedService returns a list of most used services
func (u *User) GetMostUsedService(db *gorm.DB) []Summary {
	id := u.ID
	var summary []Summary

	db.Raw(`select tt.name, t.service_id, count(*) as count from 
	transactions t
	inner join users u on u.ID = t.user_id
	JOIN transaction_types tt
	where t.user_id = 1 AND t.successful = 1 AND tt.id = service_id
	group by t.service_id
	order by count desc
	`, &id).Find(&summary)
	return summary
}

//GetTranSummary returns a summary of transactions
func (u *User) GetTranSummary(db *gorm.DB) []Summary {
	/*
	 name id count sum_amount

	*/
	id := u.ID
	var summary []Summary

	if err := db.Table("users").Raw(`select tt.name, t.service_id, sum(t.amount) as amount, count(*) as count from 
	transactions t
	inner join users u on u.ID = t.user_id
	JOIN transaction_types tt
	where t.user_id = 1 AND t.successful = 1 AND tt.id = service_id
	group by t.service_id
	`, &id).Scan(&summary).Error; err != nil {
		log.Printf("Error in GetTranSummary: %v", err)
	}
	return summary
}

//GetCards returns cards associated to this card holder
func (u *User) GetCards(db *gorm.DB) []Card {

	var cards []Card
	id := u.ID
	db.Exec(`select * from cards where user_id = ?`, id).Find(&cards)
	return cards
}

//GetMobiles returns mobile numbers associated to this card holder
func (u *User) GetMobiles(db *gorm.DB) []Mobile {
	var mobiles []Mobile
	id := u.ID

	db.Exec(`select * from mobiles where user_id = ?`, id).Find(&mobiles)
	return mobiles
}

//Transaction table
type Transaction struct {
	gorm.Model
	TerminalID    uint
	SourceID      uint
	DestinationID uint
	// DestinationID   int
	Amount     float32
	Successful bool
	ServiceID  uint
	// TransactionID   int
	UserID uint
	// SourceID uint
}

//NewTransaction returns a new Transaction objects
func NewTransaction() *Transaction {
	return &Transaction{}
}

//Marshal transaction to json
func (t *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

//Create struct Transaction with the transactions
func (t *Transaction) Create(ebs *ebs_fields.GenericEBSResponseFields, name string, db *gorm.DB) error {

	if ebs.PAN != "" {
		card := &Card{PAN: ebs.PAN}
		db.FirstOrCreate(card)
	}
	if ebs.FromAccount != "" {
		card := &Card{AccountNumber: ebs.FromAccount}
		db.FirstOrCreate(card)
	}
	if ebs.ToAccount != "" {
		card := &Card{AccountNumber: ebs.ToAccount}
		db.FirstOrCreate(card)
	}
	if ebs.TerminalID != "" {
		terminal := &Terminal{TerminalNumber: ebs.TerminalID}
		db.FirstOrCreate(terminal)
	}

	// log.Printf("tid = %v, pan = %v\n", )
	if err := db.Exec(`insert into transactions(amount, created_at, source_id, terminal_id)
			select ?, datetime('now', 'localtime'), s.id,  t.id  FROM terminals t
			inner join sources s
			where t.terminal_number = ? AND s.pan = ?`, ebs.TranAmount, ebs.TerminalID, ebs.PAN).Error; err != nil {
		log.Printf("Error in *Transaction.Create: %v", err)
		return err
	}
	return nil
}

func (t *Transaction) createAll(db *gorm.DB) error {
	// db.AutoMigrate(&User{})
	db.AutoMigrate(&TransactionType{}, &User{}, &Terminal{}, &Card{})

	if err := db.AutoMigrate(&Transaction{}).Error; err != nil {
		log.Printf("Error in AutoMigrate: Error: %v", err)
		return err
	}
	if err := db.Create(t).Error; err != nil {
		log.Printf("Error in db.Create: Error: %v", err)
		return err
	}

	return nil

}

//Commit the tables with their association to DB
func (t *Transaction) Commit(db *gorm.DB) error {
	return nil
}

// UserProfile card holder info + their associated mobile numbers
type UserProfile struct {
	Cards   []Card
	Mobiles []Mobile
}

// Card table
type Card struct {
	gorm.Model
	PAN           string `gorm:"unique_index"`
	ExpDate       string
	UserID        uint
	Transaction   []Transaction
	AccountNumber string `gorm:"unique_index"`
}

func (c *Card) migrate(db *gorm.DB) {
	// db.Exec(``)
}
func (c *Card) topCards(db *gorm.DB) []Card {
	var res []Card
	db.Raw(`select * from cards c
	join transactions t on t.source_id = c.id
	where c.pan = ?`, c.PAN).Scan(&res)
	return res
}

// Terminal is related to POS merchants
type Terminal struct {
	gorm.Model
	TerminalNumber string `gorm:"unique_index"`
	Merchant       User
	Transaction    []Transaction
}

//Init initializes a new terminal id with all of the features packed
func (t *Terminal) Init(id string) {
	t.TerminalNumber = id

}

//NewMerchant generates a new merchant in noebs system
func (t *Terminal) NewMerchant(u *User, db *gorm.DB) error {
	db.AutoMigrate(&Terminal{})
	t.Merchant = *u
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil

}

func (t *Terminal) getTerminal(name string, db *gorm.DB) error {
	t.TerminalNumber = name
	if err := db.Where(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Terminal) getTransactions(db *gorm.DB) []Transaction {
	var tt []Transaction
	if err := db.Raw(`select * from transactions t
		join terminals tt on tt.id = t.terminal_id
		where tt.terminal_number = ?`, t.TerminalNumber).Scan(&tt).Error; err != nil {
		log.Printf("Error in &Terminal.getTransactions: %v", err)
	}
	return tt
}

func (t *Terminal) getMostUsedService(db *gorm.DB) []servicesCount {
	var res []servicesCount

	db.Raw(`select tt.name as name, t.service_id as id, count(t.successful) - sum(t.successful) as failed, sum(t.successful) as succeeded, sum(t.amount) as amount, count(*) as count from transactions t
	join transaction_types tt on tt.id = t.service_id
	join terminals ts on ts.ID = t.terminal_id
	where ts.terminal_number = ?
	group by t.service_id
	order by t.amount`, t.TerminalNumber).Scan(&res)
	return res
}

type servicesCount struct {
	Count     int     // counts how many transactions
	Amount    float32 // transactions total amount (failed / succeeded)
	Failed    int
	Succeeded int
	Name      string
	ID        int // services id from transaction type
}

type purchaseCount struct {
	Count     int
	Amount    float32
	Failed    int
	Succeeded int
}

type p2pCount struct {
	Count     int
	Amount    float32
	Failed    int
	Succeeded int
}

//Account info
type Account struct {
	AccountNumber string
	AccountName   string
}

//Mobile Table
type Mobile struct {
	Number string
	Operator
}

// Operator table
type Operator struct {
	Sudani  int
	Zain    int
	MTN     int
	Unknown int
}

//TransactionType the transactions we support at noebs
type TransactionType struct {
	// P2p           bool
	// ZainTopUp     bool
	// SudaniTopUp   bool
	// MTNTopUp      bool
	// Electricity   bool
	// Account       bool
	// ZainBill      bool
	// MTNBill       bool
	// SudaniBill    bool
	// Purchase      bool
	// TransactionID uint
	gorm.Model
	Name string `gorm:"unique_index"`
}

func (tt *TransactionType) fill(db *gorm.DB) {

	db.AutoMigrate(&tt)
	//tt.Purchase = true
	t := []TransactionType{TransactionType{Name: "Purchase"},
		TransactionType{Name: "Balance"},
		TransactionType{Name: "MTN Top Up"},
		TransactionType{Name: "Zain Top Up"},
		TransactionType{Name: "Sudani Top Up"},
		TransactionType{Name: "MTN Bills"},
		TransactionType{Name: "Sudani Bills"},
		TransactionType{Name: "Zain Bills"},
		TransactionType{Name: "Electricity"},
		TransactionType{Name: "Card Transfer"}}

	db.Create(&t[0])
	db.Create(&t[1])
	db.Create(&t[2])
	db.Create(&t[3])
	db.Create(&t[4])
	db.Create(&t[5])
	db.Create(&t[6])
	db.Create(&t[7])
	db.Create(&t[8])
	db.Create(&t[9])
	// db.Create(t[5])

}

//Payer list all of the system payers
type Payer struct {
	P2p         string
	ZainTopUp   string
	SudaniTopUp string
	MTNTopUp    string
	Electricity string
	Account     string
	ZainBill    string
	MTNBill     string
	SudaniBill  string
}

//Summary summarizes transacations for a specific card holder
type Summary struct {
	Name   string  `json:"name,omitempty"`
	ID     int     `json:"id,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	Count  float32 `json:"count,omitempty"`
}
