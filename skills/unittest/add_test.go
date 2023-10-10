package unittest_test

import (
	"testing"

	"gitee.com/go-course/go12/skills/unittest"
	"github.com/go-playground/assert/v2"
)

func TestSum(t *testing.T) {
	// print
	// 如果打印不出来值, 需要调整ide设置 test Flag的设置
	// "go.testFlags": [
	// 你的单元测试和代码都修改, 单元测试会被缓存
	//     "-count=1",
	// 打印单元测试过程中的详细日志 fmt.xxx
	//     "-v"
	// ],

	// if unittest.Sum(1, 2) != 3 {
	// 	t.Fatal("测试失败")
	// }

	// 使用断言库
	assert.Equal(t, unittest.Sum(1, 2), 3)
}
