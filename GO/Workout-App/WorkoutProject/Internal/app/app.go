package app

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

// no parameters
// * is the pointer to the Application return type *Application and error are both return types.
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// & is the "address of" for pointer
	app := &Application{
		Logger: logger,
	}

	// (app, error)
	return app, nil

}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available.")

}