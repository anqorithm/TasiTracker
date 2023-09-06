package jobs

import (
	"fmt"
	"time"

	integrations "github.com/qahta0/tasitracker/integration/sahmi"
	"gorm.io/gorm"
)

func FetchAndStoreStocks(dbConnection *gorm.DB) {
	fmt.Println("Started: ", time.Now())
	comapniens, err := integrations.FetchStocks()
	if err != nil {
		panic(err)
	}
	integrations.SaveStocks(&comapniens, dbConnection)
}
