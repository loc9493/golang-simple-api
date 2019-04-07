package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"helloworld/db"

	"github.com/jinzhu/gorm"
)

type Quote struct {
	gorm.Model
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

func AddQuote(quote *Quote) {
	con := db.CreateCon()
	con.Create(quote)
	defer con.Close()
}

func UpdateQuote(quote *Quote, id int) Quote {
	con := db.CreateCon().Table("quotes")
	var new_quote Quote
	fmt.Println("Get qutoe")
	con.First(&new_quote, id)

	fmt.Println(new_quote.Quote, id)
	fmt.Println("new qutoe %d", id)
	con.Model(&new_quote).Updates(quote)
	// new_quote.Type = quote.Type
	// new_quote.Category = quote.Category
	// new_quote.Quote = quote.Quote
	// new_quote.Author = quote.Author
	// con.Save(&new_quote)
	defer con.Close()
	return new_quote
}

func GetQuote(id int) Quote {
	con := db.CreateCon()
	var quote Quote
	con.Table("quotes").Find(&quote, id)
	defer con.Close()
	return quote
}

func GetQuotes(page int) []Quote {
	limit := 50
	con := db.CreateCon()
	var result []Quote
	quotes := Quotes{}
	con.Table("quotes").Offset(page * limit).Limit(limit).Find(&result)
	fmt.Println(result)
	quotes.Quotes = result
	defer con.Close()
	return result
}
