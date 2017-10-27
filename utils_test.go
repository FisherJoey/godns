package godns

import (
	"testing"
)

func testGetCurrentIP(t *testing.T) {
	ip, _ := GetCurrentIP("http://members.3322.org/dyndns/getip")

	if ip == "" {
		t.Log("IP is empty...")
	} else {
		t.Log("IP is:" + ip)
	}
}