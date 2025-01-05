package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStatistic_BuildStatisticByFile(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name    string
		want    *Statistic
		args    args
		wantErr string
	}{
		{
			name: "can build statistic by file",
			args: args{
				inputFile: "../input.log",
			},
			want: &Statistic{
				CountRows:  66,
				CountError: 2,
			},
			wantErr: "",
		},
		{
			name: "fail when file not found",
			args: args{
				inputFile: "../notExists.log",
			},
			want: &Statistic{
				CountRows:  0,
				CountError: 0,
			},
			wantErr: "no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Statistic{}
			if err := got.BuildStatisticByFile(tt.args.inputFile); err != nil {
				assert.Error(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.want.CountRows, got.CountRows)
			assert.Equal(t, tt.want.CountError, got.CountError)
		})
	}
}

func TestStatistic_String(t *testing.T) {
	s := Statistic{
		CountRows:  66,
		StartTime:  time.Date(2024, 11, 5, 9, 31, 31, 44707000, time.UTC),
		EndTime:    time.Date(2024, 11, 5, 9, 32, 34, 594396000, time.UTC),
		CountError: 2,
		verbose:    "",
		rows:       nil,
	}
	want := "Result analyze:\n" +
		"66 rows analyzed by period [2024-11-05 09:31:31.044707 +0000 UTC - 2024-11-05 09:32:34.594396 +0000 UTC]\n" +
		"Found 2 errors"

	assert.Equal(t, want, s.String())
}
