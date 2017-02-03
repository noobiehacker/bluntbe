package swipe

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Swipe struct {
  swiperID string `json:"swiperid",omitempty`
  swipeeID string `json:"swipeeid",omitempty`
  yes string `json:"yes",omitempty`
}

var swipes []Swipe

func CreateSwipe(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  var swipe Swipe
  _ = json.NewDecoder(req.Body).Decode(&swipe)
  swipe.swiperID = params["swiperid"]
  swipe.swipeeID = params["swipeeid"]
  swipe.yes = params["yes"]
  swipes = append(swipes, swipe)
  json.NewEncoder(w).Encode(swipes)

}
