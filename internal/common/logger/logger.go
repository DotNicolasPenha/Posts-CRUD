package logger

import "fmt"

func Info(text string) {
	fmt.Printf("[info] %s \n", text)
}
func Error(err error) {
	fmt.Printf("[error] %s \n", err.Error())
}
func GroupError(errs []error) {
	fmt.Println("Errors:")
	for i, e := range errs {
		fmt.Printf("[%d] %s", i, e.Error())
	}
}
func Fatal(text string) {
	fmt.Printf("[fatal] %s \n", text)
}
