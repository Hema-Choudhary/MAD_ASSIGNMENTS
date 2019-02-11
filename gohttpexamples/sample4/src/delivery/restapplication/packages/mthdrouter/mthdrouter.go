package mthdroutr

import (
	"fmt"
	"net/http"

	resputl "delivery/restapplication/packages/resputl"
)

//ServiceAPIHandler all http methods
type ServiceAPIHandler interface {
	GetOne(r *http.Request, id string) resputl.SrvcRes
	Get(r *http.Request) resputl.SrvcRes
	Put(r *http.Request) resputl.SrvcRes
	Post(r *http.Request) resputl.SrvcRes
	Delete(r *http.Request) resputl.SrvcRes
	Patch(r *http.Request) resputl.SrvcRes
	Options(r *http.Request) resputl.SrvcRes
}

//RouteAPICall routing to method
func RouteAPICall(sah ServiceAPIHandler, r *http.Request) resputl.SrvcRes {
	switch r.Method {
	case "GET":
		fmt.Println("in GET")
		// params := mux.Vars(r)
		// id, present := params["id"]
		// if present {
		// 	return sah.GetOne(r, id)
		// } else {
		// 	return sah.Get(r)
		// }
		return sah.Get(r)
	case "PUT":
		fmt.Println("in PUT")
		return sah.Put(r)
	case "POST":
	        fmt.Println("in POST")
		return sah.Post(r)
	case "PATCH":
		fmt.Println("in Patch")
		return sah.Patch(r)
	case "DELETE":
		fmt.Println("in Delete")
		return sah.Delete(r)
	case "OPTIONS":
		fmt.Println("in Options")
		return sah.Options(r)
	}
	return resputl.SrvcRes{
		Code:     http.StatusMethodNotAllowed,
		Response: fmt.Sprintf("{\"ResponseData\" : \"%s \"}", r.Method),
		Message:  "Method not allowed",
		Headers:  nil}

}
