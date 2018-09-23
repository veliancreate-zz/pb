package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
	// "strconv"
)

var coupons = []byte(`[
	{
		“name”: “Save £20 at Tesco”,
		“brand”: "Tesco",
		“value”: 20,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £30 at Tesco”,
		“brand”: "Tesco",
		“value”: 30,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £40 at Tesco”,
		“brand”: "Tesco",
		“value”: 40,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £50 at Tesco”,
		“brand”: "Tesco",
		“value”: 50,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £20 at Asda,
		“brand”: "Asda",
		“value”: 20,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £30 at Asda,
		“brand”: "Asda",
		“value”: 30,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £40 at Asda,
		“brand”: "Asda",
		“value”: 40,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £50 at Tesco”,
		“brand”: "Tesco",
		“value”: 50,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £20 at Sainsburys,
		“brand”: "Sainsburys",
		“value”: 20,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	},
	{
		“name”: “Save £20 at Sainsburys”,
		“brand”: "Sainsburys",
		“value”: 20,
		“createdAt”: “2018-03-01 10:15:53”,
		“expiry”: “2019-03-01 10:15:53”
	}
]`)

type API interface {
	list() []Coupon
	update(id string, body interface{}) Coupon
	get(id string) Coupon
	create(body interface{}) Coupon
}

type Coupon struct {
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	Expiry    string `json:"expiry"`
}

// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// create
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		d := &Coupon{}
		err2 := json.Unmarshal(body, d)

		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(d, "POST done")
		// get
	} else {
		w.Write(coupons)
	}
}

// GetHandler handles the index route
func Get(w http.ResponseWriter, r *http.Request) {
	// id := strings.TrimPrefix(r.URL.Path, "/coupon/")
	d := &[]Coupon{}
	e := json.Unmarshal(coupons, d)

	if e != nil {
		fmt.Println(e)
	}

	// i, e2 := strconv.ParseInt(id, 10, 64)

	// if (e2 != nil) {
	// 	fmt.Println(e2)
	// }


	// if i >= 0 && i < len(d) {
	// 	item := (*d)[i]

	// }

	// json, err3 := json.Marshal(item)

	// if (err3 != nil) {
	// 	fmt.Println(err3)
	// }

	// w.Write(json)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/coupon", GetHandler)
	mux.HandleFunc("/coupon/", Get)

	http.ListenAndServe(":3000", mux)
}
