package main

import (
	"fmt"
	"net/http"
	"os"
)

func meeting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	http.StatusText(400)
	if err != nil {
		http.StatusText(400)
		return
	}
	loc := r.Form["loc"]
	fmt.Println(loc)
	//center := centerOfCoords(loc)
	//toMeet = findMeetingPlaces(center)
}

func main() {
	http.HandleFunc("/meeting", meeting)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
