package system

import (
	"fmt"

	"github.com/tengfei-xy/go-log"
	"golang.org/x/image/font/sfnt"
)

type Font struct {
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
	SubName   string `json:"sub_name"`
	CopyRight string `json:"copy_right"`
	License   string `json:"license"`
}

func GetInfo(fontBytes []byte) (Font, error) {

	// 1. 读取字体文件内容
	// fontBytes, err := os.ReadFile(filename)
	// if err != nil {
	// 	log.Error(err)
	// 	return font{}, err
	// }

	// 2. 解析字体数据
	f, err := sfnt.Parse(fontBytes)
	if err != nil {
		log.Error(err)
		return Font{}, err
	}

	// 3. 获取字体信息

	// 定义语言偏好，通常使用中文和英文
	// tags := []language.Tag{language.Chinese, language.English}

	// 获取字体家族名称 (Font Family Name)
	familyName, err := getFontName(f, sfnt.NameIDFamily)
	if err != nil {
		fmt.Println("获取字体家族名称失败:", err)
	} else {
		fmt.Println("字体家族名称:", familyName)
	}

	typographicyName, err := getFontName(f, sfnt.NameIDTypographicFamily)
	if err != nil {
		fmt.Println("NameIDTypographicFamily err:", err)
	} else {
		fmt.Println("NameIDTypographicFamily:", typographicyName)
	}

	subfamilyName, err := getFontName(f, sfnt.NameIDTypographicSubfamily)
	if err != nil {
		fmt.Println("NameIDTypographicSubfamily失败:", err)
	} else {
		fmt.Println("NameIDTypographicSubfamily:", subfamilyName)
	}

	// 获取字体版本 (Version)
	version, err := getFontName(f, sfnt.NameIDVersion)
	if err != nil {
		fmt.Println("获取字体版本失败:", err)
	} else {
		fmt.Println("字体版本:", version)
	}

	// 获取版权信息 (Copyright)
	copyright, err := getFontName(f, sfnt.NameIDCopyright)
	if err != nil {
		fmt.Println("获取版权信息失败:", err)
	} else {
		fmt.Println("版权信息:", copyright)
	}

	NameIDLicense, err := getFontName(f, sfnt.NameIDLicense)
	if err != nil {
		fmt.Println("获取版权信息失败:", err)
	} else {
		fmt.Println("版权信息:", NameIDLicense)
	}
	NameIDLicenseURL, err := getFontName(f, sfnt.NameIDLicenseURL)
	if err != nil {
		fmt.Println("NameIDLicenseURL失败:", err)
	} else {
		fmt.Println("NameIDLicenseURL:", NameIDLicenseURL)
	}

	NameIDDescription, err := getFontName(f, sfnt.NameIDDescription)
	if err != nil {
		fmt.Println("NameIDDescription失败:", err)
	} else {
		fmt.Println("NameIDDescription:", NameIDDescription)
	}

	// - 字体描述 (sfnt.NameIDDescription)
	// - 设计者 (sfnt.NameIDDesigner)
	// - 供应商 (sfnt.NameIDManufacturer)
	// - 样式 (sfnt.NameIDStyle)

	return Font{
		Name:      typographicyName,
		FullName:  familyName,
		SubName:   subfamilyName,
		CopyRight: copyright,
		License:   NameIDLicense,
	}, nil
}

// getFontName 函数用于从字体信息中根据 NameID 和语言偏好获取名称
func getFontName(font *sfnt.Font, nameID sfnt.NameID) (string, error) {
	name, err := font.Name(nil, nameID)
	if err != nil {
		return "", err
	}
	return name, nil
}
