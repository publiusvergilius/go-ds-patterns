package adapter

import "testing"

func TestXmlToJson (t *testing.T) {
	var converter IConvert = Adaptee{}

	got, _ := converter.convert("<root><data>example</data></root>")
	want := "{\"data\":\"example\"}"

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
