package main

import (
	"AsciiArtWeb/pkg/AsciiArt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path != "/":
		app.notFound(w)
		return
	}

	app.render(w, r, "home.page.html", nil)
}

func (app *application) asciiArtWeb(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path != "/ascii-art-web":
		app.notFound(w)
		return
	case r.Method != http.MethodPost:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	AsciiOutput, err := AsciiArt.AsciiArt("Hello there", "standard")
	if err != nil {
		app.serverError(w, err)
		return
	}

	td := &templateData{
		AsciiOutput: AsciiOutput,
	}

	app.render(w, r, "ascii-art.page.html", td)
}
