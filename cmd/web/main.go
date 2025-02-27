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

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// Use the Infi() method to log the starting server message at Info severity along with the listen address as an attribute
	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	// And we also use the Error() method to log any error message returned by the http.ListenAndServe() function
	logger.Error(err.Error())
	os.Exit(1)
}
