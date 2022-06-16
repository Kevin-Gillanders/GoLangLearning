package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
)

func fmtPrint(iter int) {

	for i := 0; i <= iter; i++ {
		fmt.Println("fmt")
	}

}

func bufPrint(iter int) {

	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	for i := 0; i <= iter; i++ {
		fmt.Fprintln(buf, "buf")
	}
}

func bufStackPrint(iter int) {

	buf := bufio.NewWriter(os.Stdout)
	var out [] string
	defer buf.Flush()
	for i := 0; i <= iter; i++ {
		out = append(out, "bufBuild\n")
	}
	fmt.Fprintln(buf, out)
}

func sysCallPrint(iter int) {
	for i := 0; i <= iter; i++ {
		syscall.Write(1, []byte("sysNoStack"))
	}

}


func sysCallPrintStr(row string) {
	syscall.Write(1, []byte(row))
}

func osPrint(iter int) {

	for i := 0; i <= iter; i++ {
		os.Stdout.Write([]byte("osOut\n"))
	}
}

func osStackPrint(iter int) {
	var outPut []byte
	for i := 0; i <= iter; i++ {
		outPut = append(outPut, []byte("osStackPrint\n")...)
	}
	os.Stdout.Write(outPut)

}


func sysCallPrintContent(content [][] rune) {
	for _, row := range content {
		// for _, cell := range row{
			x := string(row)
			sysCallPrintStr(x)
		// }
	}

}

func bufPrintContent(content [][] rune) {

	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	for _, row := range content {
		// for _, cell := range row{
			fmt.Fprint(buf, string(row))
		// }
	}
}

func bufPrintContentString(content []string) {

	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	for _, row := range content {
		// for _, cell := range row{
			fmt.Fprint(buf, row)
		// }
	}
}


func sysCallPrintContentString(content []string) {
	var base = 16
	var size = 16
	var data []uint16
	for _, row := range content {
		x, _ := strconv.ParseUint(row, base, size)
		data = append(data, uint16(x))
	}

	nw := uint32(len(data))

	syscall.WriteConsole(syscall.Stdout, &data[0], nw, &nw, nil)

	// length := len(data)
 //    for length > 0 {
	//     written, err := syscall.Write(2, data)
	//     if err != nil { // here
	//         panic(err)
	//     }
	//     length -= written
	//     data = data[written:]
	// }

}


func GenScreenContent(char rune, x, y int) [][] rune{
	content := [][] rune{}
	for a := 0; a < y; a++{
		row := [] rune{}
		for b := 0; b < x-1; b++{
			if a == 0{
				row = append(row, '=')
			} else{
				row = append(row, char)
			}
		}
		row = append(row, '|')

		content = append(content, row)
	} 
	return content
}

func GenScreenContentString(char string, x, y int, wallHeight float64) [] string{
	content := [] string{}
	wallStart := (float64(y) / 2) - (wallHeight / 2)
	wallend   := (float64(y) / 2) + (wallHeight / 2)
	for a := 0; a < y; a++{
		var sb strings.Builder
		for b := 0; b < x-1; b++{
			if a == 0{
				sb.WriteString("=")
			} else if a == y -1{
				sb.WriteString("=")
			} else {
				if float64(a) < wallStart{
					sb.WriteString(" ")
				} else if float64(a) < wallend{
					sb.WriteString(char)
				}else{
					sb.WriteString(" ")
				}

			}
		}
		sb.WriteString("|")
		content = append(content, sb.String())
	} 
	return content
}


