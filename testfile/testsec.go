package testfile

import (
	"os"
	"os/exec"
)

func test1111() {
	exec.Command("ls", "-la").Run() // 不安全的外部命令调用
	os.Chmod("myfile.txt", 0777)    // 不安全的文件权限
}
