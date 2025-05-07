package adapter

import "log"

type IConvert interface {
	convert( src string ) (string, error)
}

type XmlAdapter struct {
	adaptee Adaptee
}

func (a XmlAdapter) convert (src string) (string, error) {
	log.Println("Adapter process")
	return "", nil
}

type Adaptee struct {
	adapterType string
}

func (a Adaptee) convert (src string) (string, error) {
	log.Println("Adaptee conversion")
	return "", nil
}
