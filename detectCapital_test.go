package demo

import "testing"

func Test_detectCapitalUser(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"CHINA"}, true},
		{"normal", args{"FlaG"}, false},
		{"normal", args{"FLAg"}, false},
		{"normal", args{"jialin"}, true},
		{"normal", args{"J"}, true},
		{"normal", args{"j"}, true},
		{"normal", args{"jJ"}, false},
		{"normal", args{"jj"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUser(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
