package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
)

var tmp *template.Template
var score_file string = "score.txt"

func main() {

	style := http.FileServer(http.Dir("styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", style))

	var err error;
	tmp, err = template.ParseGlob("templates/*.html");
	print_error(err);

	http.HandleFunc("/", index_page);
	fmt.Println("server is running. check localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",  nil))
}

func print_error(err error){
	if err != nil{
		log.Fatal(err);
	}
}
func get_score() []int{
	file, err := os.Open(score_file);
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	var scores []int
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		print_error(err)
		scores = append(scores, line)
	}
	return scores
}


func index_page(rw http.ResponseWriter, r* http.Request){
	scores := get_score()
	if r.Method == "POST" {
		http.Redirect(rw, r, "/", http.StatusFound);
	}
	best_score := slices.Max(scores)
	tmp.ExecuteTemplate(rw, "index.html", best_score)
}
