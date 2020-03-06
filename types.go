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

//Transaction table
type Transaction struct {
	gorm.Model
	// TransactionID int
	Source      Source
	Destination Destination

	// DestinationID   int
	Amount          float32
	Successful      bool
	TransactionType TransactionType
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
	t.TransactionType.fill(id)
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
	P2p           bool
	ZainTopUp     bool
	SudaniTopUp   bool
	MTNTopUp      bool
	Electricity   bool
	Account       bool
	ZainBill      bool
	MTNBill       bool
	SudaniBill    bool
	Purchase      bool
	TransactionID uint
}

func (tt *TransactionType) fill(id int) {
	tt.Purchase = true
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
