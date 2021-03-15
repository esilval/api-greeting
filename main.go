package main

import (
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type server struct{}

func greet(w http.ResponseWriter, r *http.Request) {
	// Set the content type into the response writer
	w.Header().Set("Content-Type", "application/json")

	// Get the params
	lan := r.FormValue("lan")
	name := r.FormValue("name")

	if name == "" {
		name = "Bob"
	}

	// Resolve the bundle
	var file string
	switch lan {
	case "en":
		file = "active.en.toml"
	case "es":
		file = "active.es.toml"
	case "de":
		file = "active.de.toml"
	default:

	}

	// Register the bundle
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile(file)

	loc := i18n.NewLocalizer(bundle, lan)

	// Traduce
	helloPerson := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HelloPerson",
			Other: "Hello {{.Name}}!",
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})

	// Set the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"greeting": "` + helloPerson + `"}`))

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/greet", greet).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
