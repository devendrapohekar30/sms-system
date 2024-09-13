package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Attendance_20240914_010005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Attendance_20240914_010005{}
	m.Created = "20240914_010005"

	migration.Register("Attendance_20240914_010005", m)
}

// Run the migrations
func (m *Attendance_20240914_010005) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Attendance_20240914_010005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
