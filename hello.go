package main

import (
  "log"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func main() {

  db, err := sql.Open("mysql", "root:new-password@tcp(54.199.220.41:3306)/test00")

if err != nil {
	log.Println(err)
	return
}

  m := martini.Classic()
  // render html templates from templates directory
  m.Use(render.Renderer())

  m.Get("/", func(r render.Render) {
    //r.HTML(200, "hello", "jeremy")
    
    // Prepare statement for reading data
    stmtOut, err := db.Query("SELECT squareNumber FROM squarenum")
    
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()

    

for stmtOut.Next() {
	var temp int

	err := stmtOut.Scan(&temp)

	    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Println("The square number of 1 is:", temp)
}

/*
	//var squareNum int // we "scan" the result in here

    // Query another number.. 1 maybe?
    err = stmtOut.QueryRow(9999).Scan(&squareNum) // WHERE number = 1
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Println("The square number of 1 is: %d", squareNum)
*/
    r.JSON(200, map[string]interface{}{"hello" : "world"})


    //log.Fatal(m.RunOnAddr(":5050"))
  })
  
	m.RunOnAddr(":15050")
  m.Run()
}