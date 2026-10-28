package main

import (
	"fmt"
	"net/http"
)

func (app *application) createReelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new Reel \n")
}

//creating a show reel handler. or now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.

func (app *application) showReelHandler(w http.ResponseWriter, r *http.Request) {

	/*  After adding helpers.go fucntion removing this params
	 params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of the reel %d \n", id)
	*/

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
