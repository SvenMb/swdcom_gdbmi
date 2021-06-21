package main

import (
	"fmt"
	"time"
	"github.com/cyrus-and/gdb"
// 	"io"
//	"os"
	"strconv"
)

func main() {

	// start a new instance and pipe the target output to stdout
	gdb, err := gdb.NewCmd( []string{"gdb-multiarch", "--quiet", "--interpreter=mi2"}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("waiting for swdcom to start...")
	// fmt.Println(gdb)
	// go io.Copy(os.Stdout, gdb)

	// evaluate an expression
	// gdb.Send("var-create", "RXTX", "*", "$r11")
	// result,_ := gdb.Send("data-list-register-values","d","11")
	// fmt.Println(result)
	// fmt.Println(result["payload"].(map[string]interface{})["register-values"].([]interface{})[0].(map[string]interface{})["value"])

	// fmt.Println(gdb.Send("interpreter-exec console \"info reg\""))

	result,err := gdb.Send("exec-run")
	// fmt.Println(result)
	// fmt.Println(err)
	if (err != nil || result["class"] == "error") {
		fmt.Println("Error executing, probably openocd, st-util not started or not connected?")
		fmt.Println("GDB says (often misleading):")
		fmt.Println(result["payload"].(map[string]interface{})["msg"])
		fmt.Println("Error: ",err)
		return
	}

	time.Sleep(time.Second / 2)
	gdb.Interrupt()

	result,err = gdb.Send("data-list-register-values","d","11")
	// fmt.Println(result)
	// fmt.Println(err)
	if (err != nil || result["class"] == "error") {
		fmt.Println("Error reading r11, probably wrong gdb-architecture?")
		fmt.Println("GDB says (often misleading):")
		fmt.Println(result["payload"].(map[string]interface{})["msg"])
		fmt.Println("Error: ",err)
		return
	}

	R11,_ := strconv.Atoi(result["payload"].(map[string]interface{})["register-values"].([]interface{})[0].(map[string]interface{})["value"].(string))
	rxb := R11+4
	txb := R11+260
	fmt.Println("R11: ",R11)
	fmt.Println("RXB: ",rxb)
	fmt.Println("TXB: ",txb)

	result,_ = gdb.Send("data-read-memory",strconv.Itoa(R11),"u","1","1","4")
	// fmt.Println(result)

	RXTX := result["payload"].(map[string]interface{})["memory"].([]interface{})[0].(map[string]interface{})["data"].([]interface{})
	rxw,_ := strconv.Atoi(RXTX[0].(string))
	rxr,_ := strconv.Atoi(RXTX[1].(string))
	txw,_ := strconv.Atoi(RXTX[2].(string))
	txr,_ := strconv.Atoi(RXTX[3].(string))
	fmt.Println("rxw: ",rxw)
	fmt.Println("rxr: ",rxr)
	fmt.Println("txw: ",txw)
	fmt.Println("txr: ",txr)

	// print out the current 
	
	result,_ = gdb.Send("data-read-memory",strconv.Itoa(txb),"u","1","1",strconv.Itoa(txw-txr))
	// fmt.Println("\n\n",result)
	txt := result["payload"].(map[string]interface{})["memory"].([]interface{})[0].(map[string]interface{})["data"].([]interface{})

	for _,c := range txt {
		cint,_ := strconv.Atoi(c.(string))
		// fmt.Println(cint, ": ", string(cint))
		fmt.Print(string(cint))
	}

	gdb.Exit()
}
