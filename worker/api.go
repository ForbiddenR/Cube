package worker

import (
	"fmt"
	"net/http"
)

type ErrResponse struct {
	HttpStatusCode int
	Message        string
}

type Api struct {
	Address string
	Port    int
	Worker  *Worker
	Router  *http.ServeMux
}

func (a *Api) initRouter() {
	a.Router = http.NewServeMux()
	a.Router.HandleFunc("POST /tasks", a.StartTaskHandler)
	a.Router.HandleFunc("GET /tasks", a.GetTasksHandler)
	a.Router.HandleFunc("DELETE /tasks/{taskID}", a.StopTaskHandler)
}

func (a *Api) Start() {
	a.initRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%d", a.Address, a.Port), a.Router)
}
