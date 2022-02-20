package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourcePoolConfigBuilder_Build(t *testing.T) {
	cases := []struct{
		name    string
		builder *ResourcePoolConfigBuilder
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &ResourcePoolConfigBuilder{
				name: "",
				maxTotal: 0,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  10,
				minIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &ResourcePoolConfigBuilder{
				name: "test",
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: defaultMaxTotal,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
			wantErr: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.builder.Build()
			if !assert.Equal(t, c.wantErr, err != nil) || !assert.Equal(t, c.want, got) {
				t.Errorf("Build() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}