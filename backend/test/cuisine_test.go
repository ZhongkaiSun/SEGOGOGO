package test

import (
	"backend/common"
	"backend/config"
	"backend/model"
	"backend/route"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCuisine(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []model.Cuisine{
		model.Cuisine{
			Name:           "Mixed Chicken3",
			RestaurantName: "Popeyes",
			Price:          22.79,
			Calories:       0,
			ImgPath:        "",
		}, //sucess
		model.Cuisine{
			Name:           "Mixed Chicken",
			RestaurantName: "Popeyes2",
			Price:          22.79,
			Calories:       0,
			ImgPath:        "",
		}, // Restaurant does not exist
		model.Cuisine{
			Name:           "Mixed Chicken3",
			RestaurantName: "Popeyes",
			Price:          22.79,
			Calories:       0,
			ImgPath:        "",
		}, // cuisine already exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "http://localhost:1016/cuisine/create", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}

func TestDeleteCuisine(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []model.Cuisine{
		model.Cuisine{
			Name:           "Mixed Chicken3",
			RestaurantName: "Popeyes",
			Price:          22.79,
			Calories:       0,
			ImgPath:        "",
		}, //sucess
		model.Cuisine{
			Name:           "Mixed Chicken",
			RestaurantName: "Popeyes2",
			Price:          22.79,
			Calories:       0,
			ImgPath:        "",
		}, // Restaurant does not exist
		model.Cuisine{
			Name:           "Mixed Chicken3",
			RestaurantName: "Popeyes",
			Price:          22.79,
			Calories:       0,
			ImgPath:        "",
		}, // cuisine does not exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "http://localhost:1016/cuisine/delete", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}

func TestReadCuisine(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []string{
		"Popeyes",  // success
		"Popeyes2", // Restaurant does not exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("GET", "http://localhost:1016/cuisine/read", bytes.NewBuffer(b))
		q := req.URL.Query()
		q.Add("restaurantName", test)
		req.URL.RawQuery = q.Encode()
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}
