package logic

import "net/http"

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		Put(w, r)
		return
	}
	if m == http.MethodGet {
		Get(w, r)
		return
	}
	if m == http.MethodDelete {
		Delete(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}
