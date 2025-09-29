package main

import (
	"time"
)

// bubble内の初期時刻を指定した時刻に設定する。
// synctest.Testの冒頭で呼び出すこと。
// 2000-01-01 00:00:00 UTC以降の時刻を指定すること。
func SetNow(t T, now time.Time) {
	t.Helper()
	bubbleInitialTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if !time.Now().Equal(bubbleInitialTime) {
		t.Fatalf("SetNow must be called at the beginning of synctest.Test")
	}
	d := now.Sub(time.Now())
	if d < 0 {
		t.Fatalf("SetNow must be called with time after %v", bubbleInitialTime)
	}
	time.Sleep(d)
}

type T interface {
	Helper()
	Fatalf(format string, args ...any)
}
