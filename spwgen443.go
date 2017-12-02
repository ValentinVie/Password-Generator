////////////////////////////////////////////////////////////////////////////////
//
//  File           : spwgen443.go
//  Description    : This is the implementaiton file for the spwgen443 password
//                   generator program.  See assignment details.
//
//  Collaborators  : Valentin Vie
//  Last Modified  : Valentin Vie
//

// Package statement
package main

// Imports
import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
	"github.com/pborman/getopt"
	// There will likely be several mode APIs you need
)

// Global data
var patternval string = `pattern (set of symbols defining password)

        A pattern consists of a string of characters "xxxxx",
        where the x pattern characters include:

          d - digit
          c - upper or lower case character
          l - lower case character
          u - upper case character
          w - random word from /usr/share/dict/words (or /usr/dict/words)
              note that w# will identify a word of length #, if possible
          s - special character in ~!@#$%^&*()-_=+{}[]:;/?<>,.|\

        Note: the pattern overrides other flags, e.g., -w`

// You may want to create more global variables

//
// Functions

// Up to you to decide which functions you want to add

////////////////////////////////////////////////////////////////////////////////
//
// Function     : generatePasword
// Description  : This is the function to generate the password.
//
// Inputs       : length - length of password
//                pattern - pattern of the file ("" if no pattern)
//                webflag - is this a web password?
// Outputs      : 0 if successful test, -1 if failure

func generatePasword(length int8, pattern string, webflag bool) string {
	pwd := "" // Start with nothing and add code
	var options string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*()-_=+{}[]:;/?<>,.|\`
	len_opt := len(options)

	var options_webflag string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`
	len_opt_webflag := len(options_webflag)

	var caracters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	len_caracters := len(caracters)

	var lower_caracters string = "abcdefghijklmnopqrstuvwxyz"
	len_lower_caracters := len(lower_caracters)

	var upper_caracters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	len_upper_caracters := len(upper_caracters)

	var specials string = `~!@#$%^&*()-_=+{}[]:;/?<>,.|\`
	len_specials := len(specials)

	if len(pattern) == 0 { // No pattern
		if webflag{
			for i := int8(0); i<length; i++{
				pwd = pwd + string(options_webflag[rand.Intn(len_opt_webflag)])
			}
		} else {
			for i := int8(0); i<length; i++{
				pwd = pwd + string(options[rand.Intn(len_opt)])
			}
		}
	} else { // There is a pattern we override others flags
		for i := 0; i < len(pattern); i++{ // for every char in pattern
			if pattern[i] == byte('d'){
				pwd = pwd + strconv.Itoa(rand.Intn(10))
			} else if pattern[i] == byte('c'){
				pwd = pwd + string(caracters[rand.Intn(len_caracters)])
			} else if pattern[i] == byte('l'){
					pwd = pwd + string(lower_caracters[rand.Intn(len_lower_caracters)])
			} else if pattern[i] == byte('u'){
					pwd = pwd + string(upper_caracters[rand.Intn(len_upper_caracters)])
			} else if pattern[i] == byte('w'){
				// TO BE DONE
			} else if pattern[i] == byte('s'){
				pwd = pwd + string(specials[rand.Intn(len_specials)])
			} else {
				fmt.Printf("Wrong arguments in the pattern. They have been ignored.\n")
			}
		}
	}



	// Now return the password
	return pwd
}

////////////////////////////////////////////////////////////////////////////////
//
// Function     : main
// Description  : The main function for the password generator program
//
// Inputs       : none
// Outputs      : 0 if successful test, -1 if failure

func main() {

	// Setup options for the program content
	rand.Seed(time.Now().UTC().UnixNano())
	helpflag := getopt.Bool('h', "", "help (this menu)")
	webflag := getopt.Bool('w', "", "web flag (no symbol characters, e.g., no &*...)")
	length := getopt.String('l', "", "length of password (in characters)")
	pattern := getopt.String('p', "", patternval)

	// Now parse the command line arguments
	err := getopt.Getopt(nil)
	if err != nil {
		// Handle error
		fmt.Fprintln(os.Stderr, err)
		getopt.Usage()
		os.Exit(-1)
	}

	// Get the flags
	fmt.Printf("helpflag [%t]\n", *helpflag)
	fmt.Printf("webflag [%t]\n", *webflag)
	fmt.Printf("length [%s]\n", *length)
	fmt.Printf("pattern [%s]\n", *pattern)
	// Normally, we we use getopt.Arg{#) to get the non-flag paramters

	// Safety check length parameter
	var plength int8 = 16
	if *length != "" {
		temp, err := strconv.Atoi(*length)
		if err != nil {
			fmt.Printf("Bad length passed in [%s]\n", *length)
			fmt.Fprintln(os.Stderr, err)
			getopt.Usage()
			os.Exit(-1)
		}
		if temp <= 0 || temp > 64 {
			fmt.Printf("Length passed to long or invalid, default value (16) used.\n")
			plength = 16
		} else {
			plength = int8(temp)
		}
	}


	// Now generate the password and print it out
	pwd := generatePasword(plength, *pattern, *webflag)
	fmt.Printf("Generated password:  %s\n", pwd)

	// Return (no return code)
	return
}
