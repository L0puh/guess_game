package main

import (
	"bufio"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/tarm/serial"
)

var tmp *template.Template
var score_file string = "../data/score.txt"

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

func reset_score() {
	file, err := os.Create(score_file)
	print_error(err)
	defer file.Close()
	_, err = file.WriteString("0")
	print_error(err)
	log.Println("Updated score")
}
func print_error(err error){
	if err != nil{
		log.Fatal(err);
	}
}
func get_score() int{
	file, err := os.Open(score_file);
	print_error(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var score int;
	for scanner.Scan() {
		text := scanner.Text()
		score, err = strconv.Atoi(text)
		if err != nil { score = 0 }
	}
	return score
}
func send_start_signal(){
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	print(err)
	n, err := s.Write([]byte("START"))
	print_error(err)
	log.Printf("Signal - start. Sent %d bytes\n", n)
}

func index_page(rw http.ResponseWriter, r* http.Request){
	score := get_score()
	if r.Method == "POST" {
		err := r.ParseForm()
		print_error(err)
		action := r.FormValue("btn")
		switch action {
		case "START":
			send_start_signal()
			break
		case "RESET":
			reset_score()
			break
		default:
			log.Println("UNKNOWN ACTION")
		}
		http.Redirect(rw, r, "/", http.StatusFound);
	}
	tmp.ExecuteTemplate(rw, "index.html", score)
}
