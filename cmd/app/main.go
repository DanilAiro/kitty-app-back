package main

import (
	"github.com/DanilAiro/kitty-app-back/internal/controllers"
	"github.com/DanilAiro/kitty-app-back/internal/utils"
)

func main() {
	go controllers.Router()

	utils.HandleTermination()
}