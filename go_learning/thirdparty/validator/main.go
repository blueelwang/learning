package main

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

// RegisterReq _
type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username string `validate:"gt=0"`
	// 同上
	PasswordNew string `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email string `validate:"email"`
}

func main() {
	validate := validator.New()
	data := RegisterReq{
		Username:       "zhang3",
		PasswordNew:    "li4",
		PasswordRepeat: "li4",
		Email:          "sunjianwei1@xiaomi.com",
	}
	err := validate.Struct(data)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("OK!")
	}
}
