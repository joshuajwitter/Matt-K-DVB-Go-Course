package lesson_15

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// anything not included will be ignored
type Todo struct {
	// fields must have uppercase letter or JSON won't handle them
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Todo #{{.ID}}</title>
</head>
<body>
    <h1>Todo #{{.ID}}</h1>
    <p>Title: {{.Title}}</p>
    <p>Completed: {{.Completed}}</p>
</body>
</html>`

var tmpl = template.Must(template.New("todo").Parse(formTemplate))

func DoLesson() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		startServer()
	}()

	time.Sleep(100 * time.Millisecond)
	startClient()

	wg.Wait()
}

func startServer() {
	fmt.Println("Server listening on :8080")

	// handle a specific route
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

// create a basic handler
func handler(w http.ResponseWriter, r *http.Request) {

	const base = "https://jsonplaceholder.typicode.com"
	var fullUrl = base + r.URL.Path[1:]

	todo, err := fetchTodoFromUrl(fullUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, todo); err != nil {
		log.Printf("template execute error: %v", err)
	}
}

func startClient() {
	fmt.Println("Starting the client")
	resp, err := http.Get("http://localhost:8080/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()
}

func fetchTodoFromUrl(url string) (Todo, error) {

	resp, err := http.Get(url)
	if err != nil {
		return Todo{}, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	// Defer closing the body, capturing any potential Close() error
	var closeErr error
	defer func() {
		if err := resp.Body.Close(); err != nil {
			closeErr = fmt.Errorf("failed to close response body: %w", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return Todo{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Todo{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var item Todo
	if err := json.Unmarshal(body, &item); err != nil {
		return Todo{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Check if there was a close error before returning
	if closeErr != nil {
		return Todo{}, closeErr
	}

	return item, nil
}
