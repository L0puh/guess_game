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
	http.HandleFunc("/done", done_page);
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
func send_start_signal(speed string, repeat string){
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	print(err)

	data := fmt.Sprintf("%s,%s\n", speed, repeat)
	n, err := s.Write([]byte(data))
	print_error(err)
	log.Printf("Signal - start. Sent %d bytes\n", n)

}

func done_page(rw http.ResponseWriter, r* http.Request){
	tmp.ExecuteTemplate(rw, "done.html", nil);
}

func index_page(rw http.ResponseWriter, r* http.Request){
	score := get_score()
	if r.Method == "POST" {
		err := r.ParseForm()
		print_error(err)
		action := r.FormValue("btn")
		log.Println(action)
		switch action {
		case "START":
			speed := r.FormValue("speed")
			rep  := r.FormValue("repeat")
			send_start_signal(speed, rep)
			http.Redirect(rw, r, "/done", http.StatusFound);
			break
		case "RESET":
			reset_score()
			http.Redirect(rw, r, "/", http.StatusFound);
			break
		default:
			log.Println("UNKNOWN ACTION")
		}
	}
	tmp.ExecuteTemplate(rw, "index.html", score)
}
