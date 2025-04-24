// main.go
package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

type Request struct {
	IP               string   `json:"ip"`
	AllowedCountries []string `json:"allowed_countries"`
}

type Response struct {
	Allowed bool `json:"allowed"`
}

func main() {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/ip-authorization", func(w http.ResponseWriter, r *http.Request) {
		var req Request

		// Declare new json decoder
		decoder := json.NewDecoder(r.Body)

		// At my current company we treat payloads with extraneous fields as invalid payloads
		decoder.DisallowUnknownFields()

		// Decode will automatically unmarshal request into req Request struct, if nil or error immediately throw invalid request
		if err := decoder.Decode(&req); err != nil { //
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// ip will be nil if invalid IP format
		ip := net.ParseIP(req.IP)

		record, err := db.Country(ip)
		if err != nil {
			http.Error(w, "IP lookup failed", http.StatusInternalServerError)
			return
		}

		countryCode := record.Country.IsoCode
		allowed := false
		// Since the list of allowed countries is small (I'm guessing going to be < 100)
		// this for loop is performant. However, if the list somehow became much larger or if there were additional logic
		// needed for each request and performance considerations were needed this would be an area to look at.
		for _, code := range req.AllowedCountries {
			if code == countryCode {
				allowed = true
				break
			}
		}

		resp := Response{Allowed: allowed}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
