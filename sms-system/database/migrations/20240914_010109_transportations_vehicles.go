package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type TransportationsVehicles_20240914_010109 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TransportationsVehicles_20240914_010109{}
	m.Created = "20240914_010109"

	migration.Register("TransportationsVehicles_20240914_010109", m)
}

// Run the migrations
func (m *TransportationsVehicles_20240914_010109) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *TransportationsVehicles_20240914_010109) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
