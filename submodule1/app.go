package submodule1 //  packageはディレクトリ名と一致させる

import "fmt"

func Func1() { // 1文字目は大文字
	fmt.Println("submodele1/app.goの関数")
}

// 1文字目は大文字
const Text string = "submodele1/app.goの定数"
