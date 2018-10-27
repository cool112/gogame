package utils

import (
	"unicode"
)

//大驼峰字符串转常量表达式 eg:Abc2Def -> ABC2_DEF
func CapCamel2Const(name string) string {
	chars := []rune(name)
	var newRunes = make([]rune,0)
	for _, char := range chars {
		if unicode.IsUpper(char) {
			if len(newRunes) > 0 {
				newRunes = append(newRunes, '_')
			}
			newRunes = append(newRunes, char)
		} else {
			newRunes = append(newRunes, unicode.ToUpper(char))
		}
	}
	return string(newRunes)
}
