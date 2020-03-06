package store

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Test_createAll(t *testing.T) {

	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
	defer db.Close()
	tests := []struct {
		name string
		args Transaction
		db   *gorm.DB
		want bool
	}{
		{"test create tables", Transaction{
			UserID: 1,
			Source: Source{
				Card: Card{
					PAN:     "32323232",
					ExpDate: "323232",
				},
				CardID: 2,
				Account: Account{
					AccountNumber: "",
					AccountName:   "",
				},
				AccountID: 0,
			},
			Destination: Destination{
				ToCard:        "9222222222",
				ToAccount:     "",
				ToAccountName: "",
			},
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

	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
	defer db.Close()
	tests := []struct {
		name string
		args User
		db   *gorm.DB
		want bool
	}{
		{"test create tables", User{
			Transactions: []Transaction{{
				Source: Source{
					Card: Card{
						PAN:     "111111111111111",
						ExpDate: "",
					},
					CardID: 0,
					Account: Account{
						AccountNumber: "",
						AccountName:   "",
					},
					AccountID: 0,
				},
				Destination: Destination{
					ToCard:        "2222222222222",
					ToAccount:     "",
					ToAccountName: "",
					TransactionID: 0,
				},
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
	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
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
	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
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

func TestUser_GetMostUsedService(t *testing.T) {
	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
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
	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
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
	db := getDB("/mnt/c/Users/MSI/Documents/testx.db")
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
