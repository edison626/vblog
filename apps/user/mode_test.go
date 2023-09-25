package user_test

import (
	"crypto/md5"
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
// $2a$10$dOyuOkFXEgRkSbbPPCBEduqEYuRn.
// $2a$10$vIft2TQGT5WBxjSgAAuKye.
// $2a$10$88WM70UQgEUK8di63ZCBUOrgR6Q0fJbwTkpd.
func TestBcrypto5Hash(t *testing.T) {
	b, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	//打印
	t.Log(string(b))

	//打印16进制哈希值
	fmt.Printf("%x", b)

	//测试对比值 - salt 解密后是否一样
	err := bcrypt.CompareHashAndPassword(b, []byte("123456"))
	if err != nil {
		t.Log(err)
	}
}
