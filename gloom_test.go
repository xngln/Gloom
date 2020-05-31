package gloom

import "testing"

func TestNew(t *testing.T) {
	m := uint64(100)
	k := uint64(6)
	bf := New(m, k)

	if bf.m != m {
		t.Errorf("Got: %v, expected: %v", bf.m, m)
	}

	if bf.k != k {
		t.Errorf("Got: %v, expected: %v", bf.k, k)
	}
}

func TestAdd(t *testing.T) {
	m := uint64(100)
	k := uint64(6)
	bf := New(m, k)

	bf.AddString("golang")

	if !bf.ContainsString("golang") {
		t.Errorf("Got false, expected true")
	}

	if bf.ContainsString("python") {
		t.Errorf("Got true, expected false")
	}
}
