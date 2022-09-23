// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/teru-0529/go_webapi_sandbox/config"
)

// main
func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %+v", err)
		os.Exit(1)
	}
}

// run
func run(ctx context.Context) error {
	// 環境変数の読込み
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// ポートの確認
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %+v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	// HTTPサーバーの起動
	mux := NewMux()
	srv := NewServer(l, mux)
	return srv.Run(ctx)
}
