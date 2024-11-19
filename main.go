package main

import (
	"github.com/supakjack/isekai-shop-api-tutorial/config"
	"github.com/supakjack/isekai-shop-api-tutorial/databases"
	"github.com/supakjack/isekai-shop-api-tutorial/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())
	server.Start()
}
