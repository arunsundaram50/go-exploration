package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// define the command
	myCmd := cobra.Command{
		Use:   "greet",
		Short: "",
		Long:  "",
	}

	// define the flags for the command
	greetingFlagSet := myCmd.Flags()
	greetingFlagSet.String("person", "", "The name of the person you want to greet")
	greetingFlagSet.Int("age", 0, "Age of the person")

	// parse the input
	greetingFlagSet.Parse(os.Args)

	// get the flag values
	person, err := greetingFlagSet.GetString("person")
	if err != nil || person == "" {
		fmt.Println("Please specify the person")
		return
	}
	age, err := greetingFlagSet.GetInt("age")
	if err != nil {
		fmt.Println("Please specify the age")
		return
	}

	// using the values
	fmt.Printf("Hello %d year old %s\n", age, person)
}
