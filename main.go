package main

import (
	"fmt"
	"os"
	"syncMyConf/src"

	"github.com/devfacet/gocmd/v3"
)

func main() {
	// fmt.Println("Lets go")

	flags := struct {
		Help    bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version bool `short:"v" long:"version" description:"Display version"`

		Push struct {
		} `command:"push" description:"push the repo files to their destination"`

		Pull struct {
		} `command:"pull" description:"pull the files from the destination into the repo"`

		Pwd struct {
		} `command:"pwd" description:"print pwd"`

		// Echo struct {
		// 	Settings bool `settings:"true" allow-unknown-arg:"true"`
		// } `command:"echo" description:"Print arguments"`
		// // Math struct {
		// 	Sqrt struct {
		// 		Number float64 `short:"n" long:"number" required:"true" description:"Number"`
		// 	} `command:"sqrt" description:"Calculate square root"`
		// 	Pow struct {
		// 		Base     float64 `short:"b" long:"base" required:"true" description:"Base"`
		// 		Exponent float64 `short:"e" long:"exponent" required:"true" description:"Exponent"`
		// 	} `command:"pow" description:"Calculate base exponential"`
		// } `command:"math" description:"Math functions" nonempty:"true"`
	}{}
	// Echo command
	// gocmd.HandleFlag("Echo", func(cmd *gocmd.Cmd, args []string) error {
	// 	fmt.Printf("%s\n", strings.Join(cmd.FlagArgs("Echo")[1:], " "))
	// 	return nil
	// })

	// // Math commands
	// gocmd.HandleFlag("Math.Sqrt", func(cmd *gocmd.Cmd, args []string) error {
	// 	fmt.Println(math.Sqrt(flags.Math.Sqrt.Number))
	// 	return nil
	// })
	// gocmd.HandleFlag("Math.Pow", func(cmd *gocmd.Cmd, args []string) error {
	// 	fmt.Println(math.Pow(flags.Math.Pow.Base, flags.Math.Pow.Exponent))
	// 	return nil
	// })

	gocmd.HandleFlag("Push", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println("running push")
		src.Push()
		return nil
	})

	gocmd.HandleFlag("Pull", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println("running pull")
		return nil
	})

	gocmd.HandleFlag("Pwd", func(cmd *gocmd.Cmd, args []string) error {
		pwd, _ := os.Getwd()
		fmt.Println("pwd:", pwd)
		return nil
	})

	// Init the app
	gocmd.New(gocmd.Options{
		Name:        "basic",
		Description: "A basic app",
		Version:     fmt.Sprintf("%s", "0.1.0"),
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
