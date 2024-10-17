package foptions_test

import (
	"errors"
	"testing"

	"github.com/ciricc/go-foptions"
	"github.com/stretchr/testify/assert"
)

type CustomOption foptions.Opt[MySettings]

type MySettings struct {
	Limit int
}

var (
	errInvalidLimit = errors.New("invalid limit option")
)

func WithLimit(limit int) foptions.Opt[MySettings] {
	return func(settings *MySettings) error {
		if limit < 1 {
			return errInvalidLimit
		}

		settings.Limit = limit
		return nil
	}
}

func TestUse(t *testing.T) {
	testCases := []struct {
		title         string
		opts          []foptions.Opt[MySettings]
		expectedError error
		defaultValue  *MySettings
		expectedValue *MySettings
	}{{
		title: "Successfully update limit setting to the 1",
		opts: []foptions.Opt[MySettings]{
			WithLimit(1),
		},
		defaultValue: &MySettings{
			Limit: 0,
		},
		expectedValue: &MySettings{
			Limit: 1,
		},
	},
		{
			title: "Successfully update limit setting to the 1, but return an error due last option",
			opts: []foptions.Opt[MySettings]{
				WithLimit(1),
				WithLimit(0),
			},
			expectedError: errInvalidLimit,
			defaultValue: &MySettings{
				Limit: 0,
			},
			expectedValue: &MySettings{
				Limit: 1,
			},
		},
		{
			title: "Must return an error",
			opts: []foptions.Opt[MySettings]{
				WithLimit(0),
			},
			expectedError: errInvalidLimit,
			defaultValue: &MySettings{
				Limit: 0,
			},
			expectedValue: &MySettings{
				Limit: 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			newSettings, err := foptions.Use(tc.defaultValue, tc.opts...)

			assert.ErrorIs(t, tc.expectedError, err, "expected error is")
			assert.Equal(t, tc.expectedValue, tc.defaultValue)
			assert.Equal(t, tc.expectedValue, newSettings)
		})

	}
}
