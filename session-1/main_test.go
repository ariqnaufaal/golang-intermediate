package main

import "testing"

func TesHello2(tes *testing.T) {
	res := GetHello2()
	if res != "Hello World 2" {
		tes.Fail()
	}
}

/*
func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
*/
