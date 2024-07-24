package main

import (
	"fmt"
	"ormer/internals/orm"
	"ormer/utils"
)

type Users struct {
	ID   uint   `orm:"primary_key"`
	name string `orm:"name"`
}

type Products struct {
	ID   uint   `orm:"primary_key"`
	name string `orm:"name"`
}

func main() {
	orm, err := orm.NewOrm(utils.Postgres, "")
	if err != nil {
		fmt.Printf("failed creating new orm: %s", err.Error())
	}
	defer orm.DB.Close()

	userOne := Users{
		ID:   0,
		name: "hector",
	}

	orm.CreateNewTable(userOne)
	// rows, err := orm.NewSelectStatement(userOne , make([]string, 0))

}
