package main

import (
	"fmt"
	"go-proxy/service"
	"net/http"
	"time"
)

func main() {
	h := service.NewProxy()

	// create a http server
	http.HandleFunc("/proxy", h.ProxyAddressRequest)

	// start the main server
	port := h.Port
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on %s\n", serverAddr)
	server := &http.Server{
		Addr:           serverAddr,
		Handler:        nil, // use default Handler
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Failed to start server: %s\n", err)
		}
	}()

	select {}
}
