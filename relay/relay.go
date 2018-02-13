package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"
)

const teem_url string = "https://push.teem.com"

const url_endpoint string = "/connect/google/pushcb/"

func sendRelay(headers map[string][]string, watch_id string) bool {
    url := fmt.Sprintf("%v%v%v/", teem_url, url_endpoint, watch_id)
    client := &http.Client{}
    request, _ := http.NewRequest("POST", url, nil)
    for key, value := range headers {
        request.Header.Add(key, value[0])
    }
    resp, _ := client.Do(request)
    if resp.StatusCode == 200 {
        return true
    } else {
        return false
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
    watch_id := strings.TrimSuffix(r.URL.Path[len(url_endpoint):], "/")
    relay_header := make(map[string][]string)
    for key, value := range r.Header {
        if strings.HasPrefix(key, "X-") {
            relay_header[key] = value
        }
    }
    response_json := map[string]string{"status": "success"}
    defer sendRelay(relay_header, watch_id)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(response_json)
}

func main() {
    s := &http.Server{
        Addr: ":443",
    }
    http.HandleFunc(url_endpoint, handler)
    log.Fatal(s.ListenAndServeTLS("server.crt", "server.key"))
}
