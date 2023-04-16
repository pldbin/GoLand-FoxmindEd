package __3_worktime_test

import (
	"foxminded/1.3_worktime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkTime1(t *testing.T) {
	tests := []struct {
		name             string
		testValue        string
		expectTime1      string
		expectTime2      string
		expectErrMessage string
	}{
		{
			name:             "unsuccessful, can`t parse time with layout",
			testValue:        "03:04:05",
			expectErrMessage: "parsing time \"03:04:05\" as \"03:04:05PM\": cannot parse \"\" as \"PM\"",
		},
		{
			name:             "expect unsuccessful parsing of 'Feb 4, 2014 at 6:05pm (PST)'",
			testValue:        "Feb 4, 2014 at 6:05pm (PST)",
			expectErrMessage: "parsing time \"Feb 4, 2014 at 6:05pm (PST)\" as \"03:04:05PM\": cannot parse \"Feb 4, 2014 at 6:05pm (PST)\" as \"03\"",
		},
		{
			name:             "hour out of range",
			testValue:        "2014-Feb-04",
			expectErrMessage: "parsing time \"2014-Feb-04\": hour out of range",
		},
		{
			name:        "for test case actual ",
			testValue:   "04:04:05PM",
			expectTime1: "04:04:05PM",
			expectTime2: "16:04:05",
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel() //так ... паралельно
			gotTime1, gotTime2, err := __3_worktime.WorkTime(tt.testValue) //WorkTime(tt.testValue)
			if err != nil {
				require.Equal(t, tt.expectErrMessage, err.Error())
			} else {
				require.Equal(t, tt.expectTime1, gotTime1)
				require.Equal(t, tt.expectTime2, gotTime2)
				require.NoError(t, err)
			}
		})
	}
}
