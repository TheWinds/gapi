package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/zhiduoke/gapi"
	"github.com/zhiduoke/gapi/examples/greet/httpapi"
	"github.com/zhiduoke/gapi/examples/greet/service"
	"github.com/zhiduoke/gapi/examples/greet/updater"
	"github.com/zhiduoke/gapi/handler/httpjson"
	"github.com/zhiduoke/gapi/handler/httpview"
	"github.com/zhiduoke/gapi/metadata"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

var mdUpdater *updater.Updater

func main() {
	mdUpdater, _ = updater.NewUpdater(nil)
	go startHTTPServer()
	time.Sleep(time.Second)
	startRPCServer()
}

const rpcServerTarget = "127.0.0.1:33333"

func startHTTPServer() {
	apisrv := gapi.NewServer()
	apisrv.RegisterHandler("httpjson", &httpjson.Handler{})
	apisrv.RegisterHandler("httpview", &httpview.Handler{})

	apisrv.Use(func(ctx *gapi.Context) error {
		err := ctx.Next()
		if err != nil {
			ctx.Response().Write([]byte(err.Error()))
		}
		return nil
	})
	apisrv.RegisterMiddleware("auth", func(ctx *gapi.Context) error {
		//ctx.Request().ParseForm()
		// mock some user id
		//if ctx.Request().Form.Get("name") == "gapi" {
		ctx.Set("uid", "666")
		//}
		return ctx.Next()
	})

	// define resolve func
	apisrv.Dial = func(s string) (*grpc.ClientConn, error) {
		switch s {
		case "service.demo":
			return grpc.Dial(rpcServerTarget, grpc.WithInsecure())
		default:
			return nil, errors.New("unkonwn server name")
		}
	}

	// set a metadata updater
	mdUpdater.SetChangedFn(func(md *metadata.Metadata) {
		for key, value := range md.Routes {
			log.Println(key, value.Path)
		}
		if err := apisrv.UpdateRoute(md); err != nil {
			log.Fatal(err)
		}
	})
	e := echo.New()
	conn, err := grpc.Dial(rpcServerTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	e.POST("/v1/demo/greet", func(ctx echo.Context) error {
		client := httpapi.NewDemoAPIClient(conn)
		uid, _ := strconv.Atoi(ctx.Get("uid").(string))
		resp, err := client.Greet(context.Background(), &httpapi.GreetRequest{
			UserId: int64(uid),
			Name:   ctx.FormValue("name"),
			To:     nil,
		})
		if err != nil {
			return err
		}
		return ctx.JSON(200, resp)
	}, func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("uid", "666")
			return h(ctx)
		}
	})
	go http.ListenAndServe(":1424", e)
	http.ListenAndServe(":1323", apisrv)
}

func startRPCServer() {
	lis, err := net.Listen("tcp", rpcServerTarget)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	httpapi.RegisterDemoAPIServer(s, &service.DemoServer{})
	go func() {
		err := mdUpdater.ReportChange("demo.proto", proto.FileDescriptor("demo.proto"))
		if err != nil {
			log.Fatal(err)
		}
	}()
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
