package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	id := query.Get("id")
	user := map[string]interface{}{
		"id":   id,
		"name": "test_user",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("参数传递异常"))
		return
	}

	res := struct {
		Msg  string
		Data User
	}{
		Msg:  "user created",
		Data: user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", userHandler)
	mux.HandleFunc("/user/create", createUserHandler)

	http.ListenAndServe(":3000", mux)

}
