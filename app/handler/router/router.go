package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Artpou/elastic-search-go/controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func HandleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8080/")
	router := mux.NewRouter().StrictSlash(true)
	//db = database.InitDb()

	//Login
	router.HandleFunc("/", HomeLink)
	router.HandleFunc("/api/login/", Login).Methods("POST")
	router.HandleFunc("/api/logout/", Logout).Methods("GET")
	router.HandleFunc("/api/checkAuth/", CheckAuth).Methods("GET")
	router.HandleFunc("/api/checkAdmin/", CheckAdmin).Methods("GET")

	//Users
	router.HandleFunc("/api/users/", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")

	//Self
	router.HandleFunc("/api/self/", GetSelf).Methods("GET")
	router.HandleFunc("/api/self/", UpdateSelf).Methods("PUT")
	router.HandleFunc("/api/self/", DeleteSelf).Methods("DELETE")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", router))
}

//LOGIN

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func Login(w http.ResponseWriter, r *http.Request) {
	controllers.Signin(db, w, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	controllers.Logout(db, w, r)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	controllers.CheckAuth(w, r)
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) {
	controllers.CheckAdmin(w, r)
}

//USER

func GetUsers(w http.ResponseWriter, r *http.Request) {
	controllers.GetUsers(db, w, r)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	controllers.GetUser(db, w, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateUser(db, w, r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	controllers.CreateUser(db, w, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteUser(db, w, r)
}

func GetSelf(w http.ResponseWriter, r *http.Request) {
	controllers.GetSelf(db, w, r)
}

func UpdateSelf(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateSelf(db, w, r)
}

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteSelf(db, w, r)
}
