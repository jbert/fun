package fun

import (
	"io"
	"testing"
)

func TestPair(t *testing.T) {
	p := Pair[int, error]{5, io.EOF}

	if p.First != 5 {
		t.Fatalf("Can't get first elt")
	}
	if p.Second != io.EOF {
		t.Fatalf("Can't get second elt")
	}

	a, b := p.Unpack()
	if a != 5 {
		t.Fatalf("Can't unpack first elt")
	}
	if b != io.EOF {
		t.Fatalf("Can't unpack second elt")
	}
}
