package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/shahidshabbir-se/quirky/internal/db/sqlc"
)

type Server struct {
	app     *fiber.App
	queries *db.Queries
}

func NewServer(queries *db.Queries) *Server {
	app := fiber.New()

	server := &Server{
		app:     app,
		queries: queries,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	s.app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Quirky API is alive!")
	})

	// Example user route (if you have a CreateUser query)
	s.app.Post("/users", func(c *fiber.Ctx) error {
		type request struct {
			Username string `json:"username"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return fiber.ErrBadRequest
		}

		user, err := s.queries.CreateUser(c.Context(), body.Username)
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			return fiber.ErrInternalServerError
		}

		return c.JSON(user)
	})
}

func (s *Server) Start(address string) error {
	return s.app.Listen(address)
}
