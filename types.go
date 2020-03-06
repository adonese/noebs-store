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

func (u *User) GetFailedCount(db *gorm.DB) int {
	id := u.ID
	var count int
	db.Exec("select count(*) from transactions where user_id = 1 AND successful = 0", &id).Find(&count)
	return count
}

func (u *User) GetSucceededCount(db *gorm.DB) int {
	id := u.ID
	var count int
	db.Exec("select count(*) from transactions where user_id = 1 AND successful = 0", &id).Find(&count)
	return count
}

func (u *User) GetFailedAmount(db *gorm.DB) int {
	id := u.ID
	var count int
	db.Exec("select sum(amount) from transactions where user_id = 1 AND successful = 0", &id).Find(&count)
	return count
}

//GetSpending returns the sum of spending of this user
func (u *User) GetSpending(db *gorm.DB) int {
	id := u.ID
	var count int
	db.Exec("select sum(amount) from transactions where user_id = 1 AND successful = 1", &id).Find(&count)
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
	// TransactionID int
	Source      Source
	Destination Destination

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

//Populate struct Transaction with the transactions
func (t *Transaction) Populate(ebs *ebs_fields.GenericEBSResponseFields, name string) error {
	id := toID(name)
	t.Destination.fill(ebs)
	// t.Source.fill(ebs)
	// t.TransactionType.fill(id)
	t.ServiceID = uint(id)
	t.Amount = ebs.TranAmount
	return nil
}

func (t *Transaction) createAll(db *gorm.DB) error {
	// db.AutoMigrate(&User{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Source{})
	db.AutoMigrate(&Destination{})
	db.AutoMigrate(&TransactionType{}, &User{})

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

//Create commits result to database. It assumed a populated struct, Transaction.
func (t *Transaction) Create(db *gorm.DB) error {
	// db.AutoMigrate(&User{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Source{})
	db.AutoMigrate(&Destination{})
	db.AutoMigrate(&TransactionType{})

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

// Source is the transaction source. It can be initiated from a user
// or an account. It can happen in any place within our network.
type Source struct {
	gorm.Model
	Card
	CardID int
	Account
	AccountID int
	// UserID        int
	// Transactions []Transaction
	TransactionID uint
}

func (s *Source) fill(ebs *ebs_fields.GenericEBSResponseFields) {
	if ebs.FromAccount != "" {
		// fill account
		s.Account.AccountNumber = ebs.FromAccount
		s.AccountName = "" //FIXME: we need to get the account name
		return
	}
	s.PAN = ebs.PAN
	return
}

// Destination is the transaction source. It can be initiated from a user
// or an account. It can happen in any place within our network.
type Destination struct {
	ToCard        string
	ToAccount     string
	ToAccountName string
	TransactionID uint
}

//fill the destination struct with either account / card data
func (d *Destination) fill(ebs *ebs_fields.GenericEBSResponseFields) {
	if ebs.ToAccount != "" {
		// fill account
		d.ToAccount = ebs.ToAccount
		d.ToAccountName = "" //FIXME: we need to get the account name
		return
	}
	d.ToCard = ebs.ToCard
	return
}

// User card holder info + their associated mobile numbers
type UserProfile struct {
	Cards   []Card
	Mobiles []Mobile
}

// Card table
type Card struct {
	PAN     string
	ExpDate string
	UserID  uint
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
