package glog

import "testing"

func Test_init(t *testing.T) {
	if Config.LogService.Appid != "glog" {
		t.Error()
	}
}
