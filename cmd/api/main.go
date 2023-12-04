package main

import (
	"net/http"

	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/igorgalindop/full-cycle-go-intensive-november-2023/internal/entity"

	"github.com/labstack/echo/v4"
)

func main() {
	// routers := chi.NewRouter()
	// routers.Use(middleware.Logger)
	// routers.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", routers)
	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("1", 10, 1)
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, order)
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	order, _ := entity.NewOrder("1", 10, 1)
// 	err := order.CalculateFinalPrice()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	json.NewEncoder(w).Encode(order)
// }
