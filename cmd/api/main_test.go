package main

import "testing"

func TestHello(t *testing.T) {
	want := "salemou aalaykom"

	got := hello()

	if want != got {
		t.Fatalf("t7erbed, wanted %s, amma got %s", want, got)
	}
}
