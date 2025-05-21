package gtp

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()

	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")

	if v2.prefix != "/v1/v2" {
		t.Fatalf("expected /v1/v2, but got %s", v2.prefix)
	}

	if v3.prefix != "/v1/v2/v3" {
		t.Fatalf("expected /v1/v2/v3, but got %s", v3.prefix)
	}
}
