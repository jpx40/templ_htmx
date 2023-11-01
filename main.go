package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	foot := foot()
	count := button()
	stencil := component()
	dropdown := dropdown()

	http.Handle("/", templ.Handler(layout(foot, count, stencil, dropdown)))

	http.ListenAndServe(":8080", nil)
}

// func render() {
// 	f, err := os.Create("hello.html")
// 	if err != nil {
// 		log.Fatalf("failed to create output file: %v", err)
// 	}

// 	err = layout().Render(context.Background(), f)
// 	if err != nil {
// 		log.Fatalf("failed to write output file: %v", err)
// 	}
// }
