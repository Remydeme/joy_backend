package database

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/op/go-logging"
)

const (
	driver   = "postgres"
	host     = "localhost"
	port     = "optional"
	user     = "remy.d.w"
	password = "optional"
	dbName   = "joy"
)

// errors constants

var (
	ErrFailedToOpenDb             = errors.New("Error faile to open DB.")
	ErrDbNotConnected             = errors.New("Database not connected pointer seems to be nil ")
	DbNotSet                      = errors.New("Database pointeur not set ")
	dbTableNotCreated             = errors.New("Table not created")
	dbQueryFailed                 = errors.New("Query failed")
	dbQueryFailedUserAlreadyExist = errors.New("User already exists")
	ErrFailedDelete               = errors.New("Failed to delete element.")
	ErrNoInserterFound            = errors.New("No inserter method ofund in the object.")
)

var (
	sqlDuplicateKey   = errors.New("duplicate key value violates unique constraint \"users_pkey\"")
	sqlFailedQueryRow = errors.New("Error while fething query")
)

var log = logging.MustGetLogger("db")

type Query string

type QueryX struct {
	Query  Query
	Params []interface{}
}

type DB struct {
	*sqlx.DB
}

type Tx struct {
	*sqlx.Tx
}

// pointer on a database object use
var Database *DB = nil

//tx is an objecct use for transa ction with the database
var Transaction *Tx = nil

func init() {

	var err error

	Database, err = connect_db()

	if err != nil {

		panic(err)
		log.Error(err)
	}
}

func (db *DB) Begin() *Tx {

	tx := db.MustBegin()

	return &Tx{tx}
}

func connect_db() (*DB, error) {

	// create settings string

	dbInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", host, user, dbName)

	db, err := sqlx.Open("postgres", dbInfo)

	if err != nil {

		return nil, err

	}

	if db == nil {

		return nil, ErrFailedToOpenDb

	}

	err = db.Ping()

	return &DB{db}, err
}

func (tx *Tx) Insert(o interface{}) error {

	if u, ok := o.(Inserter); ok {

		err := u.Insert(tx.Tx)

		if err != nil {

			log.Error(err.Error(), "here")

		}

		return err

	}

	log.Debug("No inserter found for object: ", reflect.TypeOf(o))

	return ErrNoInserterFound
}

func (tx *Tx) Getx(o interface{}, qx QueryX) error {

	if u, ok := o.(Getter); ok {
		return u.Get(tx.Tx, qx.Query, qx.Params...)
	}

	stmt, err := tx.Preparex(string(qx.Query))

	if err != nil {
		return err
	}

	return stmt.Get(o, qx.Params...)
}

func (tx *Tx) Delete(o interface{}) error {

	if u, ok := o.(Deleter); ok {

		return u.Delete(tx.Tx)

	}

	return ErrFailedDelete
}

func (tx *Tx) Get(o interface{}, query Query, params ...interface{}) error {

	if u, ok := o.(Getter); ok {

		return u.Get(tx.Tx, query, params...)

	}

	stmt, err := tx.Preparex(string(query))

	if err != nil {

		return err
	}

	return stmt.Get(o, params...)
}

type Updater interface {
	Update(sqlx.Tx) error
}

type Inserter interface {
	Insert(*sqlx.Tx) error
}

type Selecter interface {
	Select(*sqlx.Tx, Query, ...interface{}) error
}
type Getter interface {
	Get(*sqlx.Tx, Query, ...interface{}) error
}
type Deleter interface {
	Delete(*sqlx.Tx) error
}
