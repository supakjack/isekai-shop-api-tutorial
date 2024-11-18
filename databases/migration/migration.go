package main

import (
	"fmt"

	"github.com/supakjack/isekai-shop-api-tutorial/config"
	"github.com/supakjack/isekai-shop-api-tutorial/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
