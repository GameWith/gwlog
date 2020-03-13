package formatter

import "testing"

func TestFieldMap_get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		f    FieldMap
		args args
		want string
	}{
		{
			name: "none",
			f:    FieldMap{},
			args: args{FieldKeyLevel},
			want: FieldKeyLevel,
		},
		{
			name: "set alias @level",
			f: FieldMap{
				FieldKeyLevel: "@level",
			},
			args: args{FieldKeyLevel},
			want: "@level",
		},
		{
			name: "set alias hoge",
			f: FieldMap{
				"hoge": "fuga",
			},
			args: args{"hoge"},
			want: "fuga",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.get(tt.args.key); got != tt.want {
				t.Errorf("FieldMap.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFields_conflict(t *testing.T) {
	type args struct {
		key string
		fm  FieldMap
	}
	tests := []struct {
		name string
		f    Fields
		args args
		want bool
	}{
		{
			name: "none",
			f:    Fields{},
			args: args{
				key: "w",
				fm:  FieldMap{},
			},
			want: false,
		},
		{
			name: "same field name",
			f: Fields{
				"hoge": "fuga",
			},
			args: args{
				key: "hoge",
				fm: FieldMap{
					"hoge": "hoge",
				},
			},
			want: true,
		},
		{
			name: "different field name",
			f: Fields{
				"hoge": "fuga",
			},
			args: args{
				key: "hoge",
				fm: FieldMap{
					"hoge": "@hoge",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.conflict(tt.args.key, tt.args.fm); got != tt.want {
				t.Errorf("Fields.conflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFields_has(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		f    Fields
		args args
		want bool
	}{
		{
			name: "none",
			f:    Fields{},
			args: args{"hoge"},
			want: false,
		},
		{
			name: "exists",
			f:    Fields{"hoge": 0},
			args: args{"hoge"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.has(tt.args.key); got != tt.want {
				t.Errorf("Fields.has() = %v, want %v", got, tt.want)
			}
		})
	}
}
