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

// $2a$10$dOyuOkFXEgRkSbbPPCBEduqEYuRn.
// $2a$10$vIft2TQGT5WBxjSgAAuKye.
// $2a$10$88WM70UQgEUK8di63ZCBUOrgR6Q0fJbwTkpd.
func TestBcrypto5Hash(t *testing.T) {
	b, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	t.Log(string(b))

	err := bcrypt.CompareHashAndPassword(b, []byte("123456"))
	if err != nil {
		t.Log(err)
	}
}
