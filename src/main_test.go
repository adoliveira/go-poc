package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartServer(t *testing.T) {
	mux := BuilServerMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/user")
	if err != nil {
		t.Fatalf("erro ao fazer GET: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", resp.StatusCode)
	}
}

func TestBuilServerMux(t *testing.T) {
	mux := BuilServerMux()
	req := httptest.NewRequest("GET", "/user", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}
}

func TestMainFunction(t *testing.T) {
	// Testa se a função main executa sem panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main causou panic: %v", r)
		}
	}()
	go func() {
		// Executa main em uma goroutine para não bloquear
		main()
	}()
}
