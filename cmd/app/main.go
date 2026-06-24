package main

import (
	"github.com/DanilAiro/kitty-app-back/internal/controllers"
	"github.com/DanilAiro/kitty-app-back/internal/initializers"
	"github.com/DanilAiro/kitty-app-back/internal/utils"
)

func main() {
	initializers.LoadEnvVariables()
	
	go controllers.Router()

	utils.HandleTermination()
}