package client

import (
	"testing"
)

func TestAddParamToUrl(t *testing.T) {
	tests := []struct {
		url      string
		param    string
		value    string
		expected string
	}{
		{"http://example.com", "param", "value", "http://example.com?param=value"},
		{"http://example.com?existing=1", "param", "value", "http://example.com?existing=1&param=value"},
		{"http://example.com", "param", "", "http://example.com"},
	}

	for _, test := range tests {
		result := addParamToUrl(test.url, test.param, test.value)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
