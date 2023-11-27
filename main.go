package main

import (
	"database/sql"
	"fmt"

	"github.com/igorgalindop/full-cycle-go-intensive-november-2023/internal/infra/database"
	"github.com/igorgalindop/full-cycle-go-intensive-november-2023/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

}
