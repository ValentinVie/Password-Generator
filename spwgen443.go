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
	"bufio" //read the dictionary
	"unicode/utf8" // read the lenght of a string with stresses
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

var dictPath string = "/usr/share/dict/words"
var nbEntries int = 0 // number of entries in the dict
// You may want to create more global variables

func lineCounter(filePath string) int{ // count the number of entries in the dictionary
	dico, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Couldn't find the file... Abort")
  	os.Exit(-1)
  }
  defer dico.Close() // close the dictionary once the function returns
	count := 0
	scanner := bufio.NewScanner(dico)

	for scanner.Scan(){
		count++
	}
	return count
}

func validWord(word string) bool{ // the word contains only alpha caracters
	for i := range word {
		if word[i] < 'A' || word[i] > 'z' {
			return false
		} else if word[i] > 'Z' && word[i] < 'a' {
			return false
		}
	}
	return true
}

func findWordFromDictionary(lenght int) string{
	if lenght <= 0 || lenght >= 15{
		return "ERROR"
	}
	// find a random word of lenght "lenght" in the dictionary file
	if nbEntries == 0 {
		nbEntries = lineCounter(dictPath)
		fmt.Printf("There are %v entries in the dictionary.\n", nbEntries)
	}

	dico, err := os.Open(dictPath)
	if err != nil {
		fmt.Printf("Couldn't find the dictionary... Abort")
  	os.Exit(-1)
  }
  defer dico.Close() // close the dictionary once the function returns

  scanner := bufio.NewScanner(dico)
	line_number := rand.Intn(nbEntries) // offset
	i := 0
	// we will go accross line_number lines before selecting a word of the correct length
	for i < line_number{ // offset
		scanner.Scan()
		i++
	}

	for scanner.Scan(){ // we take the next word that fits
		//fmt.Println(scanner.Text(), len(scanner.Text()))
		if utf8.RuneCountInString(scanner.Text()) == lenght && validWord(scanner.Text()){
			return scanner.Text()
		}
		if i == nbEntries - 1 {//we reached the end of the file we start over
			dico.Seek(0, 0)
			i = 0
		}
		i++
	}
	return "ERROR"
}

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
		i := 0
		for i < len(pattern){ // for every char in pattern
			if pattern[i] == byte('d'){ // digits
				pwd = pwd + strconv.Itoa(rand.Intn(10))
			} else if pattern[i] == byte('c'){ // char
				pwd = pwd + string(caracters[rand.Intn(len_caracters)])
			} else if pattern[i] == byte('l'){ // char lower case
					pwd = pwd + string(lower_caracters[rand.Intn(len_lower_caracters)])
			} else if pattern[i] == byte('u'){ // char upper case
					pwd = pwd + string(upper_caracters[rand.Intn(len_upper_caracters)])
			} else if pattern[i] == byte('w'){ // word where the length is specified
				var word_length1, word_length2 int
				var err1, err2 error
				if i + 1 < len(pattern){ // check if length is specified
					word_length1, err1 = strconv.Atoi(string(pattern[i+1]))
				}
				if i + 2 < len(pattern){ // check if length is specified
					word_length2, err2 = strconv.Atoi(string(pattern[i+1])+string(pattern[i+2]))
				}

				if err2 == nil && i + 2 < len(pattern){ // wXX where XX is a number
					if word_length2 > 0 && word_length2 < 15{
						pwd += findWordFromDictionary(word_length2)
						i += 2
					} else{ // lenght not correct
						fmt.Printf("The lenght of the word specified is too long or too short.\n")
						os.Exit(-1)
					}
				}else if err1 == nil && i + 1 < len(pattern){ // wX where X is a number
					if word_length1 > 0 {
						pwd += findWordFromDictionary(word_length1)
						i++
					} else{ // length <= 0...
						fmt.Printf("The lenght specified is too short.\n")
						os.Exit(-1)
					}
				} else { // w no length specified, we randomly select one.
					var x = rand.Intn(14) + 1 //avoid word of length 0
					fmt.Println(x)
					pwd += findWordFromDictionary(x)
				}
			} else if pattern[i] == byte('s'){ // special caracter
				pwd = pwd + string(specials[rand.Intn(len_specials)])
			} else { //Abort, wrong arguments in the pattern
				fmt.Printf("Wrong arguments in the pattern.\n")
				fmt.Printf(patternval+"\n")
				os.Exit(-1)
			}
			i++
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
	helpflag := getopt.Bool('h', "help (this menu)")
	webflag := getopt.Bool('w', "web flag (no symbol characters, e.g., no &*...)")
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

	if(*helpflag){
		getopt.Usage()
	}
	// Now generate the password and print it out
	pwd := generatePasword(plength, *pattern, *webflag)
	fmt.Printf("Generated password:  %s\n", pwd)

	// Return (no return code)
	return
}
