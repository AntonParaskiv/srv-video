package WebServer

import (
	"net/http"
)

type WebServer struct {
	muxServer  *http.ServeMux
	listenAddr string
}

const DefaultAddr = ":8080"

func New() (ws *WebServer) {
	ws = new(WebServer)
	ws.SetListenAddr(DefaultAddr)
	ws.SetServer(newMuxServer())
	return
}

func (ws *WebServer) SetListenAddr(listenAddr string) *WebServer {
	ws.listenAddr = listenAddr
	return ws
}

func (ws *WebServer) GetListenAddr() (addr string) {
	return ws.listenAddr
}

func (ws *WebServer) SetServer(muxServer *http.ServeMux) *WebServer {
	ws.muxServer = muxServer
	return ws
}

func newMuxServer() (muxServer *http.ServeMux) {
	muxServer = http.NewServeMux()
	return
}

func (ws *WebServer) HandleFunc(pattern string, serviceHandler func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request)) *WebServer {
	ws.muxServer.HandleFunc(pattern, serviceHandler)
	return ws
}

func (ws *WebServer) Start() (err error) {
	err = http.ListenAndServe(ws.listenAddr, ws.muxServer)
	return
}
