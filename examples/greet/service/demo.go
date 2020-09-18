package service

import (
	"context"
	"fmt"
	"github.com/zhiduoke/gapi/examples/greet/httpapi"
)

type DemoServer struct {
	count map[string]int32
	httpapi.UnimplementedDemoAPIServer
}

func (d *DemoServer) Greet(_ context.Context, req *httpapi.GreetRequest) (*httpapi.GreetReply, error) {
	return &httpapi.GreetReply{
		Msg: fmt.Sprintf("hello %s, your id is %d", req.Name, req.UserId),
	}, nil
}

func (d *DemoServer) GreetCount(_ context.Context, req *httpapi.GreetCountRequest) (*httpapi.GreetCountReply, error) {

	return &httpapi.GreetCountReply{
	}, nil
}
