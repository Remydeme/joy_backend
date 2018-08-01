package models

import (
	"time"
)

type arg interface {
}

type Object interface {
	Insert() string
	Args() []arg
	Id() (name string, pk string) // this function will permit to get an element of the object type
}

// structure of user for Json serialisation
// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func (c Credentials) Insert() arg {
	return `INSERT INTO users VALUES($1, $2);`
}

func (c Credentials) Args() []arg {
	return []arg{c.Username, c.Password}
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
	officeId          string    `json:"officeId", db:"officeId"`
	firstname         string    `json:"firstname", db:"firstname"`
	lastname          string    `json:"lastname", db:"lastname"`
	sex               bool      `json:"sex", db:"sex"`
	birth             time.Time `json:"birth", db:"birth"`
	security          string    `json:"security", db:"security"`
	phone             string    `json:"phone", db:"phone"`
	city              string    `json:"city", db:"city"`
	complement        string    `json:"complement", db:"complement"`
	addedThe          time.Time `json:"addedThe", db:"addedThe"`
	patient_pathology string    `json:"patient_pathology", db:"patient_pathology"`
	lastVist          time.Time `json:"lastVist", db:"lastVist"`
}

func (p Patient) Insert() string {
	return `INSERT INTO patient (officeId, firstname, lastname, sex, birth, 
								security, phone, city, complement, patient_pathology, lastVisit) 
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
}

func (p Patient) Args() []arg {
	return []arg{p.officeId, p.firstname, p.lastVist, p.sex, p.birth, p.security, p.phone, p.city, p.complement, p.patient_pathology, p.lastVist}
}

func (p Patient) Id() (string, string) {
	return "patient", "officeId"
}

type Owner struct {
	id        string    `json:"id", db:"id"`
	officeId  string    `json:"officeId", db:"officeId"`
	lastname  string    `json:"lastname", db:"lastname"`
	firstname string    `json:"firstname", db:"firstname"`
	birth     time.Time `json:"birth", db:"birth"`
	email     string    `json:"email", db:"email"`
	upin      string    `json:"upin", db:"upin"`
	isManager bool      `json:"isManager", db:"isManager"`
}

func (o Owner) Insert() string {
	return `INSERT INTO owner(id, officeId, lastname, firstname, birth, email, upin, isManager) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

}

func (o Owner) Args() []arg {
	return []arg{o.id, o.officeId, o.lastname, o.firstname, o.birth, o.email, o.upin, o.isManager}
}

func (o Owner) Id() (string, string) {
	return "owner", "email"
}

type Employee struct {
	officeId string `json:"officeId", db:"officeId"`
	ownerId  string `json:"ownerId", db:"ownerId"`
}

func (e Employee) Insert() string {
	return `INSERT INTO employee (officeId, ownerId) 
			VALUES ($1, $2);`
}

func (e Employee) Args() []arg {
	return []arg{e.officeId, e.ownerId}
}

func (e Employee) Id() (string, string) {
	return "employee", "officeId"
}

type Office struct {
	id             string `json:"id", db:"id"`
	name           string `json:"name", db:"name"`
	city           string `json:"city", db:"city"`
	complement     string `json:"complement", db:"complement"`
	owner_id       int    `json:"owner_id", db:"owner_id"`
	number_patient int    `json:"number_patient", db:"number_patient"`
}

func (o Office) Insert() string {
	return `INSERT INTO office (id, name, city, complement, owner_id, number_patient) 
			VALUES ($1, $2, $3, $4, $5, $6)`
}

func (o Office) Args() []arg {
	return []arg{o.id, o.name, o.city, o.complement, o.owner_id, o.number_patient}
}

func (o Office) Id() (string, string) {
	return "office", "owner_id"
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
	ownerId     string  `json:"ownerId", db:"ownerId"`
	editedBy    string  `json:"editedBy", db:"editedBy"`
	weight      int     `json:"weight", db:"weight"`
	glycemia    int     `json:"glycemia", db:"glycemia"`
	pressure    float32 `json:"pressure", db:"pressure"`
	temperature float32 `json:"temperature", db:"temperature"`
}

func (v VisitSheet) Insert() string {
	return `INSERT INTO visitSheet (ownerId, editedBy, weight, glycemia, pressure, temperature) 
			VALUES ($1, $2, $3, $4, $5, $6);`
}

func (v VisitSheet) args() []arg {
	return []arg{v.ownerId, v.editedBy, v.weight, v.glycemia, v.pressure, v.temperature}
}

func (v VisitSheet) Id() (string, string) {
	return "visitSheet", "ownerId"
}

type TimePeriod struct {
	ownerId  string `json:"ownerId", db:"ownerId"`
	day      string `json:"day", db:"day"`
	time     string `json:"time", db:"time"`
	position int    `json:"position", db:"position"`
	fullname string `json:"fullname", db:"fullname"`
	city     string `json:"city", db:"city"`
}

func (t TimePeriod) Insert() string {
	return `INSERT INTO timePeriod (ownerId, day, time, position, fullname, city) 
			VALUES ($1, $2, $3, $4, $5, $6);`
}

func (t TimePeriod) args() []arg {
	return []arg{t.ownerId, t.day, t.time, t.position, t.fullname, t.city}
}

func (t TimePeriod) Id() (string, string) {
	return "timePeriod", "ownerId"
}
