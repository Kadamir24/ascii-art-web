package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	ascii "./ascii"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Fonts struct {
	FontName []string `json:"font"`
}

func getFonts(w http.ResponseWriter, r *http.Request) {
	var fontstruct Fonts
	fontstruct = Fonts{
		FontName: []string{"standard", "shadow", "thinkertoy"},
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(fontstruct)
}

type Index struct {
	Title     string
	Body      string
	TeamMates string
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	bodyByte, _ := ioutil.ReadFile("templates/bodyform.html")
	body := string(bodyByte)
	footer, _ := ioutil.ReadFile("templates/footer.html")
	teammates := string(footer)
	page := Index{
		Title:     "Ascii-Art-Web",
		Body:      body,
		TeamMates: teammates,
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, page)
}

type Config struct {
	Text string `json:"text"`
	Font string `json:"font"`
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	var cfg Config
	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	text, es := ascii.AsciiCall(cfg.Text, cfg.Font)
	if es > 200 {
		errorHandler(w, r, es)
		return
	}
	fmt.Fprint(w, text)

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	// err, e := ascii.AsciiCall(strconv.Itoa(status)+" "+http.StatusText(status), "standard")
	if status == 500 {
		fmt.Fprint(w, "500 Internal Server Error")
		return
	} else if status == 400 {
		fmt.Fprint(w, "400 Bad Request")
		return
		// } else if status == 404 {
		// 	fmt.Fprint(w, "404 Not Found")
		// 	return
	}
	tmpl, _ := template.ParseFiles("templates/error.html")
	tmpl.Execute(w, status)
}

func main() {
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/fontsApi", getFonts)
	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/ascii/", asciiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
