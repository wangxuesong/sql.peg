
package main

import (
	// "flag"
	// "fmt"
	// "io/ioutil"
	"log"
	"runtime"
	// "time"
    "sqlparser"
)

// var (
// 	inline = flag.Bool("inline", false, "parse rule inlining")
// 	_switch = flag.Bool("switch", false, "replace if-else if-else like blocks with switch blocks")
// 	syntax = flag.Bool("syntax", false, "print out the syntax tree")
// 	highlight = flag.Bool("highlight", false, "test the syntax highlighter")
// 	test = flag.Bool("test", false, "test the PEG parser performance")
// 	print = flag.Bool("print", false, "directly dump the syntax tree")
// )

func main() {
	runtime.GOMAXPROCS(2)
	// flag.Parse()

	// if flag.NArg() != 1 {
	//		flag.Usage()
	//		log.Fatalf("FILE: the peg file to compile")
	//	}
	// file := flag.Arg(0)

	buffer := " SELECT * FROM a" // ioutil.ReadFile(file)

	p := &sqlparser.Peg{Buffer: string(buffer)}
	p.Init()
	if err := p.Parse(); err != nil {
		log.Fatal(err)
	}

	// p.Execute()

}
