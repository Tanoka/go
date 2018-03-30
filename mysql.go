package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	id int
	name string
	email string
	created time.Time
}

func (u *User) getFromRepository(db *sql.DB, id int)  {
	//Importante, trabajar con statements cuando se hacen quieries repetidas, es mucho más rápido
        stmtOut, err := db.Prepare("SELECT id, name, email, created FROM by_customers WHERE id = ?")
        checkErr(err)
        defer stmtOut.Close()
	err = stmtOut.QueryRow(id).Scan(&u.id, &u.name, &u.email, &u.created)  //Solo si se espera un único resultado
	checkErr(err)
}

func (u *User) getAllRepository(db *sql.DB, id int)  {
        stmtOut, err := db.Prepare("SELECT id, hotel_id FROM by_books ORDER BY id desc LIMIT 5 ")
        checkErr(err)
        defer stmtOut.Close()
	rows, err := stmtOut.Query()
	checkErr(err)
	defer rows.Close()

	var id_book, hotel_id int32
	for rows.Next() {
		err = rows.Scan(&id_book, &hotel_id)
		fmt.Println("hotel id: ", hotel_id)
	}

        row, err := stmtOut.Query()
        checkErr(err)
        defer row.Close()

	col, err := row.Columns()
	checkErr(err)
	colT, err := row.ColumnTypes()
	fmt.Println(col)
	//con esto podriamos hacer una funcion que construyera la estructura al llamar a scan
	for _, co := range colT {
		fmt.Print(co.Name())
		nul, ok := co.Nullable()
		fmt.Printf(" null: %v %v", nul, ok)
		fmt.Println(" tipo: ", co.ScanType())
	}
}



func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//IMPORTANTE: para poder obtener campos time, ?parseTime=true
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/byhours_test?parseTime=true")
	checkErr(err)
	defer db.Close()

	//IMPORTANTE, si se va a trabajar con gorotines y concurrencia sobre la bbdd
	db.SetMaxIdleConns(100) 
	db.SetMaxOpenConns(100)

//TODO, get decimal as string not float64!!.. una idea es trabajar con número racionales, fracciones

	user1 := new(User)
	user1.getFromRepository(db, 105671)

	fmt.Printf("user %v \n", user1)

	fmt.Printf("  fecha: %d \n", user1.created.Year())

	user1.getAllRepository(db, 105671)
}
