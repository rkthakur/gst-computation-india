package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"os"
	"strconv"
  )

// Jurisdictions struct which contains
// an array of Jurisdictions
type Jurisdictions struct {
    Jurisdictions []Jurisdiction `json:"jurisdictions"`
}

type Jurisdiction struct {
    JurisdictionId   string `json:"jurisdiction-id"`
    Taxclasses []Taxclass `json:"tax-classes"`
}

// Taxclasses struct which contains
// an array of Taxclasses
type Taxclasses struct {
	Taxclasses []Taxclass `json:"tax-classes"`
}

type Taxclass struct {
	TaxClassId string `json:"tax-class-id"`
	Slabs []Slab `json:"slabs"`
}

// Slab struct which contains
// an array of Slab
type Slab struct {
	SlabId string `json:"slab-id"`
	SlabRules SlabRules `json:"slab-rules"`
}

type SlabRules struct{
	MinSalesAmount float32 `json:"minimum-sales-amount"`
	MaxSalesAmount float32 `json:"maximum-sales-amount"`
	CgstRate float32 `json:"cgst-rate"`
	SgstRate float32 `json:"sgst-utgst-rate"`
	IgstRate float32 `json:"igst-rate"`
	CompensationCess string `json:"compensation-cess"`
}

func getTaxQuote(jurisdictionId string, taxClassId string, salesAmount float32){
	fmt.Printf("%s - %s - %v\n",jurisdictionId,taxClassId,salesAmount)
 // Open our jsonFile
 jsonFile, err := os.Open("gst-tax-slab.json")
 // if we os.Open returns an error then handle it
 if err != nil {
	 fmt.Println(err)
 }

 // defer the closing of our jsonFile so that we can parse it later on
 defer jsonFile.Close()

 // read our opened xmlFile as a byte array.
 byteValue, _ := ioutil.ReadAll(jsonFile)

 // we initialize our Jurisdictions array
 var jurisdictions Jurisdictions

 // we unmarshal our byteArray which contains our
 // jsonFile's content into 'Jurisdictions' which we defined above
 json.Unmarshal(byteValue, &jurisdictions)

 // we iterate through every user within our users array and
 // print out the user Type, their name, and their facebook url
 // as just an example
 fmt.Println("jurisdiction, Tax Class, Slab, minimum-sales-amount, maximum-sales-amount,cgst-rate,sgst-utgst-rate,igst-rate,compensation-cess")
 for i := 0; i < len(jurisdictions.Jurisdictions); i++ {
	 if jurisdictions.Jurisdictions[i].JurisdictionId == jurisdictionId {
	 for ii := 0; ii < len(jurisdictions.Jurisdictions[i].Taxclasses); ii++ {
		 if jurisdictions.Jurisdictions[i].Taxclasses[ii].TaxClassId == taxClassId {
		 for iii := 0; iii < len(jurisdictions.Jurisdictions[i].Taxclasses[ii].Slabs); iii++ {
				 fmt.Printf("%v - %v - %v -%v\n",jurisdictions.Jurisdictions[i].JurisdictionId, jurisdictions.Jurisdictions[i].Taxclasses[ii].TaxClassId, jurisdictions.Jurisdictions[i].Taxclasses[ii].Slabs[iii].SlabId, jurisdictions.Jurisdictions[i].Taxclasses[ii].Slabs[iii].SlabRules.MinSalesAmount)
				 slab := jurisdictions.Jurisdictions[i].Taxclasses[ii].Slabs[iii]
				 if slab.SlabRules.MinSalesAmount <= salesAmount && slab.SlabRules.MaxSalesAmount >= salesAmount {
					 taxQuote := salesAmount * slab.SlabRules.IgstRate
					 fmt.Printf("Computed GST amount is - %v @ %v of %v salesAmount\n", taxQuote, slab.SlabRules.IgstRate, salesAmount)
					 break
			 }
		 } 
		}
	 }
	}
 }
}

func main() {
	var SAmt float32 = 0.0
	if sAmt, err := strconv.ParseFloat(os.Args[3], 32); err == nil {
		SAmt = float32(sAmt)
	}
	fmt.Printf("%s - %s - %v\n",os.Args[1], os.Args[2], SAmt)
	getTaxQuote(os.Args[1], os.Args[2], SAmt)
}