package main

import (
	"fmt"
	"github.com/itslearninggermany/itswizard_azureactivedirctory"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, _ := gorm.Open("mysql", "admin"+":"+"pat5rxe4"+"@tcp("+"testserver.clfw49bkjzgd.eu-west-3.rds.amazonaws.com"+")/"+"Client"+"?charset=utf8&parseTime=True&loc=Local")
	fmt.Print("sadasd")
	argsWithProg := os.Args
	if len (argsWithProg) > 1 {
		zahl, err :=strconv.Atoi(argsWithProg[1])
		if err != nil {
			fmt.Println(err)
		}
		aadservice, err := itswizard_azureactivedirctory.NewAad(database, uint(zahl))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(aadservice.Sync())
	}

}

