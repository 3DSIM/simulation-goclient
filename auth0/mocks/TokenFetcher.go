package mocks

import "github.com/stretchr/testify/mock"

type TokenFetcher struct {
	mock.Mock
}

// Token provides a mock function with given fields: audience
func (_m *TokenFetcher) Token(audience string) (string, error) {
	ret := _m.Called(audience)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(audience)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(audience)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
