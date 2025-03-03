// I had to upload something xd

package main

import (
	kl "api/kartos/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

type startParams struct {
	address string
	port    int16
}

func main() {
	var startParams startParams = startParams{"localhost", 81}
	c := color.New(color.FgYellow)
	c.Println(`
 __                   __                                       .__ 
|  | _______ ________/  |_  ____  ______         _____  ______ |__|
|  |/ /\__  \\_  __ \   __\/  _ \/  ___/  ______ \__  \ \____ \|  |
|    <  / __ \|  | \/|  | (  <_> )___ \  /_____/  / __ \|  |_> >  |
|__|_ \(____  /__|   |__|  \____/____  >         (____  /   __/|__|
     \/     \/                       \/               \/|__|       
 
	`)
	kl.LogWithTimer("API Server started! 🚀")
	kl.LogWithTimer("Network: http://" + startParams.address + ":" + strconv.Itoa(int(startParams.port)))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /test/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "API WORKS %s", id)
	})

	// mux.HandleFunc("POST /test/", func(w http.ResponseWriter, r *http.Request) {
	// 	body, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		http.Error(w, "Error", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	// err := json.NewDecoder(bodys)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	fmt.Fprintf(w, "%s", name)
	// })

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Wrong usage of API!")
	})
	err := http.ListenAndServe(startParams.address+":"+strconv.Itoa(int(startParams.port)), mux)
	if err != nil {
		kl.LogWithTimer("Server failed to start: " + err.Error())
	}

}
