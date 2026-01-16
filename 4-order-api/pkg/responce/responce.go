package responce

import (
	"encoding/json"
	"net/http"
)

type HttpReply struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func JsonOkResponce(w http.ResponseWriter, data any, statusCode int) {
	reply := HttpReply{
		Success: true,
		Error:   "",
		Data:    data,
	}
	jsonResponce(w, reply, statusCode)
}

func JsonErrorResponce(w http.ResponseWriter, error string, statusCode int) {
	reply := HttpReply{
		Success: false,
		Error:   error,
		Data:    nil,
	}
	jsonResponce(w, reply, statusCode)
}

func jsonResponce(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func TextResponce(w http.ResponseWriter, text string, statusCode int) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(statusCode)
	w.Write([]byte(text))
}
