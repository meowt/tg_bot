package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Start() error {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=1234 dbname=tg_bot sslmode=disable")
	if err != nil {
		return err
	}

	rows, e := db.Query("SELECT * FROM public.chats")
	if e != nil {
		fmt.Println(e.Error())
	}
	var res []int
	for rows.Next() {
		var temp int
		e = rows.Scan(
			&temp)
		if e != nil {
			fmt.Println(e.Error())
		}
		res = append(res, temp)
	}

	fmt.Println(res)
	//_, err = sqlx.Connect("postgres", "user=postgres password=5793924ilyA dbname=tg_bot sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	return nil
}
