package main
 
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    . "fmt"
)
 
func insert(db *sql.DB) {
    stmt, err := db.Prepare("INSERT INTO user(name, password) VALUES(?, ?)")
    defer stmt.Close()
 
    if err != nil {
        Println(err)
        return
    }
    stmt.Exec("guotie", "guotie")
    stmt.Exec("testuser", "123123")
 
}
 
func main(){
    db, err := sql.Open("mysql", "test:test@tcp(10.10.4.37:3306)/test?charset=utf8")
    if err != nil {
        Println("Open database error: %s\n", err)
    }
    defer db.Close()
   
    err = db.Ping()
    if err != nil {
        Println(err)
    }
 
    insert(db)
 
    rows, err := db.Query("select id, name from user where id = ?", 2)
    if err != nil {
        Println(err)
    }
 
    defer rows.Close()
    var id int
    var name string
    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            Println(err)
        }
        Println(id, name)
    }
 
    err = rows.Err()
    if err != nil {
        Println(err)
    }
    return
}