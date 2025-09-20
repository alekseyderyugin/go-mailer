package main

import (
	"fmt"
	"go-mailer/letters/infrastructure"
	infrastructure2 "go-mailer/shared/infrastructure"
)

func main() {

	type Migratable interface {
		AutoMigrate() error
	}

	var reps []Migratable

	conn := infrastructure2.NewConnection()

	reps = append(reps, infrastructure.NewLetterRepository(conn.DB, infrastructure.NewContext(func(err error) {
		fmt.Println(err)
	})))
	reps = append(reps, infrastructure.NewClientRepository(conn.DB, infrastructure.NewContext(func(err error) {})))

	for _, rep := range reps {
		err := rep.AutoMigrate()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Success migrate")
}
