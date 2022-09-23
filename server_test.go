// main_test.go
package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	eg.Go(func() error {
		srv := NewServer(l, mux)
		return srv.Run(ctx)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// リッスンポートの表示
	t.Logf("try request to %q", url)

	res, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer res.Body.Close()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read body: %+v", err)
	}

	// HTTPサーバーの戻り値の検証
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run関数の終了通知処理検証
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
