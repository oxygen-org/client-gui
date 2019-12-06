package utils

import (
	"fmt"
	"testing"
	"os"
)

func TestShow(t *testing.T) {
	fmt.Println(getLang())
	fmt.Println(words)
	fmt.Println(Words)
}

func TestTr(t *testing.T)  {
	os.Setenv("LANG","en")
	if Tr("中国") != "China"{
		t.Error("设置英文环境未生效")
	}
	os.Setenv("LANG","zh")
	if Tr("中国") != "中国"{
		t.Error("设置中文环境未生效")
	}
	os.Setenv("LANG","sp")
	if Tr("中国") != "中国"{
		t.Error("缺省环境未生效")
	}

}
