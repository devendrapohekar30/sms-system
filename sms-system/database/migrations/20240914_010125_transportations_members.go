package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type TransportationsMembers_20240914_010125 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TransportationsMembers_20240914_010125{}
	m.Created = "20240914_010125"

	migration.Register("TransportationsMembers_20240914_010125", m)
}

// Run the migrations
func (m *TransportationsMembers_20240914_010125) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *TransportationsMembers_20240914_010125) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
