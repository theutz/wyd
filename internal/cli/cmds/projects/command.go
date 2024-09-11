package projects

import "fmt"

func (cmd *ListCmd) Run() error {
	fmt.Println("list all projects")
	return nil
}
