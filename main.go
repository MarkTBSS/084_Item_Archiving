package main

import (
	"github.com/MarkTBSS/084_Item_Archiving/config"
	"github.com/MarkTBSS/084_Item_Archiving/databases"
	"github.com/MarkTBSS/084_Item_Archiving/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
