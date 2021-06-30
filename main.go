package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test-project/pkg/management"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	//подключить сервер управления
	http.HandleFunc("/", management.Handler())
	managementServer := &http.Server{
		Addr: ":" + management.Port,
	}
	defer managementServer.Close()
	go func() {
		if err := managementServer.ListenAndServe(); err != nil {
			fmt.Errorf("Ошибка сервера управления: ", err)
		}
	}()

	log.Info("Приложение запущено")

	<-signals
}