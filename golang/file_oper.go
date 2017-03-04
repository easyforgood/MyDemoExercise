package main
import "os"
import "fmt"

func main(){

		f, err := os.OpenFile("golangtestfile", os.O_APPEND| os.O_CREATE | os.O_WRONLY , 0600) 

		fmt.Println(err)
		n,err := f.WriteString("setest1" + "\n") 

		fmt.Println(n)
		fmt.Println(err)

		n, err = f.WriteString("setest2" + "\n") 

		fmt.Println(n)
		fmt.Println(err)
		f.Close()

}
