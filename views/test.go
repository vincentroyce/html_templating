package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func Multiply(x, y int) int {
	return x * y
}

func TestTemplatingHandler(w http.ResponseWriter, r *http.Request) {

	fruitArray := []string{"apple", "orange", "mango"}

	uadmin.RenderHTML(w, r, "./templates/index.html", map[string]interface{}{
		"Title":    "HTML Templating",
		"Multiply": Multiply,
		"Fruits":   fruitArray,
	})
}
