package bluntbe

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/noobiehacker/bluntbe/internal/repo"
    "github.com/noobiehacker/bluntbe/internal/user"
    //"fmt"
    //"github.com/noobiehacker/bluntbe/internal/filter"
    //"github.com/noobiehacker/bluntbe/internal/swipe"
)

// App defines the application container
type App struct {
	repo repo.Client
}

// GetReposHandler returns a list of (public) repositories for a given GitHub user
func (a *App) GetReposHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if user == "" {
		http.Error(w, "MISSING_ARG_USER", 400)
		return
	}

	repos, err := a.repo.Get(user)
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

var users []user.User

func main() {

    app := &App{repo: repo.ReposClient{}}
    router := mux.NewRouter()
    router.HandleFunc("/users", user.GetUsers).Methods("GET")
    router.HandleFunc("/user/{id}", user.GetUser).Methods("GET")
    router.HandleFunc("/user/{id}", user.CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", user.DeleteUser).Methods("DELETE")
    router.HandleFunc("/repos", app.GetReposHandler)
    log.Fatal(http.ListenAndServe(":12345", router))
    log.Println("listening on 12345")
}

func init(){
  users = append(users, user.User{ID : "0", UserName : "Giovanni", FirstName : "David", LastName : "Chan" })
  users = append(users, user.User{ID: "1", UserName: "Soldier", FirstName: "Jon", LastName: "Huang" })
  users = append(users, user.User{ID: "2", UserName: "MayLing", FirstName: "Peter", LastName: "Ma" })
  users = append(users, user.User{ID: "3", UserName: "TightMan", FirstName: "Sadrick", LastName: "Sing" })
}
