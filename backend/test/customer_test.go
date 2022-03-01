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

func TestLogin(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []struct {
		username string
		password string
	}{
		{"ZhongkaiSun", "123456"},
		{"ZhongkaiSun", "12345"},   // wrong password
		{"ZhongkaiSun1", "123456"}, // wrong username
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "http://localhost:1016/customer/login", nil)
		q := req.URL.Query()
		q.Add("username", test.username)
		q.Add("password", test.password)
		req.URL.RawQuery = q.Encode()
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}

	}
}

func TestRegister(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []model.Customer{
		model.Customer{
			Username:     "ZhongkaiSun1",
			Password:     "123456",
			AddressLine1: "3824 SW Archer Road",
			AddressLine2: "Apt 5-501",
			Phone:        "7839291765",
			Email:        "kdlaif@gmail.com",
		}, // success
		model.Customer{
			Username:     "ZhongkaiSun",
			Password:     "123456",
			AddressLine1: "3824 SW Archer Road",
			AddressLine2: "Apt 5-501",
			Phone:        "7839291765",
			Email:        "kdlaif@gmail.com",
		}, // customer already exists
		model.Customer{
			Username:     "ZhongkaiSun2",
			Password:     "12345",
			AddressLine1: "3824 SW Archer Road",
			AddressLine2: "Apt 5-501",
			Phone:        "7839291765",
			Email:        "kdlaif@gmail.com",
		}, // password length does not match the requirement
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "http://localhost:1016/customer/register", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}

	}
}

func TestDeleteUser(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []struct {
		username string
		password string
	}{
		{"ZhongkaiSun1", "123456"}, //success
		{"ZhongkaiSun5", "12345"},  // customer does not exist
		{"ZhongkaiSun", "12345"},   // wrong password
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "http://localhost:1016/customer/delete", nil)
		q := req.URL.Query()
		q.Add("username", test.username)
		q.Add("password", test.password)
		req.URL.RawQuery = q.Encode()
		r.ServeHTTP(w, req)
		if idx == 0 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}

	}
}
