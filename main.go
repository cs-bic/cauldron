package main
import (
	"cauldron/witchery"
	"io/ioutil"
	"os"
)
func main() {
	if len(os.Args) != 2 {
		panic("main: Not enough or too many parameters were provided.")
	}
	code, issue := ioutil.ReadFile(os.Args[1])
	if issue != nil {
		panic("main: There was an issue in reading 'script.witchery' from disk.")
	}
	witchery.Initialize(string(code))
}
