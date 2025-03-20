package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"lab2/utils"
	"path/filepath"
)

type CalculatorData struct {
	Emission      float64
	GrossEmission float64
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/coal", coalCalculator)
	http.HandleFunc("/oil", oilCalculator)
	http.HandleFunc("/gas", gasCalculator)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "main.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func coalCalculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		coalMassStr := r.FormValue("mass")
		coalMass, err := strconv.ParseFloat(coalMassStr, 64)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		emission := utils.CalculateEmission(20.47, 0.8, 25.20, 1.5)
		grossEmission := utils.CalculateGrossEmission(emission, coalMass, 20.47)

		data := CalculatorData{Emission: emission, GrossEmission: grossEmission}
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "coal.html"))
		tmpl.Execute(w, data)
	} else {
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "coal.html"))
		tmpl.Execute(w, nil)
	}
}

func oilCalculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		oilMassStr := r.FormValue("mass")
		oilMass, err := strconv.ParseFloat(oilMassStr, 64)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		emission := utils.CalculateEmission(39.48, 1.0, 0.15, 0.0)
		grossEmission := utils.CalculateGrossEmission(emission, oilMass, 39.48)

		data := CalculatorData{Emission: emission, GrossEmission: grossEmission}
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "oil.html"))
		tmpl.Execute(w, data)
	} else {
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "oil.html"))
		tmpl.Execute(w, nil)
	}
}

func gasCalculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		gasMassStr := r.FormValue("mass")
		gasMass, err := strconv.ParseFloat(gasMassStr, 64)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		emission := utils.CalculateEmission(33.08, 1.0, 0.723, 0.5)
		grossEmission := utils.CalculateGrossEmission(emission, gasMass, 33.08)

		data := CalculatorData{Emission: emission, GrossEmission: grossEmission}
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "gas.html"))
		tmpl.Execute(w, data)
	} else {
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "gas.html"))
		tmpl.Execute(w, nil)
	}
}
