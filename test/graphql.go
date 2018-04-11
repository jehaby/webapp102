package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type graphErr struct {
	Message   string
	Locations []struct {
		Line, Column int
	}
}

func unmarshalBody(body io.ReadCloser, res interface{}) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		panic(err)
	}
}
