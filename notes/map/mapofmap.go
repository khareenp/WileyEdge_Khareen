package map1
import( "fmt"
)

func mapofmap(){
    mpofmp:=map[int] map[string]string{
        1 : {
            "A":"Apple",
            "B": "Banana",
        },
        2:{
            "Z":"Zebra",
            "Y": "Yalk",
        },
    }
    fmt.Println(mpofmp)

}