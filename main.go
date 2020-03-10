package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/itslearninggermany/itswizard_azureactivedirctory"
	"github.com/itslearninggermany/itswizard_basic"
	"github.com/itslearninggermany/itswizard_sync"
	"github.com/jinzhu/gorm"
)

func main() {
	database, _ := gorm.Open("mysql", "admin"+":"+"pat5rxe4"+"@tcp("+"testserver.clfw49bkjzgd.eu-west-3.rds.amazonaws.com"+")/"+"Client"+"?charset=utf8&parseTime=True&loc=Local")

	database.AutoMigrate(&itswizard_sync.DbGroup15{})
	database.AutoMigrate(&itswizard_sync.DbPerson15{})
	database.AutoMigrate(&itswizard_sync.DbGroupMembership15{})
	database.AutoMigrate(&itswizard_sync.DbMentorStudentRelationship15{})
	database.AutoMigrate(&itswizard_sync.DbStudentParentRelationship15{})
	database.AutoMigrate(&itswizard_sync.DbSyncCache15{})
	database.AutoMigrate(&itswizard_azureactivedirctory.Aad{})
	database.AutoMigrate(&itswizard_basic.DbAAdLog{})
	fmt.Print("as")
}


type toDo struct {
	db     *gorm.DB
	idList []uint
	aads   []*itswizard_azureactivedirctory.Aad
}

/*
creates a new job
*/
func NewToDo(database *gorm.DB) (td *toDo) {
	a := new(toDo)
	a.db = database
	return a
}

/*
Get All OrganisationIds who use Aad
*/
func (p *toDo) updateIdList() error{
	fmt.Println("hier")
	var allAads []itswizard_azureactivedirctory.Aad
	p.db.Find(&allAads)
	fmt.Println(allAads[0].OrganisationID)
	notExist := true
	for in := 0; in < len(allAads); in++ {
		for i := 0; i < len(p.idList); i++ {
			if p.idList[i] == allAads[in].OrganisationID {
				notExist = false
				break
			}
		}
		if notExist {
			fmt.Println("wwefwef")
			p.idList = append(p.idList, allAads[in].OrganisationID)
			aadservice, err := itswizard_azureactivedirctory.NewAad(p.db, allAads[in].OrganisationID)
			if err != nil {
				return err
			}
			p.aads = append(p.aads, aadservice)

		}
		fmt.Println(p.idList)
		fmt.Println(p.aads)
	}
	return nil
}


func (p *toDo) Sync() {
	fmt.Println("sync")
	for i := 0; i < len(p.aads); i++ {
		fmt.Println("started ", p.aads[i].OrganisationID)
		go p.aads[i].Sync()
	}
}
