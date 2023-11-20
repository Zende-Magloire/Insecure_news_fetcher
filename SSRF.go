package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running at http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			renderForm(w)
		} else if r.Method == "POST" {
			urlInput := r.FormValue("url")
			renderForm(w)
			if urlInput != "" {
				fetchAndDisplayContent(w, urlInput)
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}

func renderForm(w http.ResponseWriter) {
	form := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>News Fetcher</title>
	</head>
	<body>
		<h1>News Fetcher</h1>
		<form method="post" action="/">
			<label for="url">Enter news site URL:</label><br>
			<input type="text" id="url" name="url" placeholder="https://example.com"><br><br>
			<input type="submit" value="Fetch">
		</form>
	</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, form)
}

func fetchAndDisplayContent(w http.ResponseWriter, urlInput string) {
	// For demonstration purposes, instead of fetching the URL content,
	// we'll directly render the input URL without escaping HTML characters.
	// This introduces the XSS vulnerability.
	fmt.Fprintf(w, "<h2>Content from: %s</h2>", urlInput)
}
