// package eml contains structures that represent the EML file format.
package eml

import "encoding/xml"

type EML struct {
	XMLName        xml.Name `xml:"eml"`
	Lang           string   `xml:"xml:lang,attr"`
	SchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	Dataset        Dataset  `xml:"dataset"`
}

type Dataset struct {
	ID               string           `xml:"id,attr"`
	Title            string           `xml:"title"`
	Creator          Creator          `xml:"creator"`
	MetadataProvider MetadataProvider `xml:"metadataProvider"`
	PubDate          string           `xml:"pubDate"`
	Abstract         Abstract         `xml:"abstract"`
}

type Creator struct {
	ID                    string         `xml:"id,attr"`
	Scope                 string         `xml:"scope,attr"`
	IndividualName        IndividualName `xml:"individualName"`
	ElectronicMailAddress string         `xml:"electronicMailAddress"`
}

type MetadataProvider struct {
	IndividualName        IndividualName `xml:"individualName"`
	ElectronicMailAddress string         `xml:"electronicMailAddress"`
}

type IndividualName struct {
	GivenName string `xml:"givenName"`
	SurName   string `xml:"surName"`
}

type Abstract struct {
	Para string `xml:"para"`
}
