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

func TestRun(t *testing.T) {
	t.Skip("リファクタリング中") // TODO:

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
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
