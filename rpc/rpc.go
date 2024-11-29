package rpc
import(
	"fmt"
	"reflect"
	"encoding/json"
	"net"
)
type Call struct{
	Id int
	StructName string
	MethodName string
	Structure interface{}
	Reply interface{}
}
type Server struct{
	AddressList []string
	Structure map[string] reflect.Type
	Funcs map[string] reflect.Method
}
func NewServer() Server{
	server:=Server{}
	server.Funcs=make(map[string]reflect.Method)
	server.Structure=make(map[string]reflect.Type)
	return server
}
func (server Server)Register(name string,Member interface{}){
	typ:=reflect.TypeOf(Member)
	server.Structure[name]=typ
	for i:=0;i<typ.NumMethod();i++{
		server.Funcs[typ.Method(i).Name]=typ.Method(i)
	}
}
func (server Server)ShowServers(){
	fmt.Println(server.Structure["Arith"])
	for i,_:=range server.Funcs{
		fmt.Println(i)
	}
}
func CallServer(StructName string,MethodName string,target interface{},conn net.Conn){
	encoder:=json.NewEncoder(conn)
	decoder:=json.NewDecoder(conn)
	call:=new(Call)
	call.MethodName=MethodName
	call.Id=1
	call.StructName=StructName
	call.Structure=target
	call.Reply=0
	encoder.Encode(call)
	decoder.Decode(&call)
	fmt.Println("Result is:",call.Reply)
}
func (server Server)HandleRequest(conn net.Conn){
	encoder:=json.NewEncoder(conn)
	decoder:=json.NewDecoder(conn)
	for{
	response:=new(Call)
	decoder.Decode(&response)
	if response.Id==0{
		conn.Close()
		return 
	}
	method:=server.Funcs[response.MethodName]
	tmpstruct:=reflect.New(server.Structure[response.StructName]).Elem()
	switch s:=response.Structure.(type){
		case map[string]interface{}:
			for k,v:=range s{
			field:=tmpstruct.FieldByName(k)
			switch v:=v.(type){
			case int:
				field.Set(reflect.ValueOf(v))
			case float64:
				field.Set(reflect.ValueOf(int(v)))
			}
		}
		default:
			fmt.Println("not a struct")
		}
	reply:=method.Func.Call([]reflect.Value{tmpstruct})
	response.Reply=reply[0].Interface()
	encoder.Encode(response)
	}
}
