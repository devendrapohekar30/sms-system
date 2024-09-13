package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type MasterStaff_20240914_010247 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &MasterStaff_20240914_010247{}
	m.Created = "20240914_010247"

	migration.Register("MasterStaff_20240914_010247", m)
}

// Run the migrations
func (m *MasterStaff_20240914_010247) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *MasterStaff_20240914_010247) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
