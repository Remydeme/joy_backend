package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/joy/server/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "optional"
	user     = "remy.d.w"
	password = "optional"
	db_name  = "joy"
)

// errors constants

var (
	Err_db_not_connected               = errors.New("Database not connected pointer seems to be nil ")
	db_not_set                         = errors.New("Database pointeur not set ")
	db_table_not_created               = errors.New("Table not created")
	db_query_failed                    = errors.New("Query failed")
	db_query_failed_user_already_exist = errors.New("User already exists")
)

var (
	sql_duplicate_key    = errors.New("duplicate key value violates unique constraint \"users_pkey\"")
	sql_failed_query_row = errors.New("Error while fething query")
)

var DB *DB_s

func init() {
	DB = new(DB_s)
	err := DB.Connect_to_DB()
	if err != nil {
		panic(err)
	}
}

// database class

type DB_s struct {

	// sql database pointeur
	db_ *sql.DB

	// error status variable
	error_ error
}

type User struct {
	username string
	password string
}

type Query struct {
	query string
}

func (q *Query) get() string {
	return q.query
}

/*
/ this funtion connect the database to the rest api and return the pointer
/ of the struct DB
*/

func (database *DB_s) Connect_to_DB() error {

	// create a string that contains all the informations to connect to the database
	dbInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", host, user, db_name)

	var err error

	database.db_, err = sql.Open("postgres", dbInfo)

	if err != nil {
		fmt.Println("Error : ", err)
		return Err_db_not_connected
	}

	if database.db_ == nil {

		log.Fatal("Pointeur nil in connect server function ")
		return Err_db_not_connected
	}

	// the database is not open yet . We just know if the paramters are good
	// to connect to the database we use ping function that will check if the connection is still alive
	// if no it will open a new conention
	err = database.db_.Ping()

	if err != nil {

		database.error_ = err // store the error value

		return err
	}

	return nil
}

// fetch the user password
func (database *DB_s) GetPassword(username string) *sql.Row {

	if database.db_ != nil {

		result := database.db_.QueryRow("select password from users where users.username == $1; ", username)

		return result

	}

	return nil
}

// get username
// fetch the user username
// this function is used to check if a user already exists

func (database *DB_s) GetUsername(username string) *sql.Row {

	if database.db_ != nil {

		result := database.db_.QueryRow("select username from users where users.username == $1; ", username)

		return result

	} else {

		return nil
	}

}

// functions to add new element in database

// this function fetch

// add a new user in the database
func (database *DB_s) AddUser(username string, password []byte) error {

	if database.db_ == nil {

		return Err_db_not_connected

	} else {

		if _, err := database.db_.Query("insert into users values($1, $2);", username, password); err != nil {

			if err == sql_duplicate_key {

				return db_query_failed_user_already_exist

			}

			return db_query_failed

		} else {

			return nil

		}

	}

}

// append a new object in the database

func (database *DB_s) Add(object models.Object) error {

	if database.db_ == nil {

		return Err_db_not_connected

	}
	arguments := []interface{string("233")}
	if _, err := database.db_.Query(object.Insert(), arguments...); err != nil {

		log.Println("errOr", err)
		if err == sql_duplicate_key {

			return db_query_failed_user_already_exist

		}

		return db_query_failed

	}
	return nil
}

func (database *DB_s) Get(key string, object models.Object) (*sql.Row, error) {

	if database.db_ == nil {

		return nil, Err_db_not_connected

	}

	name, pk := object.Id()

	var query = Query{query: string("SELECT * FROM " + name + " WHERE " + pk + " = ?")}

	row := database.db_.QueryRow(query.get(), key)

	if row == nil {

		return nil, sql_failed_query_row

	}

	return row, nil
}
