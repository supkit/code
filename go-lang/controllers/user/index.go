package user

import (
	"fmt"
	"net/http"
)

func Index(res http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintln("URL", r.URL, "HOST", r.Host, "Method", r.Method, "RequestURL", r.RequestURI, "RawQuery", r.URL.RawQuery)
	fmt.Fprintln(res, info)
	res.Write([]byte("User Index"))
}
