package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Notifications_20240914_010329 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Notifications_20240914_010329{}
	m.Created = "20240914_010329"

	migration.Register("Notifications_20240914_010329", m)
}

// Run the migrations
func (m *Notifications_20240914_010329) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Notifications_20240914_010329) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
