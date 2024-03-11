package main

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"log"
	"os"
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
		log.Fatal("Binance API KEY or SECRET KEY is empty, please set env: BN_API_KEY and BN_SECRET_KEY")
		return
	}

	log.Printf("binance apiKey: %s, secretKey: %s\n", apiKey, secretKey)

	// create client
	client := binance.NewClient(apiKey, secretKey)
	futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

	// print time
	serverTime, err := client.NewServerTimeService().Do(context.Background())
	log.Printf("server(Spot) time: %v, err: %v\n", serverTime, err)
	fTime, err := futuresClient.NewServerTimeService().Do(context.Background())
	log.Printf("server(Futures) time: %v, err: %v\n", fTime, err)
	dTime, err := deliveryClient.NewServerTimeService().Do(context.Background())
	log.Printf("server(Delivery) time: %v, err: %v\n", dTime, err)

	// list open orders

	symbol := "DOTFDUSD"
	openOrders, err := client.NewListOpenOrdersService().Symbol(symbol).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range openOrders {
		fmt.Println(o)
	}

	// get order:

}
