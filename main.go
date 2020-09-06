package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	var (
		data = make(map[string]interface{})
		// result      = make(map[string]string)

	)
	fmt.Println("Input your json text (not support nested and type @ to finish):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('@') //end of input
	input = strings.Replace(input, "@", "", len(input)-1)
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Println("can not unmarshal data with error ", err)
		return
	}
	fmt.Println()
	fmt.Println("-------------DEFINITION---------------")
	fmt.Println()
	definitions := definition(data)
	fmt.Println(definitions)
}

func formatName(key string) string {
	key = strings.ReplaceAll(key, "_", " ")
	key = strings.Title(key)
	key = strings.ReplaceAll(key, " ", "")
	return key
}

func definition(data map[string]interface{}) string {
	var (
		definitions = ""
		template    = "%s%s"
		tabJSON     = "`json:\"%s\"`"
		start       = "type Object struct{\n"
		line        = "%s\t%s\t%s\n"
		end         = "}\n"
	)
	definitions = start
	for key, value := range data {
		tab := fmt.Sprintf(tabJSON, key)
		key = formatName(key)
		var item = ""
		switch reflect.TypeOf(value).Kind() {
		case reflect.Float64:
			valueInt := cast.ToInt64(value.(float64) * 1e6)
			if valueInt%1e6 == 0 {
				item = fmt.Sprintf(line, key, reflect.TypeOf(valueInt), tab)
			} else {
				item = fmt.Sprintf(line, key, reflect.TypeOf(value), tab)
			}
		default:
			item = fmt.Sprintf(line, key, reflect.TypeOf(value), tab)
		}
		definitions = fmt.Sprintf(template, definitions, item)
	}
	definitions = fmt.Sprintf(template, definitions, end)
	return definitions
}
