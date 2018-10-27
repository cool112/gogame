package main
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
	
	constArr = append(constArr, &ConstEntity{"MOD_ACC", proto.MOD_ACC})
	
	constArr = append(constArr, &ConstEntity{"MOD_LOG", proto.MOD_LOG})
	
	constArr = append(constArr, &ConstEntity{"MOD_HTTP", proto.MOD_HTTP})
	
	constArr = append(constArr, &ConstEntity{"LOGIN", proto.LOGIN})
	
	constArr = append(constArr, &ConstEntity{"AUTH", proto.AUTH})
	
	constArr = append(constArr, &ConstEntity{"ACK", proto.ACK})
	
	constArr = append(constArr, &ConstEntity{"LOG_COMMON", proto.LOG_COMMON})
	
	constArr = append(constArr, &ConstEntity{"HTTP_TEST", proto.HTTP_TEST})
	
	constArr = append(constArr, &ConstEntity{"HTTP_THTML", proto.HTTP_THTML})
	
	constArr = append(constArr, &ConstEntity{"HTTP_PATCH", proto.HTTP_PATCH})
	
	constArr = append(constArr, &ConstEntity{"HTTP_PROF", proto.HTTP_PROF})
	
	constArr = append(constArr, &ConstEntity{"HTTP_TRACE", proto.HTTP_TRACE})
	
	constArr = append(constArr, &ConstEntity{"HTTP_LOG_LIST", proto.HTTP_LOG_LIST})
	
	for _, val := range constArr {
		fmt.Fprintf(file, "%s=%d\n", val.name, val.val)
	}
}
