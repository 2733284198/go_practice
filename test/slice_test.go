package test

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		s []string
		a string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.s, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		s []string
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.s, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Uniq(tt.args.s); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Uniq() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
