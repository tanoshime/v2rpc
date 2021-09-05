package api

import (
	"fmt"
	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/tanoshime/v2rpc/src/utils"
)

func NewApi() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/api/{target}").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("/stats").To(queryStats))
	service.Route(service.GET("/stats/inbound/{tag}/{traffic-type}").To(inboundStat))
	service.Route(service.GET("/stats/user/{email}/{traffic-type}").To(userStat))
	service.Route(service.POST("/{tag}/user").To(addUser))
	service.Route(service.DELETE("/{tag}/user/{email}").To(removeUser))
	return service
}

func removeUser(request *restful.Request, response *restful.Response) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		m := make(map[string]string)
		m["message"] = fmt.Sprintf("%v", r)
		response.WriteAsJson(m)
	}()
	target := request.PathParameter("target")
	h := utils.NewRPCHelper(target)
	tag := request.PathParameter("tag")
	email := request.PathParameter("email")
	resp := h.RemoveUser(tag, email)
	m := make(map[string]string)
	m["message"] = resp
	response.WriteAsJson(m)
}

func addUser(request *restful.Request, response *restful.Response) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		m := make(map[string]string)
		message := fmt.Sprintf("%v", r)
		m["message"] = message
		response.WriteAsJson(m)
	}()
	target := request.PathParameter("target")
	h := utils.NewRPCHelper(target)
	vmessUser := new(utils.VmessUser)
	readError := request.ReadEntity(&vmessUser)
	if readError != nil {
		response.WriteAsJson(readError)
		return
	}
	tag := request.PathParameter("tag")
	err := h.AddVmessUser(tag, *vmessUser)
	m := make(map[string]string)
	if err != nil {
		m["message"] = err.Error()
		response.WriteAsJson(m)
	} else {
		response.WriteAsJson(vmessUser)
	}
}

func userStat(request *restful.Request, response *restful.Response) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		m := make(map[string]string)
		m["message"] = fmt.Sprintf("%v", r)
		response.WriteAsJson(m)
	}()
	target := request.PathParameter("target")
	h := utils.NewRPCHelper(target)
	email := request.PathParameter("email")
	traffic := request.PathParameter("traffic-type")
	reset, _ := strconv.ParseBool(request.QueryParameter("reset"))
	resp, err := h.GetUserTraffic(email, traffic, reset)
	if err != nil {
		panic(err)
	}
	response.WriteAsJson(resp)
}

func inboundStat(request *restful.Request, response *restful.Response) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		m := make(map[string]string)
		m["message"] = fmt.Sprintf("%v", r)
		response.WriteAsJson(m)
	}()
	target := request.PathParameter("target")
	h := utils.NewRPCHelper(target)
	tag := request.PathParameter("tag")
	traffic := request.PathParameter("traffic-type")
	reset, _ := strconv.ParseBool(request.QueryParameter("reset"))
	resp, _ := h.GetInboundTraffic(tag, traffic, reset)
	response.WriteAsJson(resp)
}

func queryStats(request *restful.Request, response *restful.Response) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		m := make(map[string]string)
		m["message"] = fmt.Sprintf("%v", r)
		response.WriteAsJson(m)
	}()
	target := request.PathParameter("target")
	h := utils.NewRPCHelper(target)
	pattern := request.QueryParameter("pattern")
	reset, _ := strconv.ParseBool(request.QueryParameter("reset"))
	resp, _ := h.QueryStats(pattern, reset)
	response.WriteAsJson(resp)
}
