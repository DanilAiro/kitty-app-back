package main

import (
	"github.com/DanilAiro/kitty-app-back/internal/controllers"
	"github.com/DanilAiro/kitty-app-back/internal/initializers"
	"github.com/DanilAiro/kitty-app-back/internal/repository"
	"github.com/DanilAiro/kitty-app-back/internal/utils"
)

func main() {
	initializers.LoadEnvVariables()
	repository.ConnectDB()
	initializers.CreateLogger()

	initializers.Logger.Info().Msg("server start up")
	defer initializers.Logger.Info().Msg("server stop work")
	
	go controllers.Router()
	
	utils.HandleTermination()

}