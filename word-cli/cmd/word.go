package cmd

import (
	"log"

	"github.com/JrMarcco/go-learning/word-cli/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamel
	ModeUnderscoreToLowerCamel
	ModeCamelToUnderscore
)

var desc = `
该子命令支持各种单词格式转换，模式如下：
   1：全部转大写
   2：全部转小写
   3：下划线转大写驼峰
   4：下划线转小写驼峰
   5：驼峰转下划线
`

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamel:
			content = word.UnderscoreToUpperCamel(str)
		case ModeUnderscoreToLowerCamel:
			content = word.UnderscoreToLowerCamel(str)
		case ModeCamelToUnderscore:
			content = word.CamelToUnderscore(str)
		default:
			log.Fatalln("Unsupported mode")
		}

		log.Printf("Res: %s\n", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "word", "w", "", "Please input word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Please input mode")
}
