package main 

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 获取模块名参数
	module := flag.String("module", "====", "Input your module name")
	module2 := flag.String("module2", "====2", "Input your module name2")
	flag.Parse()
	// help := flag.Usage()
	help := flag.Args()//非-option形式的参数
	help2 := flag.NArg()
	help3 := flag.NFlag()
	flag.PrintDefaults()

	// fmt.Println(help)
	fmt.Println(help)
	fmt.Println(help2)
	fmt.Println(help3)
	fmt.Println("======++++++++++++=======")
	fmt.Println(*module)
	// fmt.Println(*module.Name)
	fmt.Println(*module2)

	// 拷贝模块模板并重命名（文件夹、代码文件、测试文件）
	current_dir, _ := os.Getwd()
	fmt.Println(current_dir)

	dirhandle, err := os.Open(current_dir + "/")
	if err != nil {
		panic(err)
	}
	defer dirhandle.Close()

	fis, err := dirhandle.Readdir(0)
	fmt.Println(fis)
	for _,f := range fis {
		fmt.Println(f.Name())

		if f.IsDir() {
			fmt.Println("dir+++++++++++++++++++++++")

		} else {
			//如果是普通文件 直接写入 dir 后面已经有了 /
			fmt.Println("file=====================")
		}
	}
}