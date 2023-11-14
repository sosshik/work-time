package main

import (
	"fmt"
	"os"
	"testing"
)

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

func Test_main1(t *testing.T) {
	newArgs := []string{os.Args[0], "-time=03:04:05PM", "asfeasf"}
	tests := []struct {
		name string
	}{
		{"panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()
			os.Args = newArgs
			main()
			t.Errorf("The code did not panic")
		})
	}
}

//This test is just for learning purposes. As we don't have too much to test in this project.
//We are tring to test if main function is properly hadling different amount of os.Args.
//Test is runnig main with custom os.Args and in case of panic recovers it and test passes.
//If no panic happened test fails by t.Errorf.

func Test_main2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()
			os.Args = []string{""}
			main()
			t.Errorf("The code did not panic")
		})
	}
}

//This test is just for learning purposes. As we don't have too much to test in this project.
//We are tring to test if main function is properly hadling different amount of os.Args.
//Test is runnig main with custom os.Args and in case of panic recovers it and test passes.
//If no panic happened test fails by t.Errorf.

func BenchmarkFormatTime(b *testing.B) {

	time := "2:01:00AM"

	for i := 0; i < b.N; i++ {
		FormatTime(time)
	}

}
