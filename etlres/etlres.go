package etlres

import (
	"encoding/csv"
	"fmt"
	"log"
	//"io"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type ResourceItem struct {
	Author       string
	Identifer    string
	Name         string
	Keywords     []string
	KeywordsGTT1 []string
	KeywordsGTT2 []string
	KeywordsDT1  []string
	KeywordsDT2H []string
	KeywordsDT2S []string
	URL          string
	DataType     string
	ResourceType string
	Descriptions string
	Services     string
	Contact      string
	Email        string
	Comments     string
}

var antemp = template.Must(template.New("anbody").Parse(anbody))
var anheadertemp = template.Must(template.New("anheader").Parse(anheader))

//var tsvtemp = template.Must(template.New("tsvbody").Parse(tsvbody))

// todo this is where to do the data cleaning of the terms..   not later....
func (t *ResourceItem) Parse(in []string) {
	t.Author = in[0]
	t.Identifer = in[1]
	t.Name = in[2]
	t.Keywords = strings.Split(in[3], ",") //array of keywords
	t.KeywordsGTT1 = strings.Split(in[4], ",")
	t.KeywordsGTT2 = strings.Split(in[5], ",")
	t.KeywordsDT1 = strings.Split(in[6], ",")
	t.KeywordsDT2H = strings.Split(in[7], ",")
	t.KeywordsDT2S = strings.Split(in[8], ",")
	t.URL = in[9]
	t.DataType = in[10]
	t.ResourceType = in[11]
	t.Descriptions = in[12]
	t.Services = in[13]
	t.Contact = in[14]
	t.Email = in[15]
	t.Comments = in[16]
}

// todo  make a function to make all string a valid CDATA string type.

