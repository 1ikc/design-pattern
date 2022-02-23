package visitor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompressor_Visit(t *testing.T) {
	cases := []struct{
		name    string
		path    string
		wantErr string
	}{
		{
			name: "pdf",
			path: "./xx.pdf",
		},
		{
			name: "ppt",
			path: "./xx.ppt",
		},
		{
			name:    "404",
			path:    "./xx.xx",
			wantErr: "not found file type",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f, err := NewResourceFile(c.path)
			if c.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), c.wantErr)
				return
			}

			require.NoError(t, err)
			compressor := &Compressor{}
			assert.Nil(t, f.Accpet(compressor))
		})
	}
}