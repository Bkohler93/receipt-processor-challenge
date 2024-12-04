package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *config) handlePostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	var rr receiptRequest

	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		fmt.Printf("unable to decode request body into receipt - %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}

	receipt, err := rr.validateReceipt()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	receipt.calculatePoints()

	newReceipt, err := c.store.AddReceipt(receipt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not save receipt - %v", err)))
		return
	}

	res := struct {
		Id string `json:"id"`
	}{
		Id: newReceipt.Uuid.String(),
	}

	rb, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rb)
}

func (c *config) handleGetReceiptsPoints(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	notFoundResponse := "No receipt found for that id"

	uuid, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(notFoundResponse))
		return
	}

	receipt, err := c.store.GetReceipt(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(notFoundResponse))
	}

	res := struct {
		Points int `json:"points"`
	}{
		Points: receipt.Points,
	}

	rb, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rb)
}
