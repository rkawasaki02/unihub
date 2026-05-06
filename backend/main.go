package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rkawasaki02/unihub/backend/config"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	// TODO: Phase 2でルート登録
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("unihub server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
