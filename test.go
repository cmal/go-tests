package main

import (
	"io"
	"os"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"html/template"
	"bytes"
	"bufio"
)

func readJson(reader io.Reader, value interface{}) error {
	fmt.Println("in readJson")
	data, err := ioutil.ReadAll(reader)
	// fmt.Println(data, err)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	fmt.Println(data, value)

	err = json.Unmarshal(data, value)

	fmt.Println(err, data, value)

	if err = json.Unmarshal(data, &value); err != nil {
		fmt.Println("unmarshal ok")
		if syntaxerr, ok := err.(*json.SyntaxError); ok {
			fmt.Println("readJson ok")
			line := findLine(data, syntaxerr.Offset)
			return fmt.Errorf("JSON syntax error at line %v: %v", line, err)
		}
		return err
	}
	return nil
}

// findLine returns the line number for the given offset into data.
func findLine(data []byte, offset int64) (line int) {
	fmt.Println("in findLine")
	line = 1
	for i, r := range string(data) {
		if int64(i) >= offset {
			return
		}
		if r == '\n' {
			line++
		}
	}
	return
}


func readJsonFile(fn string, value interface{}) error {
	fmt.Println("in readJsonFile")
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	err = readJson(file, value)
	if err != nil {
		return fmt.Errorf("%s in file %s", err.Error(), fn)
	}
	return nil
}

func main() {

	// stringT := reflect.TypeOf("")
	// // m := reflect.New(reflect.MapOf(stringT, stringT))
	// // TODO holy! 搞不定上面的用法

	// type Animal struct {
	//	Name  string
	//	Order string
	// }
	// var m []Animal
	// readJsonFile("/Users/yuzhao/go-tests/demo.json", &m)
	// fmt.Printf("%+v", m)
	// fmt.Println(m)
	// t1 := reflect.TypeOf(m[0])
	// fmt.Println(t1)
	// fmt.Println(stringT)

	t := template.New("struct data demo template") // 创建一个模板
	t, _ = t.Parse("hello, {{.UserName}}! \n")     // 解析模板文件

	type Actor struct {
		UserName template.HTML
	}
	actor := Actor{UserName: template.HTML("<div></div>jsrush@structMap")}   // 创建一个数据对象
	t.Execute(os.Stdout, actor)                    // 执行模板的merger操作，并输出到控制台

	t2, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t2.ExecuteTemplate(os.Stdout, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	fmt.Println(err)


	// type Reader interface {
	//	Read(p []byte) (n int, err error)
	// }

	var r io.Reader
	r = new(bytes.Buffer)
	fmt.Println(reflect.TypeOf(r))
	r = bufio.NewReader(r)
	fmt.Println(reflect.TypeOf(r))


	r = os.Stdin
	fmt.Println(reflect.TypeOf(r), reflect.TypeOf(os.Stdin))
}
