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
	fmt.Println("Input your json text (not support nested and type ^ to finish):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('^') //end of input
	input = strings.Replace(input, "^", "", len(input)-1)
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Println("can not unmarshal data with error ", err)
		return
	}

	swagger, definitions := definition(data)
	fmt.Println()
	fmt.Println("-------------SWAGGER---------------")
	fmt.Println()
	fmt.Println(swagger)
	fmt.Println()
	fmt.Println("-------------DEFINITION---------------")
	fmt.Println()
	fmt.Println(definitions)
}

func formatName(key string) string {
	key = strings.ReplaceAll(key, "_", " ")
	key = strings.Title(key)
	key = strings.ReplaceAll(key, " ", "")
	return key
}

func definition(data map[string]interface{}) (string, string) {
	var (
		definitions = ""
		template    = "%s%s"
		tabJSON     = "`json:\"%s\"`"
		start       = "type Object struct{\n"
		line        = "\t%s\t%s\t%s\n"
		end         = "}\n"

		swagger  = ""
		startSwg = "Object:\n\ttype: object\n\tproperties:\n"
		lineSwg  = "\t\t%s:\n\t\t\ttype: %s\n"
		endSwg   = "\n"
	)
	definitions = start
	swagger = startSwg
	for key, value := range data {
		if value == nil {
			continue
		}
		//swagger
		var item = ""
		switch reflect.TypeOf(value).Kind() {
		case reflect.Float64:
			valueInt := cast.ToInt64(value.(float64) * 1e6)
			if valueInt%1e6 == 0 {
				item = fmt.Sprintf(lineSwg, key, "integer")
			} else {
				item = fmt.Sprintf(lineSwg, key, "number")
			}
		case reflect.Slice:
			item = fmt.Sprintf(lineSwg, key, "array\n\t\t\titems: \n\t\t\t\ttype: object")
		case reflect.Map:
			item = fmt.Sprintf(lineSwg, key, "object")
		case reflect.Bool:
			item = fmt.Sprintf(lineSwg, key, "boolean")
		default:
			item = fmt.Sprintf(lineSwg, key, reflect.TypeOf(value))
		}
		swagger = fmt.Sprintf(template, swagger, item)

		tab := fmt.Sprintf(tabJSON, key)
		key = formatName(key)
		item = ""
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
	swagger = fmt.Sprintf(template, swagger, endSwg)
	definitions = fmt.Sprintf(template, definitions, end)
	return swagger, definitions
}
