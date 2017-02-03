package filter

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Filter struct {
  minAge string `json:"minage",omitempty`
  maxAge string `json:"maxage",omitempty`
  minDistance string `json:"mindistance",omitempty`
  maxDistance string `json:"maxdistance",omitempty`
  sex string `json:"sex",omitempty`
  userID string `json:"userid",omitempty`
}

var filters []Filter

func CreateFilter(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
	fmt.Println(params)
	fmt.Println(req)
  var filter Filter
  _ = json.NewDecoder(req.Body).Decode(&filter)
	filter.minAge = params["minAge"]
	filter.maxAge = params["maxAge"]
	filter.minDistance = params["minDistance"]
	filter.maxDistance = params["maxDistance"]
	filter.sex = params["sex"]
	filter.userID = params["userID"]
	filters = append(filters, filter)
  json.NewEncoder(w).Encode(filters)

}
