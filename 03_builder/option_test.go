package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResourcePoolConfig(t *testing.T) {
	type args struct {
		name string
		opts []ResourcePoolConfigOptFunc
	}
	cases := []struct {
		name    string
		args    args
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			args: args{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name: "test",
				opts: []ResourcePoolConfigOptFunc{
					func(option *ResourcePoolConfigOption) {
						option.minIdle = 2
					},
				},
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: 10,
				maxIdle:  9,
				minIdle:  2,
			},
			wantErr: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := NewResourcePoolConfig(c.args.name, c.args.opts...)
			if !assert.Equal(t, c.wantErr, err != nil) || !assert.Equal(t, c.want, got) {
				t.Errorf("NewResourcePoolConfig() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}