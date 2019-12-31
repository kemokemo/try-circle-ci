package main

import (
	"bytes"
	"flag"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		out    string
		outerr string
	}{
		{"normal", args{[]string{"alice"}}, exitCodeOK, "Hello alice!\n", ""},
		{"empty args", args{[]string{}}, exitCodeFailed, "", "please specify args\n"},
		{"empty string", args{[]string{""}}, exitCodeFailed, "", "failed to execute the greeter: name is empty\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bout := new(bytes.Buffer)
			out = bout
			bouterr := new(bytes.Buffer)
			outerr = bouterr

			if got := run(tt.args.args); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}

			stdout := bout.String()
			if stdout != tt.out {
				t.Errorf("stdout = %v, want %v", stdout, tt.out)
			}

			stderr := bouterr.String()
			if stderr != tt.outerr {
				t.Errorf("stderr = %v, want %v", stderr, tt.outerr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		exitCode int
		out      string
		outerr   string
	}{
		{"sigle name", []string{"John"}, 0, "Hello John!\n", ""},
		{"multiple names", []string{"Mike", "Risa"}, 0, "Hello Mike!\nHello Risa!\n", ""},
		{"empty args", []string{}, 1, "", "please specify args\n"},
		{"empty string", []string{""}, 1, "", "failed to execute the greeter: name is empty\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("testing main", flag.ContinueOnError)
			flagSet.Parse(tt.args)
			cmdArgs = flagSet.Args()

			bout := new(bytes.Buffer)
			out = bout
			bouterr := new(bytes.Buffer)
			outerr = bouterr

			var exitCode int
			exit = func(n int) {
				exitCode = n
			}

			main()

			if exitCode != tt.exitCode {
				t.Errorf("exit code = %v, want %v", exitCode, tt.exitCode)
			}

			stdout := bout.String()
			if stdout != tt.out {
				t.Errorf("stdout = %v, want %v", stdout, tt.out)
			}

			stderr := bouterr.String()
			if stderr != tt.outerr {
				t.Errorf("stderr = %v, want %v", stderr, tt.outerr)
			}
		})
	}
}
