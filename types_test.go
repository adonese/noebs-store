package store

import (
	"reflect"
	"testing"

	"github.com/adonese/noebs/ebs_fields"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Test_createAll(t *testing.T) {

	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	defer db.Close()
	tests := []struct {
		name string
		args Transaction
		db   *gorm.DB
		want bool
	}{
		{"test create tables", Transaction{
			UserID:     1,
			TerminalID: 1,
			Amount:     0.0,
			Successful: false,
			ServiceID:  2,
		}, db, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.createAll(db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createAllUser(t *testing.T) {

	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	defer db.Close()
	tests := []struct {
		name string
		args User
		db   *gorm.DB
		want bool
	}{
		{"test create tables", User{
			Transactions: []Transaction{{
				TerminalID: 1,
				Amount:     0.0,
				Successful: false,
				ServiceID:  4,
				UserID:     3,
			}},
		}, db, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.createAllUser(db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetTranSummary(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	type fields struct {
		Model        gorm.Model
		Transactions []Transaction
		Cards        []Card
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Summary
	}{
		{"successful test", fields{Model: gorm.Model{ID: 1}}, args{db: db}, Summary{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Model:        tt.fields.Model,
				Transactions: tt.fields.Transactions,
				Cards:        tt.fields.Cards,
			}
			if got := u.GetTranSummary(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetTranSummary() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

//GetFailedAmount
func TestUser_GetFailedAmount(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	type fields struct {
		Model        gorm.Model
		Transactions []Transaction
		Cards        []Card
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Summary
	}{
		{"successful test", fields{Model: gorm.Model{ID: 1}}, args{db: db}, Summary{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Model:        tt.fields.Model,
				Transactions: tt.fields.Transactions,
				Cards:        tt.fields.Cards,
			}
			if got := u.GetFailedAmount(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetFailedAmount() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

// GetFailedCount
func TestUser_GetFailedCount(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	type fields struct {
		Model        gorm.Model
		Transactions []Transaction
		Cards        []Card
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Summary
	}{
		{"successful test", fields{Model: gorm.Model{ID: 1}}, args{db: db}, Summary{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Model:        tt.fields.Model,
				Transactions: tt.fields.Transactions,
				Cards:        tt.fields.Cards,
			}
			if got := u.GetFailedCount(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetFailedCount() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestUser_GetMostUsedService(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	type fields struct {
		Model        gorm.Model
		Transactions []Transaction
		Cards        []Card
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Summary
	}{
		{"successful test", fields{Model: gorm.Model{ID: 1}}, args{db: db}, Summary{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Model:        tt.fields.Model,
				Transactions: tt.fields.Transactions,
				Cards:        tt.fields.Cards,
			}
			if got := u.GetMostUsedService(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetTranSummary() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestUser_GetSpending(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	type fields struct {
		Model        gorm.Model
		Transactions []Transaction
		Cards        []Card
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Summary
	}{
		{"successful test", fields{Model: gorm.Model{ID: 1}}, args{db: db}, Summary{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Model:        tt.fields.Model,
				Transactions: tt.fields.Transactions,
				Cards:        tt.fields.Cards,
			}
			if got := u.GetSpending(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetTranSummary() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestTransactionType_fill(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")
	type fields struct {
		Model gorm.Model
		Name  string
	}

	tests := []struct {
		name   string
		fields fields
		args   *gorm.DB
	}{
		{"successful ", fields{}, db},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ty := &TransactionType{
				Model: tt.fields.Model,
				Name:  tt.fields.Name,
			}
			ty.fill(tt.args)
		})
	}
}

func TestUser_New(t *testing.T) {
	type fields struct {
		Model        gorm.Model
		Transactions []Transaction
		Cards        []Card
	}
	type args struct {
		db       *gorm.DB
		username string
		email    string
		mobile   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Model:        tt.fields.Model,
				Transactions: tt.fields.Transactions,
				Cards:        tt.fields.Cards,
			}
			u.New(tt.args.db, tt.args.username, tt.args.email, tt.args.mobile)
		})
	}
}

func TestTerminal_NewMerchant(t *testing.T) {
	u := &User{
		Name:   "Galal",
		Mobile: "0912141679",
	}

	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")

	type fields struct {
		Model      gorm.Model
		TerminalID string
		Merchant   User
	}
	type args struct {
		u  *User
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Create new merchant", fields{TerminalID: "12345670"}, args{u: u, db: db}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminal := &Terminal{
				Model:          tt.fields.Model,
				TerminalNumber: tt.fields.TerminalID,
				Merchant:       tt.fields.Merchant,
			}
			if err := terminal.NewMerchant(tt.args.u, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Terminal.NewMerchant() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Errorf("Terminal.NewMerchant() error = %v, wantErr %v", err, tt.wantErr)

			}
		})
	}
}

func TestTerminal_getTerminal(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")

	type fields struct {
		Model       gorm.Model
		TerminalID  string
		Merchant    User
		Transaction []Transaction
	}
	type args struct {
		name string
		db   *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"get terminal with ID", fields{}, args{db: db, name: "12345678"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminal := &Terminal{
				Model:          tt.fields.Model,
				TerminalNumber: tt.fields.TerminalID,
				Merchant:       tt.fields.Merchant,
				Transaction:    tt.fields.Transaction,
			}
			if err := terminal.getTerminal(tt.args.name, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Terminal.getTerminal() error = %v, wantErr %v\nTerminal object is: %#v", err, tt.wantErr, terminal)

			}
		})
	}
}

func TestTransaction_Create(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testt.db")

	ebs := &ebs_fields.GenericEBSResponseFields{
		TerminalID:     "12345670",
		ClientID:       "ACTS",
		PAN:            "1234",
		ServiceID:      "purchase",
		TranAmount:     40,
		ToCard:         "123456789",
		EBSServiceName: "purchase",
	}

	type args struct {
		ebs  *ebs_fields.GenericEBSResponseFields
		name string
		db   *gorm.DB
	}
	transaction := &Transaction{}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"successful create", args{db: db, ebs: ebs, name: "purchase"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := transaction.Create(tt.args.ebs, tt.args.name, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.Create() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Errorf("Transaction.Create() error = %v, wantErr %v", err, tt.wantErr)

			}
		})
	}
}

func getDB(name string) *gorm.DB {
	db, _ := getEngine(name)
	return db
}
