// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: PluginHelper)

package pipeline

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of PluginHelper interface
type MockPluginHelper struct {
	ctrl     *gomock.Controller
	recorder *_MockPluginHelperRecorder
}

// Recorder for MockPluginHelper (not exported)
type _MockPluginHelperRecorder struct {
	mock *MockPluginHelper
}

func NewMockPluginHelper(ctrl *gomock.Controller) *MockPluginHelper {
	mock := &MockPluginHelper{ctrl: ctrl}
	mock.recorder = &_MockPluginHelperRecorder{mock}
	return mock
}

func (_m *MockPluginHelper) EXPECT() *_MockPluginHelperRecorder {
	return _m.recorder
}

func (_m *MockPluginHelper) Output(_param0 string) (OutputRunner, bool) {
	ret := _m.ctrl.Call(_m, "Output", _param0)
	ret0, _ := ret[0].(OutputRunner)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockPluginHelperRecorder) Output(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Output", arg0)
}

func (_m *MockPluginHelper) PackSupply() chan *PipelinePack {
	ret := _m.ctrl.Call(_m, "PackSupply")
	ret0, _ := ret[0].(chan *PipelinePack)
	return ret0
}

func (_mr *_MockPluginHelperRecorder) PackSupply() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PackSupply")
}

func (_m *MockPluginHelper) Router() *MessageRouter {
	ret := _m.ctrl.Call(_m, "Router")
	ret0, _ := ret[0].(*MessageRouter)
	return ret0
}

func (_mr *_MockPluginHelperRecorder) Router() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Router")
}
