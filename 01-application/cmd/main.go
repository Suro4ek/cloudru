package main

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
)

func main() {
	type Response struct {
		Hostname string `json:"hostname"`
		IP       string `json:"ip"`
		Author   string `json:"author"`
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			http.Error(w, "err", http.StatusInternalServerError)
			return
		}

		addrs, err := net.InterfaceAddrs()
		if err != nil {
			http.Error(w, "err", http.StatusInternalServerError)
			return
		}

		var ip string
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}

		if ip == "" {
			ip = "not found"
		}

		author := os.Getenv("AUTHOR")

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(Response{
			Hostname: hostname,
			IP:       ip,
			Author:   author,
		})
	})

	http.ListenAndServe(":8080", nil)
}
