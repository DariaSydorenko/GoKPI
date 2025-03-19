package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Calculator1Input struct {
	Hp, Cp, Sp, Np, Op, Wp, Ap float64
}

type Calculator2Input struct {
	Hr, Cr, Sr, Or, Vr, Wr, Ar, Qdafi float64
}

type Calculator1Response struct {
	Kpc, Kpg                   float64
	Hc, Cc, Sc, Nc, Oc, Ac         float64
	Hr, Cr, Sr, Nr, Or         float64
	Qph, Qch, Qrh              float64
}

type Calculator2Response struct {
	Hp, Cp, Sp, Op, Vp, Ap     float64
	Qri												 float64
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calculator1", calculator1Handler)
	http.HandleFunc("/calculator2", calculator2Handler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Сервер запущено на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmplPath := "templates/" + tmpl + ".html"
	t, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, "Помилка завантаження шаблону", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func calculator1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "calculator1")
		return
	}

	if r.Method == http.MethodPost {
		var input Calculator1Input

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Невірні дані", http.StatusBadRequest)
			return
		}

		Kpc := 100 / (100 - input.Wp)
		Kpg := 100 / (100 - input.Wp - input.Ap)

		Hc := input.Hp * Kpc
		Cc := input.Cp * Kpc
		Sc := input.Sp * Kpc
		Nc := input.Np * Kpc
		Oc := input.Op * Kpc
		Ac := input.Ap * Kpc

		Hr := input.Hp * Kpg
		Cr := input.Cp * Kpg
		Sr := input.Sp * Kpg
		Nr := input.Np * Kpg
		Or := input.Op * Kpg

		Qph := (339 * input.Cp + 1030 * input.Hp - 108.8 * (input.Op-input.Sp) - 25 * input.Wp) / 1000
		Qch := (Qph + 0.025 * input.Wp) * (100 / (100 - input.Wp))
		Qrh := (Qph + 0.025 * input.Wp) * (100 / (100 - input.Wp - input.Ap))

		response := Calculator1Response{Kpc, Kpg, Hc, Cc, Sc, Nc, Oc, Ac, Hr, Cr, Sr, Nr, Or, Qph, Qch, Qrh}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func calculator2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "calculator2")
		return
	}

	if r.Method == http.MethodPost {
		var input Calculator2Input
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Невірні дані", http.StatusBadRequest)
			return
		}

		Hp := input.Hr * (100 - input.Wr - input.Ar) / 100
		Cp := input.Cr * (100 - input.Wr - input.Ar) / 100
		Sp := input.Sr * (100 - input.Wr - input.Ar) / 100
		Op := input.Or * (100 - input.Wr - input.Ar) / 100
		Vp := input.Vr * (100 - input.Wr) / 100
		Ap := input.Ar * (100 - input.Wr) / 100
		Qri := input.Qdafi * ((100 - input.Wr - input.Ar) / 100) - 0.025 * input.Wr

		response := Calculator2Response{Hp, Cp, Sp, Op, Vp, Ap, Qri}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
