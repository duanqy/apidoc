// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// app 提供了一些公共的函数、结构体及代码级别的设置项。
package app

import "time"

// 一些代码级别的配置项。
// 可运行 go test 来检测常量是否符合规范。
const (
	// 版本号
	//
	// 版本号按照 http://semver.org 中的规则
	Version = "3.0.0-alpha+20160721"

	// 程序的正式名称
	Name = "apidoc"

	// 源码仓库地址
	RepoURL = "https://github.com/caixw/apidoc"

	// 官网
	OfficialURL = "http://apidoc.tools"

	// 配置文件名称。
	ConfigFilename = ".apidoc.json"

	// 默认的文档标题
	DefaultTitle = "APIDOC"

	// 默认的分组名称，在不指定分组名称的时候，
	// 系统会给其加到此分组中，同时也是默认的索引文件名。
	DefaultGroupName = "index"

	// 默认的时间格式，仅对 HTML 中的默认模板启作用。自定义模板中可自定义格式。
	TimeFormat = time.RFC3339

	// 输出的 profile 文件的名称
	Profile = "apidoc.prof"

	// 需要解析的最小代码块，小于此值，将不作解析
	MiniSize = len("@api ")
)