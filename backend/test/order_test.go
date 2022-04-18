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

func TestCreateOrder(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []model.Order{
		model.Order{
			Username:       "Raindrop",
			RestaurantName: "Popeyes",
			OrderDate:      "03/02/2032",
			Price:          14.0,
			CuisineName:    "Nuggets12",
		}, //sucess
		model.Order{
			Username:       "YudiZheng",
			RestaurantName: "Popeyes",
			OrderDate:      "03/02/2022",
			Price:          14.0,
			CuisineName:    "Nuggets12",
		}, //Order already exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "http://localhost:1016/order/create", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}

func TestReadOrder(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []string{
		"YudiZheng",  // success
		"YudiZheng2", // user does not exist
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("GET", "http://localhost:1016/order/read", bytes.NewBuffer(b))
		q := req.URL.Query()
		q.Add("username", test)
		req.URL.RawQuery = q.Encode()
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}
