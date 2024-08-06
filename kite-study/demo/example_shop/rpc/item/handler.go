package main

import (
	"context"
	item "example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/stock/stockservice"
	"github.com/cloudwego/kitex/client"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct {
	stockCli stockservice.Client
}

func NewStockClient(addr string) (stockservice.Client, error) {
	return stockservice.NewClient("example.shop.stock", client.WithHostPorts(addr))
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemReq) (resp *item.GetItemResp, err error) {
	resp = item.NewGetItemResp()
	resp.Item = item.NewItem()
	resp.Item.Id = req.GetId()
	resp.Item.Title = "Kitex"
	resp.Item.Description = "Kitex is an excellent framework!"
	return
}
