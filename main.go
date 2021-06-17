package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Pokemon struct {
	ID   string `json:"pokeId"`
	Name string `json:"pokeName"`
}

type User struct {
	UserId string `json:"userid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Name   string `json:"name"`
}

type FullUser struct {
	UserId   string `json:"userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:@/PokemonDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()

	// URL API for Pokemons
	router.HandleFunc("/pokemons", getPokemons).Methods("GET")
	router.HandleFunc("/pokemons", addPokemon).Methods("POST")
	router.HandleFunc("/pokemons/{id}", getPokemon).Methods("GET")
	router.HandleFunc("/pokemons/{id}", editPokemon).Methods("PUT")
	router.HandleFunc("/pokemons/{id}", deletePokemon).Methods("DELETE")

	// URL API for Users
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", editUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("POST")

	http.ListenAndServe(":8000", router)
}

func getPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pokemons []Pokemon
	result, err := db.Query("SELECT pokeid, pokename from pokemons")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var pokemon Pokemon
		err := result.Scan(&pokemon.ID, &pokemon.Name)
		if err != nil {
			panic(err.Error())
		}
		pokemons = append(pokemons, pokemon)
	}
	json.NewEncoder(w).Encode(pokemons)
}

func addPokemon(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO pokemons(pokename) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	pokeName := keyVal["pokename"]
	_, err = stmt.Exec(pokeName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New pokemon was added")
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT pokeid, pokename FROM pokemons WHERE pokeid = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var returnPoke Pokemon
	for result.Next() {
		err := result.Scan(&returnPoke.ID, &returnPoke.Name)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(returnPoke)
}

func editPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE pokemons SET pokename = ? WHERE pokeid = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["pokename"]
	_, err = stmt.Exec(newName, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Pokemon with ID = %s was updated", params["id"])
}

func deletePokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM pokemons WHERE pokeid = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Pokemon with ID = %s was deleted", params["id"])
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allUser []User
	result, err := db.Query("SELECT userid, email, role, name from user")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.UserId, &user.Email, &user.Role, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		allUser = append(allUser, user)
	}
	json.NewEncoder(w).Encode(allUser)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO user(email,password,role,name) VALUES(?,?,0,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	email := keyVal["email"]
	password := keyVal["password"]
	name := keyVal["name"]
	_, err = stmt.Exec(email, password, name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New user has been registered")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT userid, email, role, name FROM user WHERE userid = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.UserId, &user.Email, &user.Role, &user.Name)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := db.Query("SELECT email, name FROM user WHERE userid = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("UPDATE user SET name = ?, email = ? WHERE userid = ?")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.Email, &user.Name)
		if err != nil {
			panic(err.Error())
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newEmail := keyVal["email"]
	newName := keyVal["name"]
	if newEmail == "" {
		newEmail = user.Email
	}
	if newName == "" {
		newName = user.Name
	}

	_, err = stmt.Exec(newName, newEmail, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := db.Query("SELECT userid, email, password, role, name FROM user WHERE userid = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	addstmt, err := db.Prepare("INSERT INTO deletedusers(userid,email,password,role,name) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	delstmt, err := db.Prepare("DELETE FROM user WHERE userid = ?")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	var user FullUser
	for result.Next() {
		err := result.Scan(&user.UserId, &user.Email, &user.Password, &user.Role, &user.Name)
		if err != nil {
			panic(err.Error())
		}
	}

	userid := user.UserId
	email := user.Email
	password := user.Password
	role := user.Role
	name := user.Name

	_, err = addstmt.Exec(userid, email, password, role, name)
	if err != nil {
		panic(err.Error())
	}
	_, err = delstmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}
