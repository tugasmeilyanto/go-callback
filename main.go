package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type CallbackResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Process the callback request here
        var data map[string]interface{}
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Log the received data for debugging
        log.Printf("Received callback data: %+v\n", data)

        // Respond with a 2XX status code
        response := CallbackResponse{
            Status:  "success",
            Message: "Callback received successfully",
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/callback", callbackHandler)
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
