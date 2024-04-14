package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = 1819
	dbname   = "projects"
	sslmode  = "disable"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 2

type Photos struct {
	ID          int64  `json:"id" db:"id"`
	FileName    string `json:"filename" db:"filename"`
	Description string `json:"description" db:"description"`
}

var tmpl *template.Template

var dbase *sql.DB

// serve all html templates
func init() {
	tmpl = template.Must(template.ParseGlob("../view/templates/*.html"))
}

// execute home page (index.html)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

// Project handler (projects.html)
func projectHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := dbase.Query("SELECT * FROM links;")
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()
	allPhotos := []Photos{}

	for rows.Next() {
		p := Photos{}

		err := rows.Scan(&p.ID, &p.FileName, &p.Description)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		allPhotos = append(allPhotos, p)
	}

	tmpl.ExecuteTemplate(w, "projects.html", allPhotos)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "upload.html", nil)

}

func main() {

	// connection to postgres database
	connStr := fmt.Sprintf("user=%s dbname=%s password=%d sslmode=%s", user, dbname, password, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbase = db
	defer db.Close()

	// serving all files (like css, photos, js) in direction view
	fileServer := http.FileServer(http.Dir("../view"))
	http.Handle("/view/", http.StripPrefix("/view", fileServer))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/projects", projectHandler)

	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Server is running")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
