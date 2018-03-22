package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int
	name string
	email string
}

func (u *User) getFromRepository(db *sql.DB, id int)  {
        stmtOut, err := db.Prepare("SELECT id, name, email FROM by_customers WHERE id = ?")
        checkErr(err)
        defer stmtOut.Close()
	err = stmtOut.QueryRow(id).Scan(&u.id, &u.name, &u.email)  //Solo si se espera un único resultado
	checkErr(err)
}


func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/byhours_test")
	checkErr(err)
	defer db.Close()

	user1 := new(User)
	user1.getFromRepository(db, 105671)

	fmt.Printf("user %v", user1)
}
