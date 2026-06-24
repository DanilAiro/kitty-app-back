package service

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/DanilAiro/kitty-app-back/internal/models"
)

func GetKittyPhoto() (string, error) {
	resp, err := http.Get(os.Getenv("KITTY_URL") + "?api_key=" + os.Getenv("KITTY_API_KEY"))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var kitty []models.Kitty

	err = json.NewDecoder(resp.Body).Decode(&kitty)
	if err != nil {
		return "", err
	}

	return kitty[0].URL, nil
}