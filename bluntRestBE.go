package bluntbe

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    //"fmt"
    "github.com/noobiehacker/bluntbe/internal/repos"
)
//USERS
var users []User
var swipes []Swipe
var filters []Filter

type User struct {
  UserName string `json:"username",omitempty"`
  ID string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName string `json:"lastname,omitempty"`
}

type Swipe struct {
  swiperID string `json:"swiperid",omitempty`
  swipeeID string `json:"swipeeid",omitempty`
  yes string `json:"yes",omitempty`
}

type Filter struct {
  minAge string `json:"minage",omitempty`
  maxAge string `json:"maxage",omitempty`
  minDistance string `json:"mindistance",omitempty`
  maxDistance string `json:"maxdistance",omitempty`
  sex string `json:"sex",omitempty`
  userID string `json:"userid",omitempty`
}

func CreateFilter(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  var filter Filter
  _ = json.NewDecoder(req.Body).Decode(&filter)
  filter.minAge = params["minage"]
  filter.maxAge = params["maxage"]
  filter.minDistance = params["mindistance"]
  filter.maxDistance = params["maxdistance"]
  filter.sex = params["sex"]
  filter.userID = params["userid"]
  filters = append(filters,filter)
  json.NewEncoder(w).Encode(filters)
}

func CreateSwipeEndPoint(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  var swipe Swipe
  _ = json.NewDecoder(req.Body).Decode(&swipe)
  swipe.swiperID = params["swiperid"]
  swipe.swipeeID = params["swipeeid"]
  swipe.yes = params["yes"]
  swipes = append(swipes, swipe)
  json.NewEncoder(w).Encode(swipes)

}

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for _, item := range users {
      if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
      }
  }
  json.NewEncoder(w).Encode(&User{})
}

func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  var user User
  _ = json.NewDecoder(req.Body).Decode(&user)
  user.ID = params["id"]
  users = append(users, user)
  json.NewEncoder(w).Encode(users)
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range users {
      if item.ID == params["id"] {
          users = append(users[:index], users[index+1:]...)
          break
      }
  }
  json.NewEncoder(w).Encode(users)
}

// App defines the application container
type App struct {
	repos repos.Client
}

// GetReposHandler returns a list of (public) repositories for a given GitHub user
func (a *App) GetReposHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if user == "" {
		http.Error(w, "MISSING_ARG_USER", 400)
		return
	}

	repos, err := a.repos.Get(user)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}

	b, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {

    app := &App{repos: repos.ReposClient{}}
    router := mux.NewRouter()
    users = append(users, User{ID : "0", UserName : "Giovanni", FirstName : "David", LastName : "Chan" })
    users = append(users, User{ID: "1", UserName: "Soldier", FirstName: "Jon", LastName: "Huang" })
    users = append(users, User{ID: "2", UserName: "MayLing", FirstName: "Peter", LastName: "Ma" })
    users = append(users, User{ID: "3", UserName: "TightMan", FirstName: "Sadrick", LastName: "Sing" })
    router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
    router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
    router.HandleFunc("/user/{id}", CreateUserEndpoint).Methods("POST")
    router.HandleFunc("/user/{id}", DeleteUserEndpoint).Methods("DELETE")
    router.HandleFunc("/repos", app.GetReposHandler)
    log.Fatal(http.ListenAndServe(":12345", router))
    log.Println("listening on 12345")
}
