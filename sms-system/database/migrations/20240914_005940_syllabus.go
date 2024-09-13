package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Syllabus_20240914_005940 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Syllabus_20240914_005940{}
	m.Created = "20240914_005940"

	migration.Register("Syllabus_20240914_005940", m)
}

// Run the migrations
func (m *Syllabus_20240914_005940) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Syllabus_20240914_005940) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
