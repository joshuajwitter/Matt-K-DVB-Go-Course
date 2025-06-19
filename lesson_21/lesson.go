package lesson_21

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type Item struct {
	Name  string  `json:"name"`
	Price dollars `json:"price"`
}

// NOTE: don't do this in real life
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// Response is for add, update, fetch, drop handlers
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// ListResponse is for list handler
type ListResponse struct {
	Data  []Item `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// writeJSON writes a JSON response with the given status code
func writeJSON(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("failed to write JSON response", "error", err)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	slog.Info("handling list request", "method", req.Method, "url", req.URL.String())

	if req.Method != http.MethodGet {
		slog.Warn("invalid method", "method", req.Method)
		writeJSON(w, http.StatusMethodNotAllowed, ListResponse{Error: "method must be GET"})
		return
	}

	items := make([]Item, 0, len(db))
	for item, price := range db {
		items = append(items, Item{Name: item, Price: price})
	}

	slog.Info("listing items", "count", len(items))
	if len(items) == 0 {
		writeJSON(w, http.StatusOK, ListResponse{Error: "no items in database"})
		return
	}
	writeJSON(w, http.StatusOK, ListResponse{Data: items})
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	slog.Info("handling add request", "method", req.Method, "url", req.URL.String())

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" {
		slog.Warn("missing item parameter")
		writeJSON(w, http.StatusBadRequest, Response{Error: "item must be specified"})
		return
	}
	if price == "" {
		slog.Warn("missing price parameter")
		writeJSON(w, http.StatusBadRequest, Response{Error: "price must be specified"})
		return
	}
	if _, ok := db[item]; ok {
		slog.Warn("item already exists", "item", item)
		writeJSON(w, http.StatusBadRequest, Response{Error: fmt.Sprintf("item already exists: %s", item)})
		return
	}

	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		slog.Warn("invalid price", "price", price, "error", err)
		writeJSON(w, http.StatusBadRequest, Response{Error: fmt.Sprintf("invalid price: %s", price)})
		return
	}
	if floatPrice < 0 {
		slog.Warn("negative price", "price", floatPrice)
		writeJSON(w, http.StatusBadRequest, Response{Error: "price cannot be negative"})
		return
	}

	dollarsPrice := dollars(floatPrice)
	db[item] = dollarsPrice

	slog.Info("item added", "item", item, "price", dollarsPrice)
	writeJSON(w, http.StatusCreated, Response{Message: fmt.Sprintf("item %s added with price %s", item, dollarsPrice)})
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	slog.Info("handling update request", "method", req.Method, "url", req.URL.String())

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" {
		slog.Warn("missing item parameter")
		writeJSON(w, http.StatusBadRequest, Response{Error: "item must be specified"})
		return
	}
	if price == "" {
		slog.Warn("missing price parameter")
		writeJSON(w, http.StatusBadRequest, Response{Error: "price must be specified"})
		return
	}
	if _, ok := db[item]; !ok {
		slog.Warn("item not found", "item", item)
		writeJSON(w, http.StatusNotFound, Response{Error: fmt.Sprintf("item not found: %s", item)})
		return
	}

	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		slog.Warn("invalid price", "price", price, "error", err)
		writeJSON(w, http.StatusBadRequest, Response{Error: fmt.Sprintf("invalid price: %s", price)})
		return
	}
	if floatPrice < 0 {
		slog.Warn("negative price", "price", floatPrice)
		writeJSON(w, http.StatusBadRequest, Response{Error: "price cannot be negative"})
		return
	}

	dollarsPrice := dollars(floatPrice)
	db[item] = dollarsPrice

	slog.Info("item updated", "item", item, "price", dollarsPrice)
	writeJSON(w, http.StatusOK, Response{Message: fmt.Sprintf("item %s updated with price %s", item, dollarsPrice)})
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	slog.Info("handling fetch request", "method", req.Method, "url", req.URL.String())

	if req.Method != http.MethodGet {
		slog.Warn("invalid method", "method", req.Method)
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "method must be GET"})
		return
	}

	item := req.URL.Query().Get("item")

	if item == "" {
		slog.Warn("missing item parameter")
		writeJSON(w, http.StatusBadRequest, Response{Error: "item must be specified"})
		return
	}
	price, ok := db[item]
	if !ok {
		slog.Warn("item not found", "item", item)
		writeJSON(w, http.StatusNotFound, Response{Error: fmt.Sprintf("item not found: %s", item)})
		return
	}

	slog.Info("item fetched", "item", item, "price", price)
	writeJSON(w, http.StatusOK, Response{Message: fmt.Sprintf("item %s has price %s", item, price)})
}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	slog.Info("handling drop request", "method", req.Method, "url", req.URL.String())

	item := req.URL.Query().Get("item")

	if item == "" {
		slog.Warn("missing item parameter")
		writeJSON(w, http.StatusBadRequest, Response{Error: "item must be specified"})
		return
	}
	if _, ok := db[item]; !ok {
		slog.Warn("item not found", "item", item)
		writeJSON(w, http.StatusNotFound, Response{Error: fmt.Sprintf("item not found: %s", item)})
		return
	}

	delete(db, item)

	slog.Info("item deleted", "item", item)
	writeJSON(w, http.StatusOK, Response{Message: fmt.Sprintf("item %s deleted", item)})
}

func DoLesson() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)
	http.HandleFunc("/read", db.fetch)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
