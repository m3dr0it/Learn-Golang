package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type UserTemp struct {
	id                          int
	name                        string
	username                    string
	phone_number                string
	email                       string
	registered_by_referral_code sql.NullString
	created_at                  time.Time
	created_by                  string
	updated_at                  sql.NullTime
	updated_by                  sql.NullString
	deleted_at                  sql.NullTime
	deleted_by                  sql.NullString
	otp                         int
	otp_expired                 time.Time
}

func main() {
	db, err := connectToDb()
	if err != err {
		fmt.Println("Error")
	}
	defer db.Close()

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("Error")
	}

	rows, errQ := db.Query("select * from user_temp")

	if errQ != nil {
		fmt.Println("Query error")
		return
	}
	defer rows.Close()

	var result []UserTemp

	for rows.Next() {
		var each = UserTemp{}
		var err = rows.Scan(&each.id, &each.name, &each.username, &each.phone_number,
			&each.email, &each.registered_by_referral_code,
			&each.created_at, &each.created_by,
			&each.deleted_at, &each.deleted_by,
			&each.updated_at, &each.updated_by,
			&each.otp, &each.otp_expired)
		if err != nil {
			fmt.Println("Error")
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name, each.otp)
	}

}

func connectToDb() (*sql.DB, error) {
	var conString string = "user=root dbname=sipulsa password=H8dVeZYjv66xfXbq " +
		"host=localhost port=5433 sslmode=disable"

	db, err := sql.Open("postgres", conString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
