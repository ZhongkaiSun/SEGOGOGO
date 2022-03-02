package test

import (
	"backend/common"
	"backend/config"
	"backend/route"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadRestaurant(t *testing.T) {
	config.InitConfig("test")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()
	r := route.SetupRouter()

	var tests = []string{
		"Popeyes",  // success
		"",         // get all restaurants
		"Popeyes2", //restaurant does not exit
	}
	for idx, test := range tests {
		w := httptest.NewRecorder()
		b, _ := json.Marshal(test)
		req, _ := http.NewRequest("GET", "http://localhost:1016/restaurant/read", bytes.NewBuffer(b))
		q := req.URL.Query()
		q.Add("name", test)
		req.URL.RawQuery = q.Encode()
		r.ServeHTTP(w, req)
		if idx <= 1 {
			assert.Equal(t, 200, w.Code)
		} else {
			assert.Equal(t, 422, w.Code)
		}
	}
}
