package test

import (
	"sign/service"
	"testing"
)

func TestFetch(t *testing.T) {
	service.Fetch("http://180.76.172.177/18nodone")
}
