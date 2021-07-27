package models

import (
	"testing"
)

func TestServiceToString(t *testing.T) {
	service := Service{Id: 10, UserId: 1, Name: "Borgem", Link: "http://test.test"}
	expected := "<Service id=10, UserId=1, Name=Borgem, Link=http://test.test>"
	got := service.String()

	if expected != got {
		t.Errorf("not equal")
	}
}
