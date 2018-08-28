package models

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joy/server/database"
)

var (
	insertUsers      database.Query = `INSERT INTO users(password, username) VALUES (:password, :username)`
	insertPathology  database.Query = `INSERT INTO pathology(title) VALUES(:title)`
	insertPatient    database.Query = `INSERT INTO patient(officeId, firstname, lastname, sex, birth, security, phone, city, complement, addedThe, paltientPathology, lastVisit) VALUES(:officeId, :firstname, :lastname, :sex, :birth, :security, :phone, :city, :complement, :addedThe, :paltientPathology, :lastVisit)`
	insertOwner      database.Query = `INSERT INTO owner(id, officeId, lastname, firstname, birth, email, upin, isManager) VALUES(:id, :officeId, :lastname, :firstname, :birth, :email, :upin, :isManager)`
	insertEmployee   database.Query = `INSERT INTO employee(officeId, ownerId) VALUES(:officeId, :ownerId)`
	insertOffice     database.Query = `INSERT INTO office(id, name, city, complement, ownerid, numberpatient) VALUES(:id, :name, :city, :complement, :ownerid, :numberpatient)`
	insertChef       database.Query = `INSERT INTO chef(officeId, chefId) VALUES(:officeId, :chefId)`
	insertVisitSheet database.Query = `INSERT INTO visitSheet(ownerId, editedBy, weight, glycemia, pressure, temperature) VALUES(:ownerId, :editedBy, :weight, :glycemia, :pressure, :temperature)`
	insertTimePeriod database.Query = `INSERT INTO timePeriod(ownerId, day, time, position, fullname, city) VALUES(:ownerId, :day, :time, :position, :fullname, :city)`
)

var (
	updateUsers      database.Query = `UPDATE users SET password=:password WHERE username=:username)`
	updatePathology  database.Query = `INSERT INTO pathology(title) VALUES(:title)`
	updatePatient    database.Query = `INSERT INTO patient(officeId, firstname, lastname, sex, birth, security, phone, city, complement, addedThe, paltientPathology, lastVisit) VALUES(:officeId, :firstname, :lastname, :sex, :birth, :security, :phone, :city, :complement, :addedThe, :paltientPathology, :lastVisit)`
	updateOwner      database.Query = `INSERT INTO owner(id, officeId, lastname, firstname, birth, email, upin, isManager) VALUES(:id, :officeId, :lastname, :firstname, :birth, :email, :upin, :isManager)`
	updateEmployee   database.Query = `INSERT INTO employee(officeId, ownerId) VALUES(:officeId, :ownerId)`
	updateOffice     database.Query = `INSERT INTO office(id, name, city, complement, ownerId, numberPatient) VALUES(:id, :name, :city, :complement, :ownerId, :numberPatient)`
	updateChef       database.Query = `INSERT INTO chef(officeId, chefId) VALUES(:officeId, :chefId)`
	updateVisitSheet database.Query = `INSERT INTO visitSheet(ownerId, editedBy, weight, glycemia, pressure, temperature) VALUES(:ownerId, :editedBy, :weight, :glycemia, :pressure, :temperature)`
	updateTimePeriod database.Query = `INSERT INTO timePeriod(ownerId, day, time, position, fullname, city) VALUES(:ownerId, :day, :time, :position, :fullname, :city)`
)

type arg interface {
}

type Credentials struct {
	Password string `database:"password"`
	Username string `database:"username"`
}

func (c *Credentials) Insert(tx *sqlx.Tx) error {
	_, err := tx.NamedExec(string(insertUsers), c)
	return err
}

func (c *Credentials) Update(tx *sqlx.Tx) error {
	_, err := tx.NamedExec(string(updateUsers), c)
	return err
}

func (c Credentials) Id() (string, string) {
	return "credentials", "username"
}

// Pathology Object
type Pathology struct {
	title string `json:"title", db:"title"`
}

func (p Pathology) Insert() string {
	return `INSERT INTO timePeriod (title) 
			VALUES ($1);`
}

func (p Pathology) Args() []arg {
	return []arg{p.title}
}

func (p Pathology) Id() (string, string) {
	return "pathology", "title"
}

type Patient struct {
	OfficeId         string    `json:"officeId", db:"officeId"`
	Firstname        string    `json:"firstname", db:"firstname"`
	Lastname         string    `json:"lastname", db:"lastname"`
	Sex              bool      `json:"sex", db:"sex"`
	Birth            time.Time `json:"birth", db:"birth"`
	Security         string    `json:"security", db:"security"`
	Phone            string    `json:"phone", db:"phone"`
	City             string    `json:"city", db:"city"`
	Complement       string    `json:"complement", db:"complement"`
	AddedThe         time.Time `json:"addedThe", db:"addedThe"`
	PatientPathology string    `json:"patient_pathology", db:"patient_pathology"`
	LastVisit        time.Time `json:"lastVist", db:"lastVist"`
}

