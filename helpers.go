package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//TODO complete this function and add the rest of the cases
func toID(name string) int {
	switch name {
	case "workingKey":
		return 1
	case "purchase":
		return 2
	case "cardToCard":
		return 3
	case "balance":
		return 4
	case "consumerCardTransfer":
		return 10
	case "consumerZainTopUp":
		return 11
	case "consumerMtnTopUp":
		return 12
	default:
		return -99
	}

}

func getDB(name string) *gorm.DB {
	db, _ := gorm.Open("sqlite3", name)
	return db
}
