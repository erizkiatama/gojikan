package gojikan

import "net/http"

// MockFunc is a custom type that allows setting the func that our Mock Do func will run instead
type MockFunc func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockFunc
}

// Do is a function that overrides what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}
