package main

import "testing"

func TestMakeHTTPRequest(t *testing.T) {
	tests := []struct {
		testcase string
		url      string
		expect   func(*testing.T, error)
	}{
		{
			testcase: "when request is successful",
			url:      "http://google.com",
			expect: func(t *testing.T, err error) {
				if err != nil {
					t.Fail()
				}
			},
		},
		{
			testcase: "when request is not successful",
			url:      "http://zunnorayn.com",
			expect: func(t *testing.T, err error) {
				if err == nil {
					t.Fail()
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.testcase, func(t *testing.T) {
			// call the function to get return values
			_, err := MakeHTTPRequest(tt.url)

			// handle the test case's assertions with the provided func
			tt.expect(t, err)
		})
	}
}

func TestConvertBytesToMD5(t *testing.T) {
	bytes := []byte("test_bytes")
	hash := ConvertBytesToMD5(bytes)
	if len(hash) == 0 {
		t.Fail()
	}
}
