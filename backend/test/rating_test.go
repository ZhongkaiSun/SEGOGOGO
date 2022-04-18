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

func TestCreateRating(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []model.Rating{
		model.Rating{
			Username:       "ZhongkaiSun",
			RestaurantName: "Popeyes",
			Star:           5,
			Comment:        "I would like to come again",
			RatingDate:     "2020/01/03",
		}, // success
		model.Rating{
			Username:       "ZhongkaiSun5",
			RestaurantName: "Popeyes",
			Star:           5,
			Comment:        "I would like to come again",
			RatingDate:     "2021/01/03",
		}, // customer does not exist
		model.Rating{
			Username:       "ZhongkaiSun",
			RestaurantName: "KFC1",
			Star:           5,
			Comment:        "I would like to come again",
			RatingDate:     "2021/01/03",
		}, // Restaurant does not exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "http://localhost:1016/rating/create", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}

func TestGetRating(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []string{
		"Popeyes", // success
		"KFC1",    // Restaurant does not exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "http://localhost:1016/rating/get", nil)
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