func BuildRDFFiles() {
	csvFile, err := os.Open("./Morgue/test2.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(csvFile)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	resources := make([]ResourceItem, len(lines))

	r, _ := regexp.Compile(`\W`) // todo modify this to ignore space and comma..  currently a hack to replace bad chars with space

	for i, line := range lines {
		var ritem ResourceItem
		ritem.Parse(line)
		log.Printf("%v", i)
		resources[i] = ritem
	}

	ntfile, err := os.Create("./output/resource.nt")
	if err != nil {
		panic(err)
	}
	defer ntfile.Close()

	// write triples..  use rapper to conver to turtle or RDF/XML
	//TODO  make this a template too for turtle.. then rapper to ntriples if needed (easier and more compact )
	for _, ritem := range resources {
		// ntfile.WriteString(fmt.Sprintf("<%v> <%v> <%v> . \n", ritem.Identifer, "rdfs:type", "schema:WebSite"))
		ntfile.WriteString(fmt.Sprintf("<%v> <%v> <%v> . \n", ritem.Identifer, "rdfs:type", "schema:WebSite"))

		if ritem.Author != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "foaf:name", ritem.Author))
		}
		if ritem.Name != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dc:resource", ritem.Name))
		}
		// The whole keyword section has an error in that there are dups
		// todo  read all keywords into an array (uniquely) and then write them all out at once.
		// todo phase2  (associate the provenance of how these were arrived at, scrapped, human, etc)
		if len(ritem.Keywords) > 0 {
			for _, value := range ritem.Keywords {
				if value != "" {
					ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:keyword", r.ReplaceAllString(strings.TrimSpace(value), " ")))
				}
			}
		}
		if len(ritem.KeywordsGTT1) > 0 {
			for _, value := range ritem.KeywordsGTT1 {
				if value != "" {
					ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:keyword", r.ReplaceAllString(strings.TrimSpace(value), " ")))
				}
			}
		}
		if len(ritem.KeywordsGTT2) > 0 {
			for _, value := range ritem.KeywordsGTT2 {
				if value != "" {
					ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:keyword", r.ReplaceAllString(strings.TrimSpace(value), " ")))
				}
			}
		}
		if len(ritem.KeywordsDT1) > 0 {
			for _, value := range ritem.KeywordsDT1 {
				if value != "" {
					ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:keyword", r.ReplaceAllString(strings.TrimSpace(value), " ")))
				}
			}
		}
		if len(ritem.KeywordsDT2H) > 0 {
			for _, value := range ritem.KeywordsDT2H {
				if value != "" {
					ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:keyword", r.ReplaceAllString(strings.TrimSpace(value), " ")))
				}
			}
		}
		if len(ritem.KeywordsDT2S) > 0 {
			for _, value := range ritem.KeywordsDT2S {
				if value != "" {
					ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:keyword", r.ReplaceAllString(strings.TrimSpace(value), " ")))
				}
			}
		}
		if ritem.URL != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "schema:URL", ritem.URL))
		}
		if ritem.DataType != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:datatype", ritem.DataType))
		}
		if ritem.ResourceType != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dcat:resourcetype", ritem.ResourceType))
		}
		if ritem.Descriptions != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dc:description", r.ReplaceAllString(strings.TrimSpace(ritem.Descriptions), " ")))
		}
		if ritem.Services != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "cinergi:services", r.ReplaceAllString(strings.TrimSpace(ritem.Services), " ")))
		}
		if ritem.Contact != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "dc:contact", r.ReplaceAllString(strings.TrimSpace(ritem.Contact), " ")))
		}
		if ritem.Email != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "foaf:email", ritem.Email))
		}
		if ritem.Comments != "" {
			ntfile.WriteString(fmt.Sprintf("<%v> <%v> \"%v\" . \n", ritem.Identifer, "rdfs:commnents", ritem.Comments))
		}
	}

	// write annotation file..  use rapper to conver to turtle or RDF/XML
	adata, err := os.Create("./output/annotations.xml")
	if err != nil {
		panic(err)
	}
	defer adata.Close()

	// need to find all the URL's that will be marked up first..  to add to the
	// annotation file header.
	urlCount := 0
	for _, ritem := range resources {
		if strings.Contains(ritem.URL, "http://") {
			urlCount = urlCount + 1
		}
	}

	anheadertemp.Execute(adata, urlCount)

	for _, ritem := range resources {

		if strings.Contains(ritem.URL, ".html") || strings.Contains(ritem.URL, ".htm") || strings.Contains(ritem.URL, ".php") || strings.Contains(ritem.URL, ".aspx") || strings.Contains(ritem.URL, ".asp") || strings.Contains(ritem.URL, ".jsp") || strings.Contains(ritem.URL, ".pdf") {
			fmt.Printf("Have this : %v\n", ritem.URL)
			split := strings.Split(ritem.URL, "/")
			//fmt.Printf("Want to remove : %v\n", split[len(split)-1])
			newURL := strings.Replace(ritem.URL, "/"+split[len(split)-1], "", -1)
			fmt.Printf("Want this : %v\n\n", newURL)
			// Split the string on / and remove the last index
			ritem.URL = newURL
		}

		// todo..  what to do with the FTP resources....
		if len(ritem.URL) > 0 && !strings.Contains(ritem.URL, "ftp://") {
			antemp.Execute(adata, ritem)
		}
	}

	adata.WriteString(anfooter)

	atsv, err := os.Create("./output/annotations.tsv")
	if err != nil {
		panic(err)
	}
	defer atsv.Close()

	// need to find all the URL's that will be marked up first..  to add to the
	// annotation file header.
	urlCount = 0
	for _, ritem := range resources {
		if strings.Contains(ritem.URL, "http://") {
			urlCount = urlCount + 1
		}
	}

	// write the TSV format here...
	//anheadertemp.Execute(atsv, urlCount)
	atsv.WriteString(fmt.Sprintf("%v\t%v\t%v  \n", "URL", "Label", "Label"))

	for _, ritem := range resources {

		if strings.Contains(ritem.URL, ".html") || strings.Contains(ritem.URL, ".htm") || strings.Contains(ritem.URL, ".php") || strings.Contains(ritem.URL, ".aspx") || strings.Contains(ritem.URL, ".asp") || strings.Contains(ritem.URL, ".jsp") || strings.Contains(ritem.URL, ".pdf") {
			fmt.Printf("Have this : %v\n", ritem.URL)
			split := strings.Split(ritem.URL, "/")
			//fmt.Printf("Want to remove : %v\n", split[len(split)-1])
			newURL := strings.Replace(ritem.URL, "/"+split[len(split)-1], "", -1)
			// todo replace the http:// with nothing...
			fmt.Printf("Want this : %v\n\n", newURL)
			// Split the string on / and remove the last index
			ritem.URL = newURL
		}

		// todo..  what to do with the FTP resources....
		if len(ritem.URL) > 0 && !strings.Contains(ritem.URL, "ftp://") {
			atsv.WriteString(fmt.Sprintf("%v/*\t%v\t\n", ritem.URL, "_cse_d6ieoekzj1u"))
			//tsvtemp.Execute(atsv, ritem)
		}
	}

}

//const tsvbody = `{{.URL}}/*	_cse_d6ieoekzj1u`

// odata template here
const anheader = `<?xml version="1.0" encoding="utf-8"?> 
<Annotations  start="0" num="{{.}}" total="{{.}}">`

// TODO..  look at how this URL is formed with the /* on the end
// ..  need to make a functional test for this?
const anbody = `   
   <Annotation about="{{.URL}}/*">
     <Label name="_cse_d6ieoekzj1u"/>
     <Comment>{{.Comments}}</Comment>
   </Annotation>`

const anfooter = `
 </Annotations>
`
