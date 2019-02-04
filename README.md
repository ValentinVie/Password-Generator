# spwgen443

Password generation software in Go language.

In this assignment you will be building a password generator program. This program will generate a random password. The password will be generated either randomly or via a pattern specified on the command line. Follow the following steps to make this program.

Create your virtual machine and install the go language according to the instructions on the course website. Look for the file `getting-started-with-go-11-17-v1.pdf`. Preview the document under the folders directory.
Walk through the online tutorial for the Go language. Reference the file `go-language-overview-11-17-v1.pdf`. Preview the document also located on the course website which contains lots of information about the language semantics.
Login to your virtual machine. From your virtual machine, download the starter source code provided for this assignment. To do this, create the directories, download the starter file from canvas, and rename it. The file is named spwgen443-starter.go.

```
% cd ~/go/src
% mkdir -p cmpsc443/spwgen443
% cd cmpsc443/spwgen443
% cp ~/Downloads/spwgen443-starter.go
% mv spwgen443-starter.go spwgen443.go
```
You also need to install a few things to get this assignment working including the dictionary of words and a Go library. Uses these commands to do the installs.
% sudo apt-get install --reinstall wamerican
% go get github.com/pborman/getopt
You are to create a password generator program that will create a random password. The specification for the program is as follows:
```
USAGE : spwgen443 [-h] [-w] [-l ] [-p ]

where:

   -h : help (this menu)
   -w : web flag (no symbol characters, e.g., no &*...)
   -l : length of password (in characters)
   -p : pattern (set of symbols defining password)

        A pattern consists of a string of characters "xxxxx",
        where the x pattern characters include:

          d - digit
          c - upper or lower case character
          l - lower case character
          u - upper case character
          w - random word from /usr/share/dict/words (or /usr/dict/words)
              note that w# will identify a word of length #, if possible
          s - special character in ~!@#$%^&*()-_=+{}[]:;/?<>,.|\

        Note: the pattern overrides other flags, e.g., -w
```
Program notes:
The program should print out the password to standard output.

- If no length is given, 16 characters should be assumed.
- If no pattern is given, then each character should include a uniformly random single upper or lower case character with Pr(0.33), digit with Pr(0.33), or a special character with Pr(0.33) (as specified in the program help).
- If a pattern is given, the length should be ignored.
The words identified in the pattern (of length n) should be randomly selected from the dictionary installed on the UNIX system. For example, if part of the pattern is "w4", the word should be randomly selected from all 4 letter words in the dictionary. On Ubuntu the dictionary file is located at:
`/usr/share/dict/words`

The program should disregard all words in the dictionary that include any non-alphabet characters, e.g., spaces, quotes, umlauts, etc.
The maximum length of a password is 64 characters.

Be sure to test your program thoroughly using a large number of combinations and patterns. Here are some good tests to run:
```
./spwgen
./spwgen -h
./spwgen -w
./spwgen -l 8
./spwgen -p w2dw10
./spwgen -p luddcws
```