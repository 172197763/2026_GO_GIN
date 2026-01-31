package cmd

import (
	"errors"
	"testing"
)

// 性能对比
func BenchmarkErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ErrorsNew("test error")
	}
}

// 即使不用格式化也更慢
func BenchmarkFmtErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ErrorsFmt("fixed error")
	}
}

// 单元测试
func TestErrorsNew(t *testing.T) {
	result := ErrorsNew("test error")
	expected := errors.New("test error")
	if result.Error() != expected.Error() {
		t.Errorf("ErrorsNew(\"test error\") = %d; want %d", result, expected)
	}
}
