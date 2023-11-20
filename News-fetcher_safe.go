package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the user-provided URL from the query string
		urlStr := r.URL.Query().Get("url")

		// Parse the URL to check if it is valid
		u, err := url.Parse(urlStr)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
			// The URL is invalid or doesn't have a valid scheme, return an error
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		// Make a request to the user-provided URL
		client := http.Client{}
		resp, err := client.Get(u.String())
		if err != nil {
			// Handle the error
			http.Error(w, "Failed to fetch URL", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// Handle the error
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		w.Write(body)
	})

	http.ListenAndServe(":8080", nil)
}
