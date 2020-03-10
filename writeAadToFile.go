package main

import (
	"fmt"
	"github.com/itslearninggermany/itswizard_azureactivedirctory"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func main () {
	database, _ := gorm.Open("mysql", "admin"+":"+"pat5rxe4"+"@tcp("+"testserver.clfw49bkjzgd.eu-west-3.rds.amazonaws.com"+")/"+"Client"+"?charset=utf8&parseTime=True&loc=Local")
	var allAads []itswizard_azureactivedirctory.Aad
	database.Find(&allAads)
	var out string
	for i := 0; i < len (allAads); i++ {
		if i == 0 {
			out = fmt.Sprint(allAads[i].OrganisationID)
		} else {
			out = fmt.Sprint(out,"\n",allAads[i].OrganisationID)
		}
	}
	err := ioutil.WriteFile("./aads.txt",[]byte(out),666)
	if err != nil {
		fmt.Println(err)
	}
}
