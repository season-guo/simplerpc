package main
import(
	"net"
	"Gocode/rpc"
	"time"
)
type Arith struct{
	One int
	Two int
}
func main(){
	conn,_:=net.Dial("tcp","localhost:7777")
	myarith:=Arith{4,5}
	secondarith:=Arith{30,7}
	rpc.CallServer("Arith","SayHi",&myarith,conn)
	rpc.CallServer("Arith","Add",&myarith,conn)
	rpc.CallServer("Arith","Add",&secondarith,conn)
	rpc.CallServer("Arith","Divide",&secondarith,conn)
}