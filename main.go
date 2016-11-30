package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/webercoder/go-mtg-service/lib"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rc := &lib.RequestController{}
		decoder := json.NewDecoder(r.Body)

		var req lib.Request
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Fprintf(w, "Difficulty parsing HipChat request: %+v", err)
			return
		}

		rsp, err := rc.HandleRequest(&req)

		if err != nil {
			fmt.Fprintf(w, "Error finding your card: %+v", err)
			return
		}

		b, err := json.Marshal(rsp)
		if err != nil {
			fmt.Fprintf(w, "Error creating JSON response for %+v: %+v", rsp, err)
			return
		}

		fmt.Fprintf(w, "%s", string(b))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
