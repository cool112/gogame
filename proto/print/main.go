package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"github.com/cool112/gogame/proto"
	"os"
	"regexp"
	"text/template"
)

const tmpl = `package main
import(
	"fmt"
	"io/ioutil"
	"github.com/cool112/gogame/proto"
)

type ConstEntity struct{
	name string
	val int
}

func main(){
	file,err:=ioutil.TempFile(".", "protoout")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	constArr := make([]*ConstEntity, 0)
	{{range .}}
	constArr = append(constArr, &ConstEntity{"{{.}}", proto.{{.}}})
	{{end}}
	for _, val := range constArr {
		fmt.Fprintf(file, "%s=%d\n", val.name, val.val)
	}
}
`

var (
	CONST_START = regexp.MustCompile(`^const \($`)
	CONST_END   = regexp.MustCompile(`^\)$`)
	ENUM_STATE  = regexp.MustCompile(`^\s*([A-Z_]+).*$`)
)

const (
	PARSE_READY = iota
	PARSE_START
)

var curState int
//根据proto.go文件打印出所有协议常量,注意只要是常量就会被打印
//会生成一个新的程序源码,要执行生成的程序才会在工作目录下生成一个pringprotoXXX的文件
func main() {
	mode := os.O_CREATE | os.O_RDWR
	outfile, err := os.OpenFile("./printproto.go", mode, 0666)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	protoFile, err := os.OpenFile("src/github.com/cool112/gogame/proto/proto.go", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer protoFile.Close()
	reader := bufio.NewScanner(protoFile)
	protoArr := make([]string, 0)
	for reader.Scan() {
		line := reader.Text()
		fmt.Println(line)
		switch curState {
		case PARSE_READY:
			fmt.Println("ready")
			if CONST_START.MatchString(line) {
				fmt.Println("to start")
				curState = PARSE_START
			}
		case PARSE_START:
			fmt.Println("start")
			if ENUM_STATE.MatchString(line) {
				protoArr = append(protoArr, ENUM_STATE.FindStringSubmatch(line)[1])
				fmt.Printf("%#v\n", protoArr)
			} else if CONST_END.MatchString(line) {
				fmt.Println("to ready")
				curState = PARSE_READY
			}
		}
	}
	t := template.Must(template.New("newgo").Parse(tmpl))
	t.Execute(outfile, protoArr)
}

//type ConstEntity struct{
//	name string
//	val int
//}
//func test() {
//	file, err := ioutil.TempFile(".", "protoout")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	constArr := make([]*ConstEntity, 0)
//	constArr = append(constArr, &ConstEntity{"LOG_COMMON", proto.LOG_COMMON})
//	for _, val := range constArr {
//		fmt.Fprintf(file, "%s=%d\n", val.name, val.val)
//	}
//}
