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

// 输出结果
// ns/op:每次操作平均耗时（纳秒） B/op:每次操作分配的堆内存字节数  allocs/op:每次操作发生的堆内存分配次数
// BenchmarkErrorsNew-16    	1000000000	         0.2190 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFmtErrorf-16    	11271648	       106.6 ns/op	      48 B/op	       2 allocs/op

// 单元测试
func TestErrorsNew(t *testing.T) {
	result := ErrorsNew("test error")
	expected := errors.New("test error")
	if result.Error() != expected.Error() {
		t.Errorf("ErrorsNew(\"test error\") = %d; want %d", result, expected)
	}
}
