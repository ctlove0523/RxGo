package rxgo

import (
	"context"
	"testing"
)

func Test_Single_Filter_True(t *testing.T) {
	os := FromItem(FromValue(1)).Filter(context.Background(), func(i interface{}) bool {
		return i == 1
	})
	Assert(context.Background(), t, os, HasItem(1), HasNotRaisedError())
}

func Test_Single_Filter_False(t *testing.T) {
	os := FromItem(FromValue(1)).Filter(context.Background(), func(i interface{}) bool {
		return i == 0
	})
	Assert(context.Background(), t, os, HasNoItem(), HasNotRaisedError())
}

func Test_Single_Map(t *testing.T) {
	single := FromItem(FromValue(1)).Map(context.Background(), func(i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	})
	Assert(context.Background(), t, single, HasItem(2), HasNotRaisedError())
}
