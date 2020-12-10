package timer

import (
    "time"
    _ "unsafe"
)

// NanoTime returns the current time in nanoseconds from a monotonic clock.
//go:linkname nanoTime runtime.nanotime
func nanoTime() int64

func Now() time.Duration { return time.Duration(nanoTime()) }
