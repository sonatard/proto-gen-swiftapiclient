package str

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ToAbbreviation(t *testing.T) {
	type args struct {
		s string
	}
	type rets struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want rets
	}{
		{
			name: "Userid -> Userid",
			args: args{
				s: "Userid",
			},
			want: rets{
				s: "Userid",
			},
		},
		{
			name: "UserId1 -> UserId1",
			args: args{
				s: "UserId1",
			},
			want: rets{
				s: "UserId1",
			},
		},
		{
			name: "UserId -> ID",
			args: args{
				s: "UserId",
			},
			want: rets{
				s: "UserID",
			},
		},
		{
			name: "Url -> URL",
			args: args{
				s: "Url",
			},
			want: rets{
				s: "URL",
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q", tt.name), func(t *testing.T) {
			got := ToAbbreviation(tt.args.s)
			if !cmp.Equal(tt.want.s, got, nil) {
				t.Errorf("want:%v, got:%v", tt.want.s, got)
			}
		})
	}
}
