package main

import (
	"flag"
	"fmt"
	_ "io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var path = flag.String("path", ".", "path where .proto files in")
var GO_PATH = os.Getenv("GOPATH")
var isWin = strings.Contains(runtime.GOOS, "win")

func main() {
	log.Print(isWin)
	flag.Parse()
	root, err := os.Stat(*path)
	if err != nil {
		log.Fatal(err)
	}
	if !root.IsDir() {
		log.Fatal("path must a dir!")
	}
	walkDir(*path)
	//files,err:=ioutil.ReadDir(path)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//os.FileInfo
	//files,err=filepath.Glob(*path+"/*.proto")
	//if err != nil{
	//	log.Fatal(err)
	//}
	//for _,file := range files{
	//	log.Printf("file :%s", file)
	//}
}

func walkDir(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path == dir {
			return nil
		}
		if err != nil {
			log.Panic(err)
		}
		if info.IsDir() {
			//			walkDir(path)//Walk auto recursive
		} else {
			if strings.HasSuffix(path, ".proto") {
				curDir := filepath.Dir(path)
				filename := filepath.Base(path)
				filename = filename[:len(filename)-6]
				log.Printf("start generate file :%s", path)
				gobin := GO_PATH + "/bin"
				suffix := ""
				if isWin {
					suffix = ".exe"
				}
				arg1:=fmt.Sprintf("--plugin=protoc-gen-gogofaster=%s/protoc-gen-gogofaster%s", gobin, suffix)
				arg2:=fmt.Sprintf("--gogofaster_out=%s", curDir)
				arg3:=fmt.Sprintf("--proto_path=%s", curDir)
				params := fmt.Sprintf(" --plugin=protoc-gen-gogofaster=%s/protoc-gen-gogofaster%s --gogofaster_out=%s --proto_path=%s %s", gobin, suffix, curDir, curDir, path)
				log.Print(params)
//				cmd := exec.Command(gobin+"/protoc.exe", params)
				cmd := exec.Command("protoc", arg1, arg2, arg3, path)
				out,_ := cmd.Output()
				log.Print(string(out))
//				arg1=fmt.Sprintf("--plugin=protoc-gen-msg=%s/protoc-gen-msg%s", gobin, suffix)
//				log.Print(arg1)
//				arg2=fmt.Sprintf("--msg_out=%s_reg.go:%s", filename, curDir)
//				log.Print(arg2)
//				cmd = exec.Command("protoc", arg1, arg2, arg3, path)
//				out,_= cmd.Output()
//				log.Print(string(out))
				arg1=fmt.Sprintf("--plugin=protoc-gen-msg=%s/protoc-reg-plugin%s", gobin, suffix)
				log.Print(arg1)
				arg2=fmt.Sprintf("--msg_out=%s_reg.go:%s", filename, curDir)
				log.Print(arg2)
				cmd = exec.Command("protoc", arg1, arg2, arg3, path)
				out,_= cmd.Output()
				log.Print(string(out))
				//c# compiler
//				arg1=fmt.Sprintf("-I=%s", curDir)
//				arg2=fmt.Sprintf("--csharp_out=%s", curDir)
//				cmd = exec.Command("protoc", arg1, arg2, path)
//				out,_= cmd.Output()
//				log.Print(string(out))
			}
		}
		return nil
	})
}
