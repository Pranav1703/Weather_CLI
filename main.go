package main


import (
	"fmt"
	"bufio"
	"os"
	"weatherCli/internal/cmdParser"
	"github.com/fatih/color"
	//  "weatherCli/internal/weather"

)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("-> Enter Your Command\n")
	for {
		fmt.Printf("\n=> ")
		
		scanner.Scan()
		text := scanner.Text()
		
		input,err := cmdParser.ParseInput(text)
		if err != nil {
			color.HiRed("%s",err)
			continue
		}
		
		if len(input)==0 {
			continue
		}
		
		cmd := input[0]
		validCmds := cmdParser.GetCommands()
		

		if command,ok:=validCmds[cmd]; ok {
			if len(input) == 2{
				
				command.CmdFuncMain(input[1])
			
			}else{
				
				command.CmdFunc()
				
			}

		}
	
	}

}