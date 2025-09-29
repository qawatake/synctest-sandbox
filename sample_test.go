package main

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestSetNow(t *testing.T) {
	t.Run("bubble内で呼び出さないとエラーになる", func(t *testing.T) {
		mt := newMockT()
		SetNow(mt, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
		if !mt.Failed() {
			t.Fatal("expected SetNow to fail")
		}
	})

	t.Run("過去の時刻を指定するとエラーになる", func(t *testing.T) {
		mt := newMockT()
		SetNow(mt, time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC))
		if !mt.Failed() {
			t.Fatal("expected SetNow to fail")
		}
	})

	t.Run("bubble内ですでに時刻が進んでいるとエラーになる", func(t *testing.T) {
		mt := newMockT()
		synctest.Test(t, func(t *testing.T) {
			time.Sleep(1 * time.Second)
			SetNow(mt, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
		})
		if !mt.Failed() {
			t.Fatal("expected SetNow to fail")
		}
	})
}

type mockT struct {
	pass bool
}

func newMockT() *mockT {
	return &mockT{pass: true}
}

func (m *mockT) Helper() {}

func (m *mockT) Fatalf(format string, args ...interface{}) {
	m.pass = false
}

func (m *mockT) Failed() bool {
	return !m.pass
}
