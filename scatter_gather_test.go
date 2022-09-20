package scattergather

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScatterGatherWithInputParams(t *testing.T) {
	testcases := []struct {
		name      string
		input     []string
		batchSize int
		fn        func([]string) ([]string, error)
		want      []string
		err       error
	}{
		{
			name:      "test input size less than batch size",
			input:     []string{"bird", "cat", "dog", "fish"},
			batchSize: 5,
			fn: func(input []string) ([]string, error) {
				var rs []string
				for _, str := range input {
					rs = append(rs, "my "+str)
				}

				return rs, nil
			},
			want: []string{
				"my bird",
				"my cat",
				"my dog",
				"my fish",
			},
			err: nil,
		},
		{
			name:      "test input size equal batch size",
			input:     []string{"bird", "cat", "dog", "fish"},
			batchSize: 4,
			fn: func(input []string) ([]string, error) {
				var rs []string
				for _, str := range input {
					rs = append(rs, "my "+str)
				}

				return rs, nil
			},
			want: []string{
				"my bird",
				"my cat",
				"my dog",
				"my fish",
			},
			err: nil,
		},
		{
			name:      "test input size great than batch size",
			input:     []string{"bird", "cat", "dog", "fish"},
			batchSize: 2,
			fn: func(input []string) ([]string, error) {
				var rs []string
				for _, str := range input {
					rs = append(rs, "my "+str)
				}

				return rs, nil
			},
			want: []string{
				"my bird",
				"my cat",
				"my dog",
				"my fish",
			},
			err: nil,
		},
		{
			name:      "test invalid batch size",
			input:     []string{"bird", "cat", "dog", "fish"},
			batchSize: -1,
			fn:        nil,
			want:      []string{},
			err:       ErrInvalidBatchSize,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ScattergatherWithInputParams(tc.input, tc.batchSize, tc.fn)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
				return
			}

			sort.Strings(tc.want)
			sort.Strings(got)
			assert.Equal(t, tc.want, got)
		})
	}
}
