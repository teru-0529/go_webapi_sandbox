package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// JSONオブジェクトの検証
func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %+v", want, err)
	}
	if err := json.Unmarshal(got, jg); err != nil {
		t.Fatalf("cannot unmarshal got %q: %+v", got, err)
	}
	if diff := cmp.Diff(jg, jw); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

// レスポンスの検証
func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })

	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}

	// ステータスコードの検証
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}

	if len(gb) == 0 && len(body) == 0 {
		// レスポンスボディ無しの場合は終了
		return
	}

	// JSONオブジェクトの検証
	AssertJSON(t, body, gb)
}

// テストデータファイルの読み込み
func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read from %q: %+v", path, err)
	}
	return bt
}
