package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	_ = godotenv.Load()

	c := cron.New()
	c.AddFunc("0 3 * * *", CheckIn)

	log.Println("Starting scheduler...")
	c.Start()
	log.Println("Started")

	// Wait for terminating signal
	end := make(chan os.Signal, 1)
	signal.Notify(end, syscall.SIGINT, syscall.SIGTERM)
	<-end

}

func CheckIn() {
	resp, err := resty.New().R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json;charset=UTF-8",
			"Cookie":       os.Getenv("COOKIE"),
		}).
		SetBody(map[string]string{
			"act_id": "e202102251931481",
		}).
		Post("https://hk4e-api-os.mihoyo.com/event/sol/sign?lang=vi-vn")
	if err != nil {
		log.Fatal("CheckIn error: ", err)
	}

	log.Println("OK: ", string(resp.Body()))
}
