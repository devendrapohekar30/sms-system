package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Transportations_20240914_010053 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Transportations_20240914_010053{}
	m.Created = "20240914_010053"

	migration.Register("Transportations_20240914_010053", m)
}

// Run the migrations
func (m *Transportations_20240914_010053) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Transportations_20240914_010053) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
