package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Teacher_20240914_005952 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Teacher_20240914_005952{}
	m.Created = "20240914_005952"

	migration.Register("Teacher_20240914_005952", m)
}

// Run the migrations
func (m *Teacher_20240914_005952) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Teacher_20240914_005952) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
