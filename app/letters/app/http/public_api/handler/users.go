package handler

import (
	"fmt"
	"go-mailer/letters/domain"
	"go-mailer/letters/infrastructure"
	core "go-mailer/shared/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func PostUser(fiberContext *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	type Response struct {
		Status  string
		Message string
		Payload map[string]string
	}

	request := Request{}

	if err := fiberContext.BodyParser(&request); err != nil {
		return err
	}

	clientId := domain.NewClientId()

	clientName, err := domain.NewClientName(request.Name)

	if err != nil {
		return fiberContext.Status(400).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	client, err := domain.NewClient(clientId, clientName)

	app := core.NewConnection()

	appContext := infrastructure.NewContext(func(err error) {
		fmt.Println(err)
	})

	repo := infrastructure.NewClientRepository(app.DB, appContext)

	err = repo.Save(client)

	if err != nil {
		return fiberContext.Status(500).JSON(Response{})
	}

	return fiberContext.JSON(Response{
		Status: "success",
		Payload: map[string]string{
			"id": string(clientId),
		},
	})
}
