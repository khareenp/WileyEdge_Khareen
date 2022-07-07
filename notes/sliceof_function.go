package notes
import "fmt"


func A( ){
    fmt.Println("Calling A" )
}
func B(t int){
    fmt.Println("Calling B" ,t)
}

func C(A string ,f float32){

    fmt.Println("Calling C" ,A , " ", f)
}
func Sliceof() {
	//creates slice of functions
    f:=[] interface{}{A,B,C}
	
    f[0].(func())()
    f[1].(func(int))(15)
    f[2].(func(string,float32))("Hello",123)



}