package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
)

const (
    host     = "database" // Use the service name from docker-compose.yml
    port     = 5432
    user     = "your_user" // Replace with your actual PostgreSQL username
    password = "your_password" // Replace with your actual PostgreSQL password
    dbname   = "your_db" // Replace with your actual PostgreSQL database name
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database!")

    // Query the users table and print the results to the terminal
    rows, err := db.Query("SELECT name, email FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var name string
        var email string
        err := rows.Scan(&name, &email)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Name: %s, Email: %s\n", name, email)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}

