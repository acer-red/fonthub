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
	Version   string `json:"version"`
	SHA256    string `json:"sha256"`
}

func GetInfo(fontBytes []byte) (Font, error) {

	f, err := sfnt.Parse(fontBytes)
	if err != nil {
		log.Error(err)
		return Font{}, err
	}

	familyName, err := getFontName(f, sfnt.NameIDFamily)
	if err != nil {
		fmt.Println("获取字体家族名称失败:", err)
	}

	typographicyName, err := getFontName(f, sfnt.NameIDTypographicFamily)
	if err != nil {
		fmt.Println("NameIDTypographicFamily err:", err)
	}

	subfamilyName, err := getFontName(f, sfnt.NameIDTypographicSubfamily)
	if err != nil {
		fmt.Println("NameIDTypographicSubfamily失败:", err)
	}

	version, err := getFontName(f, sfnt.NameIDVersion)
	if err != nil {
		fmt.Println("获取字体版本失败:", err)
	}

	copyright, err := getFontName(f, sfnt.NameIDCopyright)
	if err != nil {
		fmt.Println("获取版权信息失败:", err)
	}

	NameIDLicense, err := getFontName(f, sfnt.NameIDLicense)
	if err != nil {
		fmt.Println("获取版权信息失败:", err)
	}

	return Font{
		Name:      typographicyName,
		FullName:  familyName,
		SubName:   subfamilyName,
		CopyRight: copyright,
		License:   NameIDLicense,
		Version:   version,
		SHA256:    Sha256(fontBytes),
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
