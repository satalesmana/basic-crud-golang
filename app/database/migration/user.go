package migration

import (
	"database/sql"
	"log"
)

func UserMigration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `user` (" +
		"`id` INT NOT NULL AUTO_INCREMENT," +
		"`name` VARCHAR(50) NULL DEFAULT NULL," +
		"`email` VARCHAR(50) NULL DEFAULT NULL," +
		"`password` VARCHAR(50) NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[User Table] berhasil dibuat")
	}

}
