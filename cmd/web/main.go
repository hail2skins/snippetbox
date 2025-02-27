package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the web application.
type application struct {
	logger *slog.Logger
}

func main() {

	// Define a new command-line flag with the name 'addr', a default value of ":4000" and some short help text explaining what the flag controls
	addr := flag.String("addr", ":4000", "HTTP network address")

	// We use flag.Parse() to parse the command line flag.
	// flag.Parse() will update the value of addr which we then pass to the http.ListenAndServe() function.
	flag.Parse()

	// Use the slog.New() function to initialize a new structured logger, which is used to log messages with different severity levels.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		logger: logger,
	}

	// Use the Infi() method to log the starting server message at Info severity along with the listen address as an attribute
	logger.Info("starting server", "addr", *addr)

	// Call the new app.routes() method to get the servemux containing our routes
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
