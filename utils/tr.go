package utils

import (
	"os"
	"strings"
)

const (
	// LangZH 中文文本标识
	LangZH = "zh_CN"
	// LangEN 英文文本标识
	LangEN = "en_US"
)

func getLang() string {
	envLang := strings.ToLower(os.Getenv("LANG"))
	var lang string
	switch {
	case strings.Contains(envLang, "en"):
		lang = LangEN
	case strings.Contains(envLang, "zh"):
		lang = LangZH
	default:
		lang = LangZH
	}
	return lang
}



// Word 文本结构
type Word struct {
	Ori   string
	MapTo map[string]string
}

var words = []*Word{
	{Ori: "中国", MapTo: map[string]string{LangEN: "China"}},
	{Ori: "Oxygen计算平台客户端", MapTo: map[string]string{LangEN: "Oxygen Computing Platform Client"}},
	{Ori: "用户登录", MapTo: map[string]string{LangEN: "User Login"}},
	{Ori: "🐼：欢迎使用Oxygen计算平台", MapTo: map[string]string{LangEN: "🐼：Welcome Using Oxygen Computing Platform"}},
	{Ori: "用户名/邮箱", MapTo: map[string]string{LangEN: "User-Name/Email"}},
	{Ori: "密码", MapTo: map[string]string{LangEN: "Password"}},

	{Ori: "注册", MapTo: map[string]string{LangEN: "Register"}},
	{Ori: "登录", MapTo: map[string]string{LangEN: "Login"}},
	{Ori: "登录错误", MapTo: map[string]string{LangEN: "Login Error"}},
	{Ori: "密码错误或用户不存在", MapTo: map[string]string{LangEN: "Password Error or User not Exist"}},
	{Ori: "恭喜！登录成功", MapTo: map[string]string{LangEN: "Congratulations! Login success"}},
	{Ori: "你已经通过用户密码验证", MapTo: map[string]string{LangEN: "You have been authenticated with the user password"}},



}

// Words 多语言文本集合
var Words = make(map[string]*Word)

func init() {
	for _, w := range words {
		Words[w.Ori] = w
	}
}

func (w Word) String() string {
	return w.Ori
}

// Tr 文本语言转换
func Tr(text string) string {
	if _, ok := Words[text]; ok && getLang() != LangZH {
		return Words[text].MapTo[getLang()]
	} else {
		return text
	}
}
