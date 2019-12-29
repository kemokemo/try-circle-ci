package main

import "testing"

func Test_greeter(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"bob", args{"bob"}, "Hello bob!", false},
		{"empty", args{""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := greeter(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("greeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("greeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
