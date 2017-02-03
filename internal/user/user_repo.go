package user

import (
	"encoding/json"
	"fmt"
	"net/http"
  "github.com/gorilla/mux"
)

type User struct {
  UserName string `json:"username",omitempty"`
  ID string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName string `json:"lastname,omitempty"`
}

// Client provides an interface to the repos package
type Client interface {
	Get(string) ([]Repo, error)
}

// ReposClient provides an implmentation of the Client interface
type ReposClient struct{}

// Repo representes GitHub repository
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetUser(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for _, item := range users {
      if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
      }
  }
  json.NewEncoder(w).Encode(&User{})
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  var user User
  _ = json.NewDecoder(req.Body).Decode(&user)
  user.ID = params["id"]
  users = append(users, user)
  json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range users {
      if item.ID == params["id"] {
          users = append(users[:index], users[index+1:]...)
          break
      }
  }
  json.NewEncoder(w).Encode(users)
}

var users []User

// Get calls the GitHub API to and returns a Repo object for a given user
func (c ReposClient) Get(user string) ([]Repo, error) {
	var r []Repo

	reposURL := fmt.Sprintf("https://api.github.com/users/%s/repos", user)

	res, err := http.Get(reposURL)
	if err != nil {
		return nil, fmt.Errorf("github api: unknown error, %s", err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&r)
		if err != nil {
			return nil, fmt.Errorf("github api: error decoding response %s", err)
		}
	case 404:
		return nil, fmt.Errorf("github api: no results found")
	default:
		return nil, fmt.Errorf("github api: unknown error")
	}

	return r, nil
}
