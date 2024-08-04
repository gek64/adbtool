package main

import (
	"bufio"
	"bytes"
	"github.com/gek64/gek/gExec"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// GetAppsFromFile 从文件中读取APP列表
func GetAppsFromFile(appFile string) (apps []string, err error) {
	// 打开文件
	file, err := os.Open(appFile)
	if err != nil {
		return nil, err
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	// 按行读取文件中数据
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 跳过注释
		if strings.Contains(scanner.Text(), "//") {
			continue
		}

		// 截取每行 package: 后半的包名,并存储到apps切片中
		apps = append(apps, strings.ReplaceAll(scanner.Text(), "package:", ""))
	}
	return apps, nil
}

// GetAppListFromADB 使用adb获取全部应用列表
func GetAppListFromADB() (apps []string, err error) {
	cmd := exec.Command("adb", "shell", "pm", "list", "packages")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	bytesReader := bytes.NewReader(output)
	// 按行读取数据
	scanner := bufio.NewScanner(bytesReader)
	for scanner.Scan() {
		// 截取每行 package: 后半的包名,并存储到apps切片中
		apps = append(apps, strings.ReplaceAll(scanner.Text(), "package:", ""))
	}
	return apps, nil
}

func PMClear(app string) error {
	return gExec.Run(exec.Command("adb", "shell", "pm", "clear", app))
}

func PMUninstall(app string) error {
	return gExec.Run(exec.Command("adb", "shell", "pm", "uninstall", app))
}

func PMUninstallUser(app string, uid int) error {
	return gExec.Run(exec.Command("adb", "shell", "pm", "uninstall", "--user", strconv.Itoa(uid), app))
}

func PMReinstall(app string) error {
	return gExec.Run(exec.Command("adb", "shell", "cmd", "package", "install-existing", app))
}

func PMDisable(app string) error {
	return gExec.Run(exec.Command("adb", "shell", "pm", "disable", app))
}

func PMDisableUser(app string, uid int) error {
	return gExec.Run(exec.Command("adb", "shell", "pm", "disable-user", "--user", strconv.Itoa(uid), app))
}

func PMEnable(app string) error {
	return gExec.Run(exec.Command("adb", "shell", "pm", "enable", app))
}
