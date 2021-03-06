package cmac

import (
	"bytes"
	"testing"
)

func TestGenSubkeys(t *testing.T) {
	for i, tv := range nistvectors {
		c, err := tv.cipher(tv.key)
		if err != nil {
			t.Fatal(err)
		}
		k1, k2 := gensubkeys(c)
		if !bytes.Equal(k1, tv.k1) {
			t.Errorf("tv[%d]: expected: %x got %x\n", i, tv.k1, k1)
		}
		if !bytes.Equal(k2, tv.k2) {
			t.Errorf("tv[%d]: expected: %x got %x\n", i, tv.k2, k2)
		}
	}
}

func TestNew(t *testing.T) {
	for i := 0; i < 3; i++ {
		tv := nistvectors[i]

		m, err := New(tv.key)
		if err != nil {
			t.Fatalf("tv[%d]: New() err: %s\n", i, err)
		}
		for j, tc := range tv.cases {
			m.Write(tc.msg)
			mac := m.Sum(nil)
			if !bytes.Equal(mac, tc.mac) {
				t.Errorf("tv[%d,%d]: expected: %x got %x\n", i, j, tc.mac, mac)
			}
			m.Reset()
		}
	}
}

func TestNewWithCipher(t *testing.T) {
	for i, tv := range nistvectors {
		c, err := tv.cipher(tv.key)
		if err != nil {
			t.Fatal(err)
		}
		m, err := NewWithCipher(c)
		if err != nil {
			t.Fatalf("tv[%d]: NewWithCipher() err: %s\n", i, err)
		}
		for j, tc := range tv.cases {
			m.Write(tc.msg)
			mac := m.Sum(nil)
			if !bytes.Equal(mac, tc.mac) {
				t.Errorf("tv[%d,%d]: expected: %x got %x\n", i, j, tc.mac, mac)
			}
			m.Reset()
		}
	}
}
