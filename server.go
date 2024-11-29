package main
import(
	"Gocode/rpc"
	"net"
)
type Arith struct{
	One int
	Two int
}
func(a Arith)Add() int{
	return a.One+a.Two
}
func(a Arith)Divide() int{
	return a.One/a.Two
}
func(a Arith)SayHi() string{
	return "Hello!"
}
func main(){
	MyServer:=rpc.NewServer()
	arith:=Arith{}
	MyServer.Register("Arith",arith)
	listener,_:=net.Listen("tcp","localhost:7777")
	for{
		conn,_:=listener.Accept()
		go MyServer.HandleRequest(conn)	
	}
}