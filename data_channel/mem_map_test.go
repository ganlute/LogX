package data_channel

import (
	"github.com/ganlute/LogX/data_channel/api"
	"testing"
)

func TestDataModule(t *testing.T) {
	var dataChannel api.DataChannel
	dataChannel = &MemMap{}

	dataChannel.InitContext()

	err := dataChannel.SetContext("123")
	if err != nil {
		t.Error(err)
	}
	context,err := dataChannel.GetContext()
	if err != nil {
		t.Error(err)
	}
	ctx := context.(string)
	if ctx != "123" {
		t.Error("ctx != 123")
	}
	err = dataChannel.UpdateContext( "234")
	if err != nil {
		t.Error(err)
	}
	context,err = dataChannel.GetContext()
	if err != nil {
		t.Error(err)
	}
	ctx = context.(string)
	if ctx != "234" {
		t.Error("ctx != 234")
	}
	err = dataChannel.DeleteContext()
	if err != nil {
		t.Error(err)
	}
	context,err = dataChannel.GetContext()
	if err != nil {
		if err.Error() == Record_Not_Found {
			// pass
		}else {
			t.Error(err)
		}
	}else {
		t.Error("err is nil")
	}
}
