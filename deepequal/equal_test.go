package equal

import (
	"reflect"
	"testing"
)

// go test -test.bench=".*" -count=3

func BenchmarkIsEqual(b *testing.B) {
	m1 := &member{}
	m2 := &member{}
	for i := 0; i < b.N; i++ { //use b.N for looping
		m1.IsEqual(m2)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	m1 := &member{}
	m2 := &member{}
	for i := 0; i < b.N; i++ { //use b.N for looping
		reflect.DeepEqual(m1, m2)
	}
}
