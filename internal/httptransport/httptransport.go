package httptransport

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"palatte-api/internal/service"
	"palettecalculator"
)

type handler struct {
	service.Calculator
}

func NewHandler(c service.Calculator) http.Handler {
	h := &handler{c}
	r := mux.NewRouter()
	r.HandleFunc("/predominant", h.calculatePredominant).Methods("POST")
	r.HandleFunc("/complimentary", h.calculateComplimentary).Methods("POST")
	r.HandleFunc("/split-complimentary", h.calculateSplitComplimentary).Methods("POST")
	r.HandleFunc("/triadic", h.calculateTriadic).Methods("POST")
	r.HandleFunc("/tetradic", h.calculateTetradic).Methods("POST")

	return r
}

func (h *handler) calculatePredominant(w http.ResponseWriter, r *http.Request) {
	var fp string
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&fp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, err := h.Calculator.CalculatePredominant(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	color, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(color)
}

func (h *handler) calculateComplimentary(w http.ResponseWriter, r *http.Request) {
	var inColor palettecalculator.Color
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&inColor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := h.Calculator.CalculateComplimentary(&inColor)

	color, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(color)
}
func (h *handler) calculateSplitComplimentary(w http.ResponseWriter, r *http.Request) {
	var inColor palettecalculator.Color
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&inColor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := h.Calculator.CalculateSplitComplimentary(&inColor)

	color, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(color)
}
func (h *handler) calculateTriadic(w http.ResponseWriter, r *http.Request) {
	var inColor palettecalculator.Color
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&inColor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := h.Calculator.CalculateTriadic(&inColor)

	color, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(color)
}
func (h *handler) calculateTetradic(w http.ResponseWriter, r *http.Request) {
	var inColor palettecalculator.Color
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&inColor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := h.Calculator.CalculateTetradic(&inColor)

	color, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(color)
}
