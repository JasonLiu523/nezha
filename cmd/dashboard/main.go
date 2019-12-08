package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/patrickmn/go-cache"

	"github.com/p14yground/nezha/cmd/dashboard/controller"
	"github.com/p14yground/nezha/cmd/dashboard/rpc"
	"github.com/p14yground/nezha/model"
	"github.com/p14yground/nezha/service/dao"
)

func init() {
	var err error
	dao.Conf, err = model.ReadInConfig("data/config.yaml")
	if err != nil {
		panic(err)
	}
	dao.Admin = &model.User{
		Login: dao.Conf.GitHub.Admin,
	}
	dao.DB, err = gorm.Open("sqlite3", "data/sqlite.db")
	if err != nil {
		panic(err)
	}
	dao.Cache = cache.New(5*time.Minute, 10*time.Minute)
}

func main() {
	go controller.ServeWeb()
	go rpc.ServeRPC()
	select {}
}