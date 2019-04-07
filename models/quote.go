package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"helloworld/db"
)

type Quote struct {
	ID       string `json:"id"`
	Type     int    `json:"type"`
	Quote    string `json:"quote"`
	Category string `json:"category"`
	Author   string `json:"author"`
}
type Quotes struct {
	Quotes []Quote `json:"quote"`
}

var con *sql.DB

func AddQuote(quote *Quote) error {
	con := db.CreateCon()
	sql := "INSERT INTO quotes(type, quote, category, author) VALUES( ?, ?, ?, ?)"
	stmt, err := con.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(quote.Type, quote.Quote, quote.Category, quote.Author)

	// Exit if we get an error
	if err2 != nil {
		return err2
		fmt.Print(err2.Error())
	}
	fmt.Println(result.LastInsertId())
	return nil
}

func GetQuotes() Quotes {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT id,type, quote, category, author FROM quotes order by id desc limit 50"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Quotes{}
	for rows.Next() {
		quote := Quote{}

		err2 := rows.Scan(&quote.ID, &quote.Type, &quote.Quote, &quote.Category, &quote.Author)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Quotes = append(result.Quotes, quote)
	}
	return result
}
