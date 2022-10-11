package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"desafio-goweb-franconiz/cmd/server/handler"
	"desafio-goweb-franconiz/internal/domain"
	"desafio-goweb-franconiz/internal/tickets"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("./tickets.csv")
	if err != nil {
		panic(err)
	}

	_ = godotenv.Load()
	repo := tickets.NewRepository(list)
	service := tickets.NewService(repo)
	s := handler.NewService(service)
	r := gin.Default()

	ts := r.Group("/tickets")
	ts.GET("/", s.GetAll())
	ts.GET("/getByCountry/:dest", s.GetTicketsByCountry())
	ts.GET("/getAverage/:dest", s.GetAverageDestination())
	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("\"aqui\": %v\n", "aqui")
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
