// Copyright 2020 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gvalid_test

import (
	"fmt"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/util/gvalid"
)

func ExampleCheckMap() {
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := []string{
		"passport@required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
		"password2@required|length:6,16#",
	}
	if e := gvalid.CheckMap(params, rules); e != nil {
		fmt.Println(e.Map())
		fmt.Println(e.FirstItem())
		fmt.Println(e.FirstString())
	}
	// May Output:
	// map[required:账号不能为空 length:账号长度应当在6到16之间]
	// passport map[required:账号不能为空 length:账号长度应当在6到16之间]
	// 账号不能为空
}

func ExampleCheckMap2() {
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := []string{
		"passport@length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
		"password2@required|length:6,16#",
	}
	if e := gvalid.CheckMap(params, rules); e != nil {
		fmt.Println(e.Map())
		fmt.Println(e.FirstItem())
		fmt.Println(e.FirstString())
	}
	// Output:
	// map[same:两次密码输入不相等]
	// password map[same:两次密码输入不相等]
	// 两次密码输入不相等
}

// Empty string attribute.
func ExampleCheckStruct() {
	type Params struct {
		Page      int    `v:"required|min:1         # page is required"`
		Size      int    `v:"required|between:1,100 # size is required"`
		ProjectId string `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := gvalid.CheckStruct(obj, nil)
	fmt.Println(err)
	// Output:
	// <nil>
}

// Empty pointer attribute.
func ExampleCheckStruct2() {
	type Params struct {
		Page      int       `v:"required|min:1         # page is required"`
		Size      int       `v:"required|between:1,100 # size is required"`
		ProjectId *gvar.Var `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := gvalid.CheckStruct(obj, nil)
	fmt.Println(err)
	// Output:
	// <nil>
}

// Empty integer attribute.
func ExampleCheckStruct3() {
	type Params struct {
		Page      int `v:"required|min:1         # page is required"`
		Size      int `v:"required|between:1,100 # size is required"`
		ProjectId int `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := gvalid.CheckStruct(obj, nil)
	fmt.Println(err)
	// Output:
	// project id must between 1, 10000
}
