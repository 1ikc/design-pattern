package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputText_Snapshot(t *testing.T) {
	in := &InputText{}
	snapshots := []*Snapshot{}

	cases := []struct{
		input string
		want  string
	}{
		{
			input: ":list",
			want:  "",
		},
		{
			input: "hello",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
		{
			input: "world",
			want:  "",
		},
		{
			input: ":list",
			want:  "helloworld",
		},
		{
			input: ":undo",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			switch c.input {
			case ":list":
				assert.Equal(t, c.want, in.GetText())
			case ":undo":
				in.Restore(snapshots[len(snapshots)-1])
				snapshots = snapshots[:len(snapshots)-1]
			default:
				snapshots = append(snapshots, in.Snapshot())
				in.Append(c.input)
			}
		})
	}
}