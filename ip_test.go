package ins

import "testing"

// test ip to 4
func TestGetTo4(t *testing.T) {

	ip, err := GetTo4()
	if err != nil {
		t.Log(err)
	}

	t.Log(ip.String())

}