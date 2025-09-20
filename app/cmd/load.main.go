package main

import (
	"fmt"
	"go-mailer/letters/domain"
	"go-mailer/letters/infrastructure"
	core "go-mailer/shared/infrastructure"
	"os"
	"strconv"
)

func main() {
	app := core.NewConnection()

	context := infrastructure.NewContext(func(err error) {
		fmt.Println(err)
	})

	letterRepository := infrastructure.NewLetterRepository(app.DB, context)

	clientId := domain.NewClientId()
	client := domain.NewClient(clientId)
	clientRepository := infrastructure.NewClientRepository(app.DB, context)
	err := clientRepository.Save(client)

	if err != nil {
		panic(err)
	}

	count := getCount()
	letters := make([]*domain.Letter, count)

	for i := 0; i < count; i++ {
		letterId := domain.NewLetterID()
		uniqId := string(letterId)

		letter := domain.NewLetter(
			letterId,
			uniqId+"_sender@test.test",
			[]domain.Address{domain.NewAddress(uniqId+"_recipient@test.test", "")},
			"some subject",
			domain.HtmlMessage("some html message "+uniqId),
			domain.PlainMessage("some plain message "+uniqId),
			clientId,
		)

		letters[i] = letter
	}

	err = letterRepository.CreateBatch(letters)

	if err != nil {
		panic(err)
	}

	fmt.Println(count, "new letters loaded!")
}

func getCount() int {
	defaultCount := 1000

	args := os.Args[1:]

	if len(args) < 1 {
		return defaultCount
	}

	count, err := strconv.Atoi(args[0])

	if err != nil {
		panic(err)
	}

	if count < 0 {
		panic("Укажите корректное количество писем")
	} else if count == 0 {
		count = 1000
	}

	return count
}
