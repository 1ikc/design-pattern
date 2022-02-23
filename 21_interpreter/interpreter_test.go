package interpreter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAndExpression_Interpret(t *testing.T) {
	stats := map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	cases := []struct{
		name  string
		stats map[string]float64
		rule  string
		want  bool
	}{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a < 2 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, err := NewAlertRule(c.rule)
			require.Nil(t, err)
			assert.Equal(t, c.want, r.Interpret(c.stats))
		})
	}
}