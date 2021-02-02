package defs

import "encoding/xml"

// Client is
type Client struct {
	XMLName xml.Name `xml:"client"`
	Host    string   `xml:"host,attr"`
}

// Clients is
type Clients struct {
	XMLName xml.Name `xml:"clients"`
	Clients []Client `xml:"client"`
}

// Server is
type Server struct {
	XMLName xml.Name `xml:"server"`
	Host    string   `xml:"host,attr"`
	Port    string   `xml:"port,attr"`
	Type    string   `xml:"type,attr"`
}

// Servers is
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Servers []Server `xml:"server"`
}

// Monitor is
type Monitor struct {
	XMLName xml.Name `xml:"monitor"`
	Host    string   `xml:"host,attr"`
	Type    string   `xml:"type,attr"`
}

// Monitoring is
type Monitoring struct {
	XMLName xml.Name `xml:"monitoring"`
	Monitor Monitor  `xml:"monitor"`
}

// Users is
type Users struct {
	XMLName     xml.Name `xml:"users"`
	ArrivalRate string   `xml:"arrivalrate,attr"`
	Unit        string   `xml:"unit,attr"`
}

// ArrivalPhase is
type ArrivalPhase struct {
	XMLName  xml.Name `xml:"arrivalphase"`
	Phase    string   `xml:"phase,attr"`
	Duration string   `xml:"duration,attr"`
	Unit     string   `xml:"unit,attr"`
}

// Load is
type Load struct {
	XMLName       xml.Name       `xml:"load"`
	ArrivalPhases []ArrivalPhase `xml:"arrivalphase"`
}

// UserAgent is
type UserAgent struct {
	XMLName     xml.Name `xml:"user_agent"`
	Probability string   `xml:"probability"`
	InnerTxt    string   `xml:",innerxml"`
}

// Option is
type Option struct {
	XMLName xml.Name `xml:"option"`
	Type    string   `xml:"type,attr"`
	Name    string   `name:"name"`
}

// Options is
type Options struct {
	XMLName xml.Name `xml:"options"`
	Options []Option `xml:""option`
}

// HttpHeader is ..
type HttpHeader struct {
	XMLName xml.Name `xml:"http_header"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

// Http is ..
type Http struct {
	XMLName          xml.Name `xml:"http"`
	Url              string   `xml:"url,attr"`
	Method           string   `xml:"method,attr"`
	Version          string   `xml:"version,attr,omitempty"`
	ContentsFromFile string   `xml:"contents_from_file,attr,omitempty"`
}

// Request is ..
type Request struct {
	XMLName xml.Name `xml:"request"`
	Http    Http     `xml:"http"`
}

// Thinktime is
type Thinktime struct {
	XMLName xml.Name `xml:"thinktime"`
	Value   string   `xml:"value,attr"`
	Random  string   `xml:"random,attr"`
}

// Session is
type Session struct {
	XMLName     xml.Name `xml:"session"`
	Name        string   `xml:"name,attr"`
	Probability string   `xml:"probability,attr,omitempty"`
	Type        string   `xml:"type,attr"`
	Requests    []interface{}
}

// Sessions is
type Sessions struct {
	XMLName  xml.Name  `xml:"sessions"`
	Sessions []Session `xml:"session"`
}

// Tsung is
type Tsung struct {
	XMLName    xml.Name   `xml:"tsung"`
	Loglevel   string     `xml:"loglevel,attr"`
	Version    string     `xml:"version,attr"`
	Clients    Clients    `xml:"clients"`
	Servers    Servers    `xml:"servers"`
	Monitoring Monitoring `xml:"monitor"`
	Load       Load       `xml:"load"`
	Options    Options    `xml:"options"`
	Sessions   Sessions   `xml:"sessions"`
}
