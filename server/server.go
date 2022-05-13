package main

import (
	"github.com/gofiber/fiber/v2"
	zmq "github.com/pebbe/zmq4/draft"
)

type Input struct {
	Text string `json: "text"`
}
type Type int

func main() {
	app := fiber.New()
	client, _ := zmq.NewContext()

	socket, _ := client.NewSocket(zmq.REQ)
	socket.Bind("tcp://*:6666")

	app.Get("/", func(c *fiber.Ctx) error {
		input := new(Input)
		if err := c.BodyParser(input); err != nil {
			panic(err)
		}
		socket.Send(input.Text, 0)

		if msg, err := socket.Recv(0); err != nil {
			panic(err)
		} else {
			return c.JSON(fiber.Map{"Length": msg})
		}
	})

	app.Listen(":3000")
}
