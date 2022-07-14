package util

import "testing"

func TestEncodeMD5(t *testing.T) {
	md5 := EncodeMD5("hello")
	t.Log(md5)
}

/*=== RUN   TestEncodeMD5
    md5_test.go:7: 5d41402abc4b2a76b9719d911017c592
--- PASS: TestEncodeMD5 (0.00s)
PASS
*/
