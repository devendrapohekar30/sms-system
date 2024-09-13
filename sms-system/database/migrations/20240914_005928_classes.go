package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Classes_20240914_005928 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Classes_20240914_005928{}
	m.Created = "20240914_005928"

	migration.Register("Classes_20240914_005928", m)
}

// Run the migrations
func (m *Classes_20240914_005928) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Classes_20240914_005928) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
