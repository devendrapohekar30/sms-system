package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ExamesResults_20240914_010202 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ExamesResults_20240914_010202{}
	m.Created = "20240914_010202"

	migration.Register("ExamesResults_20240914_010202", m)
}

// Run the migrations
func (m *ExamesResults_20240914_010202) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *ExamesResults_20240914_010202) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
