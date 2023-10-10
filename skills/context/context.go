package context

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// curl, ctx 都是函数的第一个参数
func Curl(ctx context.Context, url string) error {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 携带上下文
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	// resp, err := client.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

type UserName struct{}
type UserRole struct{}

// 1. 用户认证 Login() tk ---> 用户的身份
// 2. 转账 Payment()
func Payment(
	ctx context.Context,
	tk string,
) error {
	// Login
	// tk ---> bob
	// user := token.ValidateToken(tk)
	// 上文
	ctx1 := context.WithValue(ctx, UserName{}, "bob")
	ctx2 := context.WithValue(ctx1, UserRole{}, "admin")

	// 下文
	DoPayment(ctx2)
	return nil
}

// 执行Payment的操作, 主流程 payment goroutine, cancel 监听外部信号
func DoPayment(ctx context.Context) {
	fmt.Println(ctx)
	fmt.Printf("%#v\n", ctx)

	// commit
	fmt.Println(ctx.Value(UserName{}))
	fmt.Println(ctx.Value(UserRole{}))

	// for {
	// 	select {
	// 	// 取消
	// 	case <-ctx.Done():
	// 	default:
	// 		fmt.Println(ctx.Value(UserName{}))
	// 		fmt.Println(ctx.Value(UserRole{}))

	// 		// orm.update(xxx)
	// 		// 数据库操作

	// 		// select {
	// 		// case <-ctx.Done():
	// 		// 	// 直接退出
	// 		// default:
	// 		// 	// 操作数据库
	// 		// }
	// 	}
	// }
}
