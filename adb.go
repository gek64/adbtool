package main

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/unix755/xtools/xExec"
)

// GetAppsFromFile 从文件中读取 APP 列表, 排除包名前的注释
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
		// 去除整行中所有的 空格 与 package:
		app := strings.ReplaceAll(scanner.Text(), " ", "")
		app = strings.ReplaceAll(app, "package:", "")
		// 跳过注释及空白行
		if strings.Contains(app, "//") || strings.Contains(app, "#") || app == "" {
			continue
		}
		// 符合条件的 app 存储到 apps 切片中
		apps = append(apps, app)
	}
	return apps, nil
}

// GetAllAppsFromFile 从文件中读取 APP 列表, 不排除包名前的注释
func GetAllAppsFromFile(appFile string) (apps []string, err error) {
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
		// 去除整行中所有的 空格 与 package:
		app := strings.ReplaceAll(scanner.Text(), "package:", "")
		// 跳过注释及空白行
		if strings.Contains(app, "//") || strings.Contains(app, "# ") || strings.Contains(app, "#!") || app == "" {
			continue
		}
		// 去除整行中所有的 #
		app = strings.ReplaceAll(app, "#", "")
		// 符合条件的 app 存储到 apps 切片中
		apps = append(apps, app)
	}
	return apps, nil
}

// GetAppListFromADB 使用 adb 获取全部应用列表
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
	return xExec.Run(exec.Command("adb", "shell", "pm", "clear", app))
}

func PMUninstall(app string) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "uninstall", app))
}

func PMUninstallUser(app string, uid int) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "uninstall", "--user", strconv.Itoa(uid), app))
}

func PMReinstall(app string) error {
	return xExec.Run(exec.Command("adb", "shell", "cmd", "package", "install-existing", app))
}

func PMDisable(app string) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "disable", app))
}

func PMDisableUser(app string, uid int) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "disable-user", "--user", strconv.Itoa(uid), app))
}

func PMEnable(app string) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "enable", app))
}

func PMSuspend(app string) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "suspend", app))
}

func PMUnsuspend(app string) error {
	return xExec.Run(exec.Command("adb", "shell", "pm", "unsuspend", app))
}
