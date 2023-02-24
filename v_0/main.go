package main

import "github.com/set2002satoshi/my-site-api/infrastructure"

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	db.DBInit()
	r.Run()
}
