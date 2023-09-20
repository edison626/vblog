package user_test

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// https://www.somd5.com/
func TestMd5Hash(t *testing.T) {
	h := md5.New()
	_, err := h.Write([]byte("123456"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%x", h.Sum(nil))
}

// bcrypt 散列算法 - https://gitee.com/infraboard/go-course/blob/master/day09/go-hash.md
// 加了盐后可以看到前面的部分是相同的
// JDJhJDEwJC43T29qaGFnR0VDMmUucDhIcUJ5bU9vRUJlUU5HZlNYNUxVdERSNjBVMmZIY29oTmJEcVgy
// JDJhJDEwJDI4VVliNG9tS3VIQ1diUURTWm5IWS5oblhhUmNNUmhNSHczcE5tWkZiYTN1dGJrRGlaSDF5
// JDJhJDEwJGY4MVpxT0xzbm9FZjkvWUx3WHBZM094UU9sSkNDUi82YWk5S2NmalZ2RUxMdUhDemN2T0oy
func TestBcrypto5Hash(t *testing.T) {
	b, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	//打印16进制哈希值
	fmt.Printf("%x", b)

	//打印base64 比较常用
	t.Log(base64.StdEncoding.EncodeToString(b))

	//测试对比值 - salt 解密后是否一样
	err := bcrypt.CompareHashAndPassword(b, []byte("123456"))
	if err != nil {
		t.Log(err)
	}
}