func main() {

	// fmt
	t1 := time.Now()

	// testAmount := 10608

	// fmtPrint(testAmount)
	// fmtTime := fmt.Sprintf("=======fmt took %v========\n", time.Now().Sub(t1))

	// t1 = time.Now()
	// bufPrint(testAmount)
	// bufTime := fmt.Sprintf("=======bufPrint took %v========\n", time.Now().Sub(t1))

	// t1 = time.Now()
	// bufPrint(testAmount)
	// bufStackPrintTime := fmt.Sprintf("=======bufStackPrint took %v========\n", time.Now().Sub(t1))

	// t1 = time.Now()
	// sysCallPrint(testAmount)
	// sysCallPrintTime := fmt.Sprintf("=======sysCallPrint took %v========\n", time.Now().Sub(t1))

	// t1 = time.Now()
	// osPrint(testAmount)
	// osPrintTime := fmt.Sprintf("=======osPrint took %v========\n", time.Now().Sub(t1))

	// t1 = time.Now()
	// osStackPrint(testAmount)
	// osStackPrintTime := fmt.Sprintf("=======osStackPrint took %v========\n", time.Now().Sub(t1))


	//			powershell
	// =======fmt took 3.5557721s========

	// =======bufPrint took 801.3875ms========

	// =======bufStackPrint took 797.5128ms========

	// =======sysCallPrint took 64.0004ms========

	// =======osPrint took 3.4981111s========

	// =======osStackPrint took 913.5978ms========


	// cmd fmt much slower but others are faster
	// =======fmt took 4.5672253s========

	// =======bufPrint took 460.1196ms========

	// =======bufStackPrint took 486.3681ms========

	// =======sysCallPrint took 63.8008ms========

	// =======osPrint took 4.4705161s========

	// =======osStackPrint took 507.1621ms========


	// fmt.Println("======================================")
	// fmt.Println(fmtTime)
	// fmt.Println(bufTime)
	// fmt.Println(bufStackPrintTime)
	// fmt.Println(sysCallPrintTime)
	// fmt.Println(osPrintTime)
	// fmt.Println(osStackPrintTime)
	// fmt.Println("======================================")


	ScreenTest()
	t1 = time.Now()

	content := GenScreenContent('X', 211, 50)
	contentGenTime := fmt.Sprintf("=======contentGen took %v========\n", time.Now().Sub(t1))

	fmt.Println("======================================")
	t1 = time.Now()

	strContent := GenScreenContentString("@", 211, 50, 50)
	strContentGenTime := fmt.Sprintf("=======strContentGen took %v========\n", time.Now().Sub(t1))

	fmt.Println("======================================")
	t1 = time.Now()


	bufPrintContent(content)
	bufPrintContentTime := fmt.Sprintf("=======bufPrintContent took %v========\n", time.Now().Sub(t1))
	t1 = time.Now()
	fmt.Println("======================================")

	sysCallPrintContent(content)
	sysCallPrintContentTime := fmt.Sprintf("=======sysCallPrintContent took %v========\n", time.Now().Sub(t1))
	t1 = time.Now()
	fmt.Println("======================================")

	bufPrintContentString(strContent)
	bufPrintStringContentTime := fmt.Sprintf("=======bufPrintContentString took %v========\n", time.Now().Sub(t1))
	t1 = time.Now()
	fmt.Println("======================================")

	sysCallPrintContentString(strContent)
	sysCallPrintStringContentTime := fmt.Sprintf("=======sysCallPrintContentString took %v========\n", time.Now().Sub(t1))
	t1 = time.Now()
	fmt.Println("======================================")


	fmt.Println("======================================")
	fmt.Println(contentGenTime)
	fmt.Println(bufPrintContentTime)
	fmt.Println(sysCallPrintContentTime)
	fmt.Println("======================================")
	fmt.Println(strContentGenTime)
	fmt.Println(bufPrintStringContentTime)
	fmt.Println(sysCallPrintStringContentTime)
	fmt.Println("======================================")


	ClearScreen()
	//Remove cursor
	fmt.Print("\033[?25l")
	t1 = time.Now()
	counter := 0
	chr := "@"
	for {
		fmt.Printf("\033[%d;%dH", 0, 0)
		counter++

		// if counter % 200 > 100{
		// 	chr = "@"
		// } else {
		// 	chr = "#"
		// }
		bufPrintContentString(GenScreenContentString(chr, 211, 50, float64(counter % 50)))

		if time.Now().After(t1.Add(time.Second*10)){
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("======================================")
	fmt.Println("frames drawn ", counter)
	fmt.Println("This gives an FPS of ", counter/10)
	fmt.Println("======================================")
	//Re show cursor
	fmt.Print("\033[?25h")
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func ScreenTest() {
	for x:=0; x < 230; x++{
 		// power shell terminal
		// X 156
		// cmd
		// X 221
		if x % 10 == 0{
			fmt.Printf("\n=============%v===========\n", x)
		}
		fmt.Println(strings.Repeat("|", x))
	}
	fmt.Println("\n============================")
	// power shell terminal
	// X 39
	// cmd
	// X 50
	for y:= 0; y < 200; y++{
		fmt.Println(y)
	}

}


func CheckInput(keyPress chan <- rune){
	_, err := keyboard.GetKeys(10)
	if err != nil{
		log.Println(err)
	}	



}