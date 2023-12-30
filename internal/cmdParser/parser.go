package cmdParser

import (
	"errors"
	"fmt"
	"strings"
)

func ParseInput(s string) ([]string , error) { //parse input and check for number of words.return a error if there are more than one command given as a input or invalid arguments are provided
	text := strings.ToLower(s)
	words := strings.Fields(text)
	validCmds := GetCommands()
	
	command,ok:=validCmds[words[0]]

	if len(words) ==1 && ok{
		return words,nil
	}
	if !ok{
	
		return nil,errors.New("invalid command. type 'help' to see all commands")
	}

	
	if len(words) > 1 && !command.hasInput{
			
		return nil ,errors.New("error: more than one command provided. only one command is accepted")

	}
	
	if command.hasInput && len(words) > 2 || len(words) < 2 {   //checking for invalid args condition 
			
		return nil ,fmt.Errorf("error: invalid number of argument provided. '%s' -comamnd accepts only one argument",words[0])

	}

	
	
	return words , nil
	
}

// cmd arg1 arg 2  -
//cmd arg1         -
// cmd             -
// cmd inv inv inv - done