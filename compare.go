package main

import (
	"fmt"
	"slices"
)

func Compare(a string, b string) (err error) {
	var newA, newB []string

	// 从 A B 两个文件中获取所有的包名
	appsFromFileA, err := GetAllAppsFromFile(a)
	if err != nil {
		return err
	}
	appsFromFileB, err := GetAllAppsFromFile(b)
	if err != nil {
		return err
	}

	// 切片去重
	appsFromFileA = slices.Compact(appsFromFileA)
	appsFromFileB = slices.Compact(appsFromFileB)

	// 找出只存在于 A 中的包名
	for _, app := range appsFromFileA {
		if slices.Contains(appsFromFileB, app) {
			continue
		}
		newA = append(newA, app)
	}
	// 找出只存在于 B 中的包名
	for _, app := range appsFromFileB {
		if slices.Contains(appsFromFileA, app) {
			continue
		}
		newB = append(newB, app)
	}

	// 打印对比结果
	fmt.Println("package only in A:")
	for _, na := range newA {
		fmt.Println(na)
	}

	fmt.Println("")

	fmt.Println("package only in B:")
	for _, nb := range newB {
		fmt.Println(nb)
	}

	return nil
}
