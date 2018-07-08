package main

import "fmt"
import "os"
import "strings"

import (
	flag "github.com/ogier/pflag"
)

// flags 
var (
	user string
)

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	// we can move 
	if flag.NFlag() == 0 {
		printUsage()
	}
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s) %s \n", users)

	// return user data 
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`username:   `, result.Login)
		fmt.Println(`Email:   `, result.Email)
		fmt.Println(`Bio:  `, result.Bio)
	}
}

func init() {
	// args var, double-dash, single-dash, default, description
	flag.StringVarP(&user, "user", "u", "", "Search User")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options")
	flag.PrintDefaults()
	os.Exit(1)
}