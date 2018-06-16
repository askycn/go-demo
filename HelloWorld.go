// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "./demo"

func main() {
	print()
}

// add some comments
func print() {
	fmt.Println("Hello, 世界")
	//msg := demo.I18NMessage{"EN", "Hello World!"}
	msg := demo.I18NMessage{}
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("Handling unknown language")
			fmt.Println(fmt.Sprintf("%s : %s", "Recovering from", err))
			msg.Clear()
			fmt.Println(msg.GetLen())
		}
	}()

	msg.AddMsg("EN", "Hello World!")
	msg.Say("EN")
	fmt.Println(msg.GetLen())

	msg.Say("CH")

}
