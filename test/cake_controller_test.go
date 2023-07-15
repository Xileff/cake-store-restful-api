package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"felixsavero/cake-store-restful-api/app"
	"felixsavero/cake-store-restful-api/controller"
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/model/domain"
	"felixsavero/cake-store-restful-api/repository"
	"felixsavero/cake-store-restful-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	err := godotenv.Load("../.env")
	helper.PanicIfError(err)
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/cake_store_restful_api_test")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func truncateCake(db *sql.DB) {
	db.Exec("TRUNCATE cakes")
	db.Exec("ALTER TABLE cakes AUTO_INCREMENT = 1")
}

func setupRouter(db *sql.DB) http.Handler {
	validator := validator.New()

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validator)
	cakeController := controller.NewCakeController(cakeService)

	router := app.NewRouter(cakeController)

	return router
}

func TestCreateCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
			"title": "Kue Lapis",
			"description": "Ini kue terenak menurut Felix",
			"rating": 10,
			"image": "https://endeus.tv/resep/kue-lapis-legit"
		}`)
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:5000/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "Kue Lapis", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Ini kue terenak menurut Felix", responseBody["data"].(map[string]interface{})["description"])
	assert.Equal(t, 10, int(responseBody["data"].(map[string]interface{})["rating"].(float64)))
	assert.Equal(t, "https://endeus.tv/resep/kue-lapis-legit", responseBody["data"].(map[string]interface{})["image"])
}

func TestCreateCakeFail(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
			"title": "",
			"description": "",
			"rating": -1,
			"image": ""
		}`)
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:5000/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "fail", responseBody["status"])
}

func TestUpdateCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Kue Lapis",
		Description: "Ini kue terenak menurut Felix",
		Rating:      9.5,
		Image:       "https://endeus.tv/resep/kue-lapis-legit",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
			"title": "Kue Lapis Update",
			"description": "Ini kue terenak menurut Felix Update",
			"rating": 9.5,
			"image": "https://endeus.tv/resep/kue-lapis-legit"
		}`)
	request, _ := http.NewRequest(http.MethodPut, "http://localhost:5000/cakes/"+strconv.Itoa(cake.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, cake.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Kue Lapis Update", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Ini kue terenak menurut Felix Update", responseBody["data"].(map[string]interface{})["description"])
	assert.Equal(t, 9.5, float64(responseBody["data"].(map[string]interface{})["rating"].(float64)))
	assert.Equal(t, "https://endeus.tv/resep/kue-lapis-legit", responseBody["data"].(map[string]interface{})["image"])
}

func TestUpdateCakeFail(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Kue Lapis",
		Description: "Ini kue terenak menurut Felix",
		Rating:      9.5,
		Image:       "https://endeus.tv/resep/kue-lapis-legit",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
			"title": "",
			"description": "",
			"rating": -9.5,
			"image": "https://endeus.tv/resep/kue-lapis-legit"
		}`)
	request, _ := http.NewRequest(http.MethodPut, "http://localhost:5000/cakes/"+strconv.Itoa(cake.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "fail", responseBody["status"])
}

func TestGetCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Kue Lapis",
		Description: "Ini kue terenak menurut Felix",
		Rating:      9.5,
		Image:       "https://endeus.tv/resep/kue-lapis-legit",
	})
	tx.Commit()

	router := setupRouter(db)

	request, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/cakes/"+strconv.Itoa(cake.Id), nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, cake.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Kue Lapis", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Ini kue terenak menurut Felix", responseBody["data"].(map[string]interface{})["description"])
	assert.Equal(t, 9.5, float64(responseBody["data"].(map[string]interface{})["rating"].(float64)))
	assert.Equal(t, "https://endeus.tv/resep/kue-lapis-legit", responseBody["data"].(map[string]interface{})["image"])
}

func TestGetCakeFail(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	router := setupRouter(db)

	request, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/cakes/1", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "fail", responseBody["status"])
}

func TestDeleteCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Kue Lapis",
		Description: "Ini kue terenak menurut Felix",
		Rating:      9.5,
		Image:       "https://endeus.tv/resep/kue-lapis-legit",
	})
	tx.Commit()

	router := setupRouter(db)

	request, _ := http.NewRequest(http.MethodDelete, "http://localhost:5000/cakes/"+strconv.Itoa(cake.Id), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["status"])
}

func TestDeleteCakeFail(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	router := setupRouter(db)

	request, _ := http.NewRequest(http.MethodDelete, "http://localhost:5000/cakes/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "fail", responseBody["status"])
}

func TestListCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake1 := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Kue Lapis",
		Description: "Ini kue terenak menurut Felix",
		Rating:      9.5,
		Image:       "https://endeus.tv/resep/kue-lapis-legit",
	})

	cake2 := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Pisang Goreng",
		Description: "Cemilan enak + murah",
		Rating:      9.9,
		Image:       "https://endeus.tv/resep/pisang-goreng",
	})
	tx.Commit()

	router := setupRouter(db)

	request, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/cakes", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["status"])

	var cakes = responseBody["data"].([]interface{})

	// sorted by rating and alphabetically
	cakeResponse1 := cakes[1].(map[string]interface{})
	cakeResponse2 := cakes[0].(map[string]interface{})

	assert.Equal(t, cake1.Id, int(cakeResponse1["id"].(float64)))
	assert.Equal(t, cake1.Title, cakeResponse1["title"])
	assert.Equal(t, cake1.Description, cakeResponse1["description"])
	assert.Equal(t, cake1.Rating, cakeResponse1["rating"].(float64))
	assert.Equal(t, cake1.Image, cakeResponse1["image"])

	assert.Equal(t, cake2.Id, int(cakeResponse2["id"].(float64)))
	assert.Equal(t, cake2.Title, cakeResponse2["title"])
	assert.Equal(t, cake2.Description, cakeResponse2["description"])
	assert.Equal(t, cake2.Rating, cakeResponse2["rating"].(float64))
	assert.Equal(t, cake2.Image, cakeResponse2["image"])
}
