package analyzer

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		word string
	}

	type want struct {
		isValid bool
	}

	type test struct {
		name string
		args args
		want want
	}

	tests := []test{
		{
			name: "valid",
			args: args{
				word: "hello",
			},
			want: want{
				isValid: true,
			},
		},
		{
			name: "empty",
			args: args{
				word: "",
			},
			want: want{
				isValid: false,
			},
		},
		{
			name: "wrong_size",
			args: args{
				word: "hi",
			},
			want: want{
				isValid: false,
			},
		},
		{
			name: "with_numbers",
			args: args{
				word: "10x",
			},
			want: want{
				isValid: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.word); got != tt.want.isValid {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
