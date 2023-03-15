package main

import (
	"fmt"

	"github.com/fatih/structs"
    "github.com/mitchellh/mapstructure"
)
import surrealdb "github.com/surrealdb/surrealdb.go"


type Database struct {
    db *surrealdb.DB
}

func ConnectDatabse(port int, user string, pass string, namespace string, database string) (*Database, error) {
    db, err := surrealdb.New(fmt.Sprintf("ws://localhost:%d/rpc", port))
    if err != nil {
        return nil, err
    }

    _, err = db.Signin(map[string]interface{} {
        "user": user,
        "pass": pass,
    })
    if err != nil {
        return nil, err
    }

    _, err = db.Use(namespace, database)
    if err != nil {
        return nil, err
    }

    return &Database{db: db}, nil
}

func (db *Database) Close() {
    db.db.Close()
}

func (db *Database) Select(table string, output interface{}) error {
    data, err := db.db.Select(table)
    if err != nil {
        return err
    }
    return mapstructure.Decode(data, &output)
}

func (db *Database) Create(table string, entry interface{}) (interface{}, error) {
    result, err := db.db.Create(table, structs.Map(entry))
    if err != nil {
        return nil, err
    }
    return result, nil
}

func (db *Database) Delete(what string) error {
    _, err := db.db.Delete(what)
    return err
}

