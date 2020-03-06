package store

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Test_createAll(t *testing.T) {

	db := getDB("/mnt/c/Users/MSI/Documents/test000.db")
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
			TransactionType: TransactionType{
				P2p:         false,
				ZainTopUp:   false,
				SudaniTopUp: false,
				MTNTopUp:    false,
				Electricity: false,
				Account:     false,
				ZainBill:    false,
				MTNBill:     false,
				SudaniBill:  false,
				Purchase:    true,
			},
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

	db := getDB("/mnt/c/Users/MSI/Documents/test2.db")
	defer db.Close()
	tests := []struct {
		name string
		args User
		db   *gorm.DB
		want bool
	}{
		{"test create tables", User{
			Transactions: []Transaction{{

				// Source: Source{
				// 	Card: Card{
				// 		PAN:     "111111111111111",
				// 		ExpDate: "",
				// 	},
				// 	CardID: 0,
				// 	Account: Account{
				// 		AccountNumber: "",
				// 		AccountName:   "",
				// 	},
				// 	AccountID: 0,
				// },
				Destination: Destination{
					ToCard:        "2222222222222",
					ToAccount:     "",
					ToAccountName: "",
					TransactionID: 0,
				},
				Amount:     0.0,
				Successful: false,
				TransactionType: TransactionType{
					P2p:           false,
					ZainTopUp:     false,
					SudaniTopUp:   false,
					MTNTopUp:      false,
					Electricity:   false,
					Account:       false,
					ZainBill:      false,
					MTNBill:       false,
					SudaniBill:    false,
					Purchase:      false,
					TransactionID: 0,
				},
				UserID: 0,
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
