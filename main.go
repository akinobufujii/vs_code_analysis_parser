package main

import (
	"fmt"
	"log"

	"github.com/akinobufujii/vs_code_analysis_parser/vscodeanalysis"
)

func main() {
	defectsData, err := vscodeanalysis.LoadDefectsData("vc.nativecodeanalysis.all.xml")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%v\n", defectsData)
}
