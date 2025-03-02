package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTag(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Tags
		wantErr error
	}{
		{
			name: "Normal/single tag",
			args: `name:"test"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "test",
					Options: nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/single tag with empty name",
			args: `name:",test"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "",
					Options: Options{Option{Name: "test"}},
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/single tag with backslash comma",
			args: `na\,me:"test"`,
			want: Tags{
				{
					Key:     "na,me",
					Name:    "test",
					Options: nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/single tag with double options",
			args: `name:"test,opt1,opt2,"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "test",
					Options: Options{Option{Name: "opt1"}, Option{Name: "opt2"}},
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/single tag with double options (use option with value)",
			args: `name:"test,opt1=abc,opt2=abc\,"`,
			want: Tags{
				{
					Key:  "name",
					Name: "test",
					Options: Options{
						Option{Name: "opt1", Value: "abc"},
						Option{Name: "opt2", Value: "abc,"},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/single tag with double options (irregular case1)",
			args: `name:"test,opt1:ab,opt2,"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "test",
					Options: Options{Option{Name: "opt1:ab"}, Option{Name: "opt2"}},
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/single tag with double options (irregular case2)",
			args: `name:"test,opt1$%ab,opt2*s,"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "test",
					Options: Options{Option{Name: "opt1$%ab"}, Option{Name: "opt2*s"}},
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/double tag",
			args: `name:"test" json:"test,omitempty"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "test",
					Options: nil,
				},
				{
					Key:     "json",
					Name:    "test",
					Options: Options{Option{Name: "omitempty"}},
				},
			},
			wantErr: nil,
		},
		{
			name: "Normal/double tag with double options",
			args: `name:"test,omitempty,debug" json:"test,omitempty"`,
			want: Tags{
				{
					Key:     "name",
					Name:    "test",
					Options: Options{Option{Name: "omitempty"}, Option{Name: "debug"}},
				},
				{
					Key:     "json",
					Name:    "test",
					Options: Options{Option{Name: "omitempty"}},
				},
			},
			wantErr: nil,
		},
		{
			name:    "NG/tag format error case1",
			args:    `name:"test`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/tag format error case2",
			args:    `name:"`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/tag format error case3",
			args:    `name"`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/tag format error case4",
			args:    `name`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/tag format error case5",
			args:    `name:`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/tag format error case6",
			args:    `name:"":`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/tag format error case7",
			args:    `n,ame:""`,
			want:    nil,
			wantErr: ErrTagSyntax,
		},
		{
			name:    "NG/key format error",
			args:    `:"test"`,
			want:    nil,
			wantErr: ErrTagKeySyntax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTag(tt.args)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
