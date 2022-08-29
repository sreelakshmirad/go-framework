package callbacks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

/*
 *	This controller contains the endpoints for product listing and product detail page
 *	@optional: a valid Session
 *
 *	HTTP Header:
 *		- Authorization :
 *			- Bearer Token : {CustomerID}|{SessionID}|{HashToken}
 *
 */

// API Endpoint : /fruits
// Method : GET
// Get product/fruit list
// @embedded : s *Server pointer
// @param : w http.ResponseWriter HTTP response object
// @param : r *http.Request pointer HTTP request
// @return: void
func GetAllFruits(w http.ResponseWriter, r *http.Request) {

	// var db data.Repo
	// fmt.Println("Hello World")

	// fmt.Println(w, "wwww")
	// fmt.Println(r, "rrrrr")

	// Check for valid session
	// ses, err := db.HasSession(r)
	// ses := nil
	// var sessID int64
	// if err != nil {
	// 	log.Error("Error 250 :", err.Error())
	// sessID = 0
	// } else {
	// 	sessID = ses.ID
	// }
	// fmt.Println("Hello World2")

	// Log listProducts
	// db.LogMetrics(sessID, 10, "")

	// a, err := db.GetFruits()
	// fmt.Println(a, "a")
	// fmt.Println(err, "Error")

	var a []*data.Fruit
	var fruit = data.Fruit{
		ID:   1,
		Name: "Apple",
	}
	a = append(a, &fruit)

	// if err != nil {
	// 	log.Error("Error 251 :", err.Error())
	// 	render.Render(w, r, ErrRender(err))
	// 	return
	// }
	// fmt.Println("Hello World3")

	SendResponse(w, r, a, nil)
}

// API Endpoint : /fruits/{fruitID}
// Method : GET
// Get a product/fruit detail
// @embedded : s *Server pointer
// @param : w http.ResponseWriter HTTP response object
// @param : r *http.Request pointer HTTP request
// @return: void
func GetFruitByID(w http.ResponseWriter, r *http.Request) {
	var db data.Repo

	// Check for valid session

	fmt.Println("Here")
	ses, err := db.HasSession(r)
	var sessID int64
	if err != nil {
		log.Error("Error 250 :", err.Error())
		sessID = 0
	} else {
		sessID = ses.ID
	}

	fruitID, err := strconv.Atoi(chi.URLParam(r, "fruitID"))
	if err != nil {
		log.Error(err.Error())
	}

	a, err := db.GetFruitByID(fruitID)

	if err != nil {
		log.Error("Error 253 :", err.Error())
		render.Render(w, r, ErrRender(err))
		return
	}

	j, err := json.Marshal(a)

	if err != nil {
		log.Error("Error 254 :", err.Error())
		render.Render(w, r, ErrRender(err))
		return
	}

	frStr := strconv.Itoa(fruitID)
	// Log productDetail
	db.LogMetrics(sessID, 11, "a="+frStr)

	SendResponse(w, r, j, ses)
}
