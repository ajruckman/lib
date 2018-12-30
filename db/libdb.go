package libdb

import (
    "fmt"

    "github.com/jmoiron/sqlx"

    "github.com/ajruckman/lib/err"
)

type Database struct {
    User         string
    Password     string
    Hostname     string
    DatabaseName string
    DB           *sqlx.DB
}

func (db *Database) Init() (err error) {
    var dbString string

    if db.Password != "" {
        dbString = fmt.Sprintf(
            "host=%s user=%s password=%s dbname=%s sslmode=disable",
            db.Hostname,
            db.User,
            db.Password,
            db.DatabaseName,
        )
    } else {
        dbString = fmt.Sprintf(
            "host=%s user=%s dbname=%s sslmode=disable",
            db.Hostname,
            db.User,
            db.DatabaseName,
        )
    }

    db.DB, err = sqlx.Connect("postgres", dbString)
    return
}

func (db *Database) MustInit() {
    liberr.Err(db.Init())
}

/*
Example:

    func init() {
        DB = libdb.Database{
            User:         "user",
            Password:     "pass",
            Hostname:     "ip",
            DatabaseName: "database",
        }
        err := DB.Init()
        liberr.Err(err)
    }

*/
