package database

import (
	"app-basic-crud/app/database/migration"
)

func Migrate() {
	db := GetCoon()
	migration.UserMigration(db)

}
