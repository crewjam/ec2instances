package ec2instances

import (
	"testing"
)

func Test(t *testing.T) {
	instances := GetInstance("t2.micro")
	if instances.Memory != 1.0 {
		t.Errorf("instance[t2.micro].Memory should be 1.0 not %v", instances.Memory)
	}
}
