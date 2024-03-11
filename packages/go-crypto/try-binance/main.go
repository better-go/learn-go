package main

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/k0kubun/pp/v3"
	"os"
	"reflect"
	"strings"
)

func main() {
	println("try-binance")
	tryBN()
}

func tryBN() {
	var (
		apiKey    = "your api key"
		secretKey = "your secret key"
	)

	apiKey = os.Getenv("BINANCE_API_KEY")
	secretKey = os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		pp.Printf("Binance API KEY or SECRET KEY is empty, please set env: BN_API_KEY and BN_SECRET_KEY")
		return
	}

	pp.Printf("binance apiKey: %s, secretKey: %s\n", apiKey, secretKey)

	// create client
	client := binance.NewClient(apiKey, secretKey)
	futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

	// print time
	serverTime, err := client.NewServerTimeService().Do(context.Background())
	pp.Printf("server(Spot) time: %v, err: %v\n", serverTime, err)

	fTime, err := futuresClient.NewServerTimeService().Do(context.Background())
	pp.Printf("server(Futures) time: %v, err: %v\n", fTime, err)

	dTime, err := deliveryClient.NewServerTimeService().Do(context.Background())
	pp.Printf("server(Delivery) time: %v, err: %v\n", dTime, err)

	// get account info
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// struct pointer
	t := reflect.TypeOf(res).Elem()
	v := reflect.ValueOf(res).Elem()

	pp.Println("user account info:")
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			//判断是否是嵌套结构, skip balance field
			if v.Field(i).Type().Kind() != reflect.Slice {
				pp.Printf("[%v, %v]: %+v\n", t.Field(i).Name, t.Field(i).Type.Name(), v.Field(i).Interface())
			}

			//判断是否是嵌套结构, skip balance field
			//pp.Printf("%s %s = %v -tag:%s \n",
			//	t.Field(i).Name,
			//	t.Field(i).Type,
			//	v.Field(i).Interface(),
			//	t.Field(i).Tag)
		}
	}

	pp.Println("user balance:")
	if res.Balances != nil {
		for _, b := range res.Balances {
			fb := strings.TrimLeft(strings.TrimRight(b.Free, "0"), "0")
			lb := strings.TrimLeft(strings.TrimRight(b.Locked, "0"), "0")

			// balance > 0
			if fb != "." || lb != "." {
				pp.Printf("%+v\n", b)
			}
		}

		// list open orders
		pp.Println("open orders:")
		symbol := "DOTFDUSD"
		symbol = "" // get all symbols orders
		openOrders, err := client.NewListOpenOrdersService().Symbol(symbol).
			Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, o := range openOrders {
			pp.Printf("%+v\n", o)
		}

		// get order:
	}
}
