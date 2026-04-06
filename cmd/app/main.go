package main

import (
	controllers "github.com/DanilAiro/kitty-app-back/internal/controllers"
	initializers "github.com/DanilAiro/kitty-app-back/internal/initializers"
)

func main() {
	initializers.LoadEnvVariables()
	controllers.Router()
}