// +build linux

package netlink

import (
	"testing"
	"time"
)

func TestSetGetSocketTimeout(t *testing.T) {
	timeout := 10 * time.Second
	if err := SetSocketTimeout(10 * time.Second); err != nil {
		t.Fatalf("Set socket timeout for default handle failed: %v", err)
	}

	if val, err := GetSocketTimeout(); err != nil {
		t.Fatalf("Get socket timeout for default handle failed: %v", err)
	} else {
		if val != timeout {
			t.Fatalf("Unexpcted socket timeout value: got=%v, expect=%v", val, timeout)
		}
	}
}
