package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/tanoshime/v2rpc/src/api"
)

func main() {
	restful.Filter(restful.OPTIONSFilter())
	restful.Filter(func(r1 *restful.Request, r2 *restful.Response, fc *restful.FilterChain) {
		log.Println("ProcessFilter")
		r2.Header().Set("Access-Control-Allow-Origin", r1.HeaderParameter("origin"))
		r2.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		r2.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		fc.ProcessFilter(r1, r2)
	})
	restful.Add(api.NewApi())
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Println("listen at: http://127.0.0.1:12580")
	http.ListenAndServe("127.0.0.1:12580", nil)
}
