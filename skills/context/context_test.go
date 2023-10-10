package context_test

import (
	"context"
	"testing"
	"time"

	mycontext "gitee.com/go-course/go12/skills/context"
)

func TestCurl(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := mycontext.Curl(ctx, "http://httpbin.org/delay/5")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPayment(t *testing.T) {
	ctx := context.Background()
	err := mycontext.Payment(ctx, "xxx")
	if err != nil {
		t.Fatal(err)
	}
}
