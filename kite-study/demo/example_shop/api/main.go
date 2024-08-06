package main

import (
	"context"
	"example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/item/itemservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

var (
	cli itemservice.Client
)

func main() {
	c, err := itemservice.NewClient("example.shop.item", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	cli = c

	hertz := server.New(server.WithHostPorts("0.0.0.0:8889"))

	hertz.GET("/api/item", Handler)

	if err := hertz.Run(); err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx context.Context, c *app.RequestContext) {
	req := item.NewGetItemReq()
	req.Id = 1024
	resp, err := cli.GetItem(ctx, req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	c.String(200, resp.String())
}
