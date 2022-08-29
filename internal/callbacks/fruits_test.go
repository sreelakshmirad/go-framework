package callbacks

import (
	"fmt"
	helper "framework_v1/internal/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestgetAllFruits(t *testing.T) {

	c := new(Config)
	c.DBConnString = helper.MustReadMysqlConf()
	s := NewServer(c)
	req, err := http.NewRequest("GET", "/fruitqs", nil)
	fmt.Printf("t: %v\n", t)
	t.Logf("Hello there")
	if err != nil {
		t.Fatalf("Error : %v", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.getAllFruits)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: Got %v Want %v", status, http.StatusOK)
	}

}

// func TestGetFruitByID(t *testing.T) {

// 	c := new(Config)
// 	c.DBConnString = helper.MustReadMysqlConf()
// 	s := NewServer(c)
// 	req, err := http.NewRequest("GET", "/fruits", nil)
// 	fmt.Printf("t: %v\n", t)
// 	if err != nil {
// 		t.Fatalf("Error : %v", err.Error())
// 	}

// 	q := req.URL.Query()
// 	q.Add("fruitID", "1")
// 	req.URL.RawQuery = q.Encode()

// }
