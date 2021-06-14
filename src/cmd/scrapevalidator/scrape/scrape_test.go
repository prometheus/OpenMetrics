package scrape

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
)

func TestValidate(t *testing.T) {
	tcs := []testCase{
		{
			name: "bad_counter_decreasing",
			exports: []string{
				`# TYPE a counter
# HELP a help
a_total 2
# EOF`,
				`# TYPE a counter
# HELP a help
a_total 1
# EOF`,
			},
			expectedErr: errMustNotCounterValueDecrease,
		},
		{
			name: "good_counter_increasing",
			exports: []string{
				`# TYPE a counter
# HELP a help
a_total 1
# EOF`,
				`# TYPE a counter
# HELP a help
a_total 2
# EOF`,
			},
		},
		{
			name: "bad_metric_disappearing",
			exports: []string{
				`# TYPE a counter
# HELP a help
a_total 1
# EOF`,
				`# TYPE b counter
# HELP b help
b_total 2
# EOF`,
			},
			expectedErr: errMustNotSeriesDisappear,
		},
		{
			name: "good_not_duplicate_labels",
			exports: []string{
				`# TYPE a1 counter
# HELP a1 help
a1_total{bar="baz1"} 1
# TYPE a2 counter
# HELP a2 help
a2_total{bar="baz2"} 1
# EOF`,
			},
		},
		{
			name: "bad_duplicate_labels",
			exports: []string{
				`# TYPE a1 counter
# HELP a1 help
a1_total{bar="baz"} 1
# TYPE a2 counter
# HELP a2 help
a2_total{bar="baz"} 1
# EOF`,
			},
			expectedErr: errShouldNotDuplicateLabel,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			run(t, tc)
		})
	}
}

type testCase struct {
	name        string
	exports     []string
	expectedErr error
}

func run(t *testing.T, tc testCase) {
	s := newScraperLoop()
	var mErr error
	for _, export := range tc.exports {
		err := s.parseAndValidate([]byte(export), time.Now())
		mErr = multierr.Append(mErr, err)
	}
	if tc.expectedErr == nil {
		require.NoError(t, mErr)
		return
	}
	require.Equal(t, mErr.Error(), tc.expectedErr.Error())
}
