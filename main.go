package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/smartgolang/gosmartsqlite"
)

func main() {
	db, _ := sql.Open("sqlite3", "./my.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, email TEXT)")
	statement.Exec()

	if len(os.Args) == 1 {
		fmt.Printf("\nUsage:\n")

		fmt.Println(os.Args[0], "-Upsert"+" Constantine Vassil Constantine@Vassil.com", "")
		fmt.Println(os.Args[0], "-Search"+" Constantine", "")
		fmt.Println(os.Args[0], "-Delete"+"  Constantine", "")
		fmt.Println(os.Args[0], "-List"+"  ", "")

	} else {
		service := ""

		fmt.Printf("%s|%d\n", "Args", len(os.Args))
		if len(os.Args) >= 1 {
			service = os.Args[1]
			service = strings.Replace(service, " ", "", -1)
			switch service {
			case "-Upsert": // ...
				if len(os.Args) == 5 {
					FirstName := os.Args[2]
					LastName := os.Args[3]
					Email := os.Args[4]

					var newPerson gosmartsqlite.Person
					newPerson.FirstName = FirstName
					newPerson.LastName = LastName
					newPerson.Email = Email
					gosmartsqlite.UpsertPerson(db, newPerson)
				}
			case "-Search": // ...
				if len(os.Args) == 3 {
					FirstName := os.Args[2]
					people, _ := gosmartsqlite.SearchForPerson(db, FirstName)
					for _, ourPerson := range people {
						fmt.Printf("\n----\nID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s", ourPerson.ID, ourPerson.FirstName, ourPerson.LastName, ourPerson.Email)
					}
				}
			case "-Delete": // ...
				if len(os.Args) == 3 {
					FirstName := os.Args[2]
					people, _ := gosmartsqlite.SearchForPerson(db, FirstName)
					for _, ourPerson := range people {
						gosmartsqlite.DeletePerson(db, ourPerson.ID)
					}
				}
			case "-List": // ...
				if len(os.Args) == 2 {

					people := gosmartsqlite.List(db)
					for _, ourPerson := range people {
						fmt.Printf("\n----\nID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s", ourPerson.ID, ourPerson.FirstName, ourPerson.LastName, ourPerson.Email)
					}
				}
			default:
				panic(fmt.Sprintf("invalid service %q", service)) //
			}
		}
	}

}
