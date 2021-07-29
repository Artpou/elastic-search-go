package respond

import (
	"encoding/json"
	"net/http"
)

func GetJson(w http.ResponseWriter, url string, target interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	data := json.NewDecoder(r.Body).Decode(target)
	json.NewEncoder(w).Encode(data)
	return nil
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}
