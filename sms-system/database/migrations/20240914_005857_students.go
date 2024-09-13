package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Students_20240914_005857 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Students_20240914_005857{}
	m.Created = "20240914_005857"

	migration.Register("Students_20240914_005857", m)
}

// Run the migrations
func (m *Students_20240914_005857) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Students_20240914_005857) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
