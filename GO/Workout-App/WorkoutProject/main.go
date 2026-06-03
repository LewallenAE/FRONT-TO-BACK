package main

import (
	"net/http"
	"time"
	"fmt"
	"github.com/LewallenAE/WorkoutApp/Internal/app"
)

// Beginning of main()
func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our app.")

	// Declare Server
	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:        ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

	// end main ()
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available.")

}
