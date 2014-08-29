package etlvoc

import (
	"encoding/csv"
	"fmt"
	"log"
	//"io"
	"os"
	"strings"
)

// todo..   deal with comment lines
// todo..   parse the define strings for quotes

type TopResourceItem struct {
	Term     string
	Def      string
	Narrower []string
}

type SecondResourceItem struct {
	Term string
	Def  string
}

// todo this is where to do the data cleaning of the terms..   not later....
func (t *TopResourceItem) Parse(in []string) {
	t.Term = in[0]
	t.Def = in[1]
	t.Narrower = strings.Split(in[2], ",")
}

func (t *SecondResourceItem) Parse(in []string) {
	t.Term = in[0]
	t.Def = in[1]
}

// todo  make a function to make all string a valid CDATA string type.

func BuildVocFiles() {
	csvFile, err := os.Open("./dataSets/Glossary-Vocab-TopLevel.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(csvFile)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	resources := make([]TopResourceItem, len(lines))

	for i, line := range lines {
		var ritem TopResourceItem
		ritem.Parse(line)
		log.Printf("%v", i)
		resources[i] = ritem
	}

	ntfile, err := os.Create("./output/skos.n3")
	if err != nil {
		panic(err)
	}
	defer ntfile.Close()

	// write triples..  use rapper to conver to turtle or RDF/XML
	//TODO  make this a template too for turtle.. then rapper to ntriples if needed (easier and more compact )
	for _, ritem := range resources {
		// ntfile.WriteString(fmt.Sprintf("<%v> <%v> <%v> . \n", ritem.Identifer, "rdfs:type", "schema:WebSite"))
		ntfile.WriteString(fmt.Sprintf("c4p:%v rdf:type skos:Concept ; \n", ritem.Term))
		ntfile.WriteString(fmt.Sprintf("	skos:prefLabel   \"%v\"  ; \n", ritem.Term))
		for _, value := range ritem.Narrower {
			if value != "" {
				ntfile.WriteString(fmt.Sprintf("	skos:narrower   c4p:%v  ; \n", strings.TrimSpace(value)))
			}
		}
		ntfile.WriteString(fmt.Sprintf("	skos:definition   \"%v\"  . \n\n", ritem.Def))
	}

	// do second file for lower level terms
	csvFileS, err := os.Open("./dataSets/Glossary-Vocab-SecondLevel.csv")
	defer csvFileS.Close()
	if err != nil {
		panic(err)
	}
	readerS := csv.NewReader(csvFileS)

	linesS, err := readerS.ReadAll()
	if err != nil {
		log.Fatalf("error reading all linesS: %v", err)
	}

	resourcesS := make([]SecondResourceItem, len(linesS))

	for i, line := range linesS {
		var ritem SecondResourceItem
		ritem.Parse(line)
		log.Printf("%v", i)
		resourcesS[i] = ritem
	}

	for _, ritem := range resourcesS {
		// ntfile.WriteString(fmt.Sprintf("<%v> <%v> <%v> . \n", ritem.Identifer, "rdfs:type", "schema:WebSite"))
		ntfile.WriteString(fmt.Sprintf("c4p:%v rdf:type skos:Concept ; \n", ritem.Term))
		ntfile.WriteString(fmt.Sprintf("	skos:prefLabel   \"%v\"  ; \n", ritem.Term))
		ntfile.WriteString(fmt.Sprintf("	skos:definition   \"%v\"  . \n\n", ritem.Def))
	}

}
