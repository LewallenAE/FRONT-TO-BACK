package main

import (
	"log"
	"net/http"
)

// Begin Main
func main() {

	http.HandleFunc("/", ShowAllTasksFunc)
	http.HandleFunc("/add/", AddNewTaskFunc)
	http.HandleFunc("/change/", ChangeTaskFunc)
	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/completed/", CompletedTaskFunc)
	http.HandleFunc("/delete/", DeleteTaskFunc)
	http.HandleFunc("/deleted/", DeletedTaskFunc)
	http.HandleFunc("/edit/", EditTaskFunc)
	http.HandleFunc("/edited/", ShowEditedTaskFunc)
	http.HandleFunc("/show_login/", ShowLoginTaskFunc)
	http.HandleFunc("/login/", LoginTaskFunc)
	http.HandleFunc("/logout/", LogOutTaskFunc)
	http.HandleFunc("/show_register/", ShowRegisterTaskFunc)
	http.HandleFunc("/register/", RegisterTaskFunc)
	http.HandleFunc("/restore/", RestoreTaskFunc)
	http.HandleFunc("/trash/", TrashTaskFunc)
	http.HandleFunc("/update/", UpdateTaskFunc)

	// Port creation
	PORT := "127.0.0.1:8080"
	http.HandleFunc("/", CompleteTaskFunc)
	log.Fatal(http.ListenAndServe(PORT, nil))

	// Simple static webserver:
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("Running server on: " + PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//  End Main
}

// Method: "/"
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// Method: "/add/"
func AddNewTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "add all tasks GET"
	} else {
		message = "add all tasks POST"
	}
	w.Write([]byte(message))
}

// Method: "/change/"
func ChangeTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "change task GET"
	} else {
		message = "change task POST"
	}
	w.Write([]byte(message))
}

// Method: "/complete/"
func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/tasks/"):]
		w.Write([]byte(r.URL.Path))
	}
}

// Method: "/completed/"
func CompletedTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "completed task GET"
	} else {
		message = "completed task POST"
	}
	w.Write([]byte(message))
}

// Method: "/delete/"
func DeleteTaskFunc(w http.ResponseWriter, r *http.Request){
	var message string
	if r.Method == "GET" {
		message = "delete task GET"
	} else {
		message = "delete task POST"
	}
	w.Write([]byte(message))
}

// Method: "/deleted/"
func DeletedTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "deleted task GET"
	} else {
		message = "deleted task POST"
	}
	w.Write([]byte(message))
}

// Method: "/edit/"
func EditTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "edit task GET"
	} else {
		message = "edit task POST"
	}
	w.Write([]byte(message))
}

// Method: "/edited/"
func ShowEditedTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "show edited task GET"
	} else {
		message = "show edited task POST"
	}
	w.Write([]byte(message))
}

// Method: "/show_login/"
func ShowLoginTaskFunc(w http.ResponseWriter, r *http.Request){
	var message string
	if r.Method == "GET" {
		message = "show login task GET"
	} else {
		message = "show login task POST"
	}
	w.Write([]byte(message))
}

// Method: "/login/"
func LoginTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "login task GET"
	} else {
		message = "login task POST"
	}
	w.Write([]byte(message))
}

// Method: "/logout/"
func LogOutTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "logout task GET"
	} else {
		message = "logout task POST"
	}
	w.Write([]byte(message))
}

// Method: "/show_register/"
func RegisterTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "register task GET"
	} else {
		message = "register task POST"
	}
	w.Write([]byte(message))
}

// Method: "/register/"
func ShowRegisterTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method ==  "GET" {
		message = "show register task GET"
	} else {
		message = "show register task POST"
	}
	w.Write([]byte(message))
}

// Method: "/restore/"
func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "restore task GET"
	} else {
		message = "restore task POST"
	}
	w.Write([]byte(message))
}

// Method: "/trash/"
func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "trash task GET"
	} else {
		message = "trash task POST"
	}
	w.Write([]byte(message))
}

// Method: "/update/"
func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET"{
		message = "update task GET"
	} else {
		message = "update task POST"
	}
	w.Write([]byte(message))
}

// Method: "/tasks/"
func GetTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/tasks/"):]
		w.Write([]byte("Get the task: " + id))
	}
}

// Method:
func ServeStaticFile(w http.ResponseWriter, r *http.Request) {

}
