package handler

import (
	"fmt"
	"net/http"
	"os"

	"desafio-goweb-franconiz/internal/tickets"
	"desafio-goweb-franconiz/pkg/web"

	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := tokenValidator(c); err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
			return
		}

		tickets, err := s.service.GetTotalTickets()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := tokenValidator(c); err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
			return
		}

		destination := c.Param("dest")

		tickets, err := s.service.GetByCountry(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) GetAverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := tokenValidator(c); err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
			return
		}

		destination := c.Param("dest")

		avg, err := s.service.GetAverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}

func tokenValidator(ctx *gin.Context) error {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {

		return fmt.Errorf("Error: Token invalido.")

	}
	return nil
}
