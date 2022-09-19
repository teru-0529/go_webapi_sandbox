// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// main
func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %+v", err)
	}
}

// run
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)

	// 別ゴルーチンでHTTPサーバーを起動
	eg.Go(func() error {
		// http.ErrServerClosedは、http.Server.ShutDown()の正常終了であり異常ではない。
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// チャネルからの通知を待機
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	// 起動した別ゴルーチンの終了を待つ
	return eg.Wait()
}