func (p Patient) Insert() string {
	return `INSERT INTO patient (officeId, firstname, lastname, sex, birth, 
								security, phone, city, complement, patient_pathology, lastVisit) 
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
}

func (p Patient) Id() (string, string) {
	return "patient", "officeId"
}

type Owner struct {
	ID        string    `json:"id", db:"id"`
	OfficeId  string    `json:"officeId", db:"officeId"`
	Lastname  string    `json:"lastname", db:"lastname"`
	Firstname string    `json:"firstname", db:"firstname"`
	Birth     time.Time `json:"birth", db:"birth"`
	Email     string    `json:"email", db:"email"`
	Upin      string    `json:"upin", db:"upin"`
	IsManager bool      `json:"isManager", db:"isManager"`
}

func (o Owner) Insert() string {
	return `INSERT INTO owner(id, officeId, lastname, firstname, birth, email, upin, isManager) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

}

func (o Owner) Args() []arg {
	return []arg{o.ID, o.OfficeId, o.Lastname, o.Firstname, o.Birth, o.Email, o.Upin, o.IsManager}
}

func (o Owner) Id() (string, string) {
	return "owner", "email"
}

type Employee struct {
	OfficeId string `json:"officeId", db:"officeId"`
	OwnerId  string `json:"ownerId", db:"ownerId"`
}

func (e Employee) Insert() string {
	return `INSERT INTO employee (officeId, ownerId) 
			VALUES ($1, $2);`
}

func (e Employee) Args() []arg {
	return []arg{e.OfficeId, e.OwnerId}
}

func (e Employee) Id() (string, string) {
	return "employee", "officeId"
}

type Office struct {
	ID            string `db:"id"`
	Name          string `db:"name"`
	City          string `db:"city"`
	Complement    string `db:"complement"`
	OwnerID       int    `db:"ownerid"`
	NumberPatient int    `db:"numberpatient"`
}

func (o Office) Insert(tx *sqlx.Tx) error {
	_, err := tx.NamedExec(string(insertOffice), o)
	log.Println("In the insert funtion this is the error : ", err)
	return err
}

func (o *Office) Update(tx *sqlx.Tx) error {
	_, err := tx.NamedExec(string(updateUsers), o)
	return err
}

type Chef struct {
	officeId string `json:"officeId", db:"officeId"`
	ownerId  string `json:"ownerId", db:"ownerId"`
}

func (c Chef) Insert() string {
	return `INSERT INTO chef (office_id, ownerId) 
			VALUES ($1, $2);`
}

func (c Chef) Args() []arg {
	return []arg{c.officeId, c.ownerId}
}

func (c Chef) Id() (string, string) {
	return "chef", "office_id"
}

type VisitSheet struct {
	OwnerId     string  `json:"ownerId", db:"ownerId"`
	EditedBy    string  `json:"editedBy", db:"editedBy"`
	Weight      int     `json:"weight", db:"weight"`
	Glycemia    int     `json:"glycemia", db:"glycemia"`
	Pressure    float32 `json:"pressure", db:"pressure"`
	Temperature float32 `json:"temperature", db:"temperature"`
}

func (v VisitSheet) Insert() string {
	return `INSERT INTO visitSheet (ownerId, editedBy, weight, glycemia, pressure, temperature) 
			VALUES ($1, $2, $3, $4, $5, $6);`
}

func (v VisitSheet) args() []arg {
	return []arg{v.OwnerId, v.EditedBy, v.Weight, v.Glycemia, v.Pressure, v.Temperature}
}

func (v VisitSheet) Id() (string, string) {
	return "visitSheet", "ownerId"
}

type TimePeriod struct {
	OwnerId  string `json:"ownerId", db:"ownerId"`
	Day      string `json:"day", db:"day"`
	Time     string `json:"time", db:"time"`
	Position int    `json:"position", db:"position"`
	Fullname string `json:"fullname", db:"fullname"`
	City     string `json:"city", db:"city"`
}

func (t TimePeriod) Insert() string {
	return `INSERT INTO timePeriod (ownerId, day, time, position, fullname, city) 
			VALUES ($1, $2, $3, $4, $5, $6);`
}

func (t TimePeriod) args() []arg {
	return []arg{t.OwnerId, t.Day, t.Time, t.Position, t.Fullname, t.City}
}

func (t TimePeriod) Id() (string, string) {
	return "timePeriod", "ownerId"
}
