package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(request, para string) *Response {
	//return "ECHO:" + request
	return &Response{Code: "200"}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("From Client1", "From server")
	resp2, _ := client1.Call("From Client2", "From server")
	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1.Code, "resp2:", resp2.Code)
	}
	client1.Close()
	client2.Close()
}
