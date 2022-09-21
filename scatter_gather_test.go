package scattergather

import (
	"errors"
	"reflect"
	"sort"
	"testing"
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
			want:      nil,
			err:       ErrInvalidBatchSize,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ScattergatherWithInputParams(tc.input, tc.batchSize, tc.fn)
			if tc.err != err {
				t.Fatalf("tc: %s, got: %+v, want: %+v", tc.name, err, tc.err)
			}

				sort.Strings(tc.want)
				sort.Strings(got)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("tc: %s, got: %v, want %v, err: %v %v %v", tc.name, got, tc.want, tc.err, err, tc.err == err)
				}
		})
	}
}

func TestScatterGatherWithInput_ReturnError(t *testing.T) {
	want := errors.New("random failure")

	input := []string{"bird", "cat", "dog", "fish"}
	_, err := ScattergatherWithInputParams(input, 2, func(k []string) ([]string, error) {
		return nil, want
	})

	if err != want {
		t.Fatalf("got: %v, want: %v", err, want)
	}
	// assert.EqualError(t, err, "failed to handle")
}
