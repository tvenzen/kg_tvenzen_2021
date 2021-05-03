package main

import (
    "fmt"
	"strings"
	"os"
)

func main() {
	
	//To take in all command line arguments and assign them to a variable
	nums := os.Args[1:]
	
	//Defines number of all arguments to be used in for-loop
	length := len(nums)

	for k := 0; k < length; k++{ 
		var number string = nums[k]
		
		//All if statments are necessary to catch all number cases.
		//If-else and Switch seem to only do the first case per string.
		if (strings.Contains(number,"0")){
			//strings.Replace replaces a number, "99" is to assure 
			//it replaces all instances of a specified number 

			number = strings.Replace(number,"0","Zero", 99)
		}
		if (strings.Contains(number,"1")){
			number = strings.Replace(number,"1","One", 99)
		}
		if (strings.Contains(number,"2")){
			number = strings.Replace(number,"2","Two", 99)
		}
		if (strings.Contains(number,"3")){
			number = strings.Replace(number,"3","Three", 99)
		}
		if (strings.Contains(number,"4")){
			number = strings.Replace(number,"4","Four", 99)
		}
		if (strings.Contains(number,"5")){
			number = strings.Replace(number,"5","Five", 99)
		}
		if (strings.Contains(number,"6")){
			number = strings.Replace(number,"6","Six", 99)
		}
		if (strings.Contains(number,"7")){
			number = strings.Replace(number,"7","Seven", 99)
		}
		if (strings.Contains(number,"8")){
			number = strings.Replace(number,"8","Eight", 99)
		}
		if (strings.Contains(number,"9")){
			number = strings.Replace(number,"9","Nine", 99)
		}
		//Prints the new converted string.
		fmt.Print(number, ",")
}  
}