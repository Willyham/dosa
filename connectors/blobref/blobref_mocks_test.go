// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/uber-go/dosa/connectors/blobref (interfaces: BlobStore)

package blobref_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	dosa "github.com/uber-go/dosa"
	io "io"
	url "net/url"
)

// MockBlobStore is a mock of BlobStore interface
type MockBlobStore struct {
	ctrl     *gomock.Controller
	recorder *_MockBlobStoreRecorder
}

// Recorder for MockBlobStore (not exported)
type _MockBlobStoreRecorder struct {
	mock *MockBlobStore
}

// NewMockBlobStore creates a new mock
func NewMockBlobStore(ctrl *gomock.Controller) *MockBlobStore {
	mock := &MockBlobStore{ctrl: ctrl}
	mock.recorder = &_MockBlobStoreRecorder{mock}
	return mock
}

// EXPECT adds an expectation
func (_m *MockBlobStore) EXPECT() *_MockBlobStoreRecorder {
	return _m.recorder
}

// Read is a mock implementation of MockBlobStore.Read
func (_m *MockBlobStore) Read(_param0 context.Context, _param1 *url.URL) (io.Reader, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0, _param1)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBlobStoreRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1)
}

// Remove is a mock implementation of MockBlobStore.Remove
func (_m *MockBlobStore) Remove(_param0 context.Context, _param1 *dosa.EntityInfo, _param2 *dosa.ColumnDefinition, _param3 map[string]dosa.FieldValue) error {
	ret := _m.ctrl.Call(_m, "Remove", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockBlobStoreRecorder) Remove(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0, arg1, arg2, arg3)
}

// Shutdown is a mock implementation of MockBlobStore.Shutdown
func (_m *MockBlobStore) Shutdown() error {
	ret := _m.ctrl.Call(_m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockBlobStoreRecorder) Shutdown() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shutdown")
}

// Write is a mock implementation of MockBlobStore.Write
func (_m *MockBlobStore) Write(_param0 context.Context, _param1 *dosa.EntityInfo, _param2 *dosa.ColumnDefinition, _param3 map[string]dosa.FieldValue, _param4 io.Reader) (*url.URL, error) {
	ret := _m.ctrl.Call(_m, "Write", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBlobStoreRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}