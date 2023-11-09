package main

import "testing"

func TestFormatTime(t *testing.T) {
	t.Parallel()

	type args struct {
		timeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"12:01:00PM", args{"12:01:00PM"}, "12:01:00", false},
		{"12:01:00AM", args{"12:01:00AM"}, "00:01:00", false},
		{"2:01:00PM", args{"02:01:00PM"}, "14:01:00", false},
		{"2:01:00AM", args{"02:01:00AM"}, "02:01:00", false},
		{"lowercase", args{"04:52:00pm"}, "16:52:00", false},
		{"only 1 letter lowercase", args{"09:25:00Am"}, "09:25:00", false},
		{"wrong format", args{"5:55PM"}, "", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := FormatTime(tt.args.timeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFormatTime(b *testing.B) {

	time := "2:01:00AM"

	for i := 0; i < b.N; i++ {
		FormatTime(time)
	}

}
