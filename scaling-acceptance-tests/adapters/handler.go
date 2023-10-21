package adapters

import (
	"fmt"
	"github.com/isedaniel/go-specs-greet/domain/interactions"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
