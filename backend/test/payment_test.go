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

func TestCreatePayment(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []model.Payment{
		model.Payment{
			Username:     "YudiZheng",
			CardHolder:   "YudiZheng",
			CardNumber:   "2234567890XXXXXX",
			ExpDate:      "09/25",
			SecurityCode: "456",
			AddressLine1: "",
			AddressLine2: "",
			City:         "",
			State:        "",
			Zipcode:      "",
		}, //sucess
		model.Payment{
			Username:     "YudiZheng",
			CardHolder:   "YudiZheng",
			CardNumber:   "2234567890XXXXXX",
			ExpDate:      "09/25",
			SecurityCode: "456",
			AddressLine1: "",
			AddressLine2: "",
			City:         "",
			State:        "",
			Zipcode:      "",
		}, //Payment already exist
		model.Payment{
			Username:     "YudiZheng1",
			CardHolder:   "YudiZheng",
			CardNumber:   "2234567890XXXXXX",
			ExpDate:      "09/25",
			SecurityCode: "456",
			AddressLine1: "",
			AddressLine2: "",
			City:         "",
			State:        "",
			Zipcode:      "",
		}, //Customer does not exist
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
