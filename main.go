package main

import (
	"bagus-category-api/database"
	"bagus-category-api/handlers"
	"bagus-category-api/handlers/utils"
	"bagus-category-api/model/repositories"
	"bagus-category-api/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	//Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer db.Close()

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	//GET localhost:8080/categories
	//POST localhost:8080/categories
	http.HandleFunc("/categories", categoryHandler.HandleCategory)

	//GET localhost:8080/categories/{id}
	//PUT localhost:8080/categories/{id}
	//DELETE localhost:8080/categories/{id}
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryById)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteResponse(w, http.StatusOK, "Server is Running, please use the path /categories and /categories{id}", nil)
		w.Header().Set("Content-Type", "application/json")
	})

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Server Failed")
	}
}
