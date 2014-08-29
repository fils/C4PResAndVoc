package main

import (
	etlres "oceandrilling.org/C4PResAndVoc/etlres"
	etlvoc "oceandrilling.org/C4PResAndVoc/etlvoc"
)

func main() {
	etlres.BuildRDFFiles()
	etlvoc.BuildVocFiles()
}
