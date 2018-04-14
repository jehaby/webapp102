package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	defaultHTTPAddr = "http://localhost:8899/"

	jsonContentType = "application/json"
)

type graphErr struct {
	Message   string
	Locations []struct {
		Line, Column int
	}
}

func queryGraphql(q string, res interface{}, checkResponse func()) {
	q = prepareQuery(q)

	// TODO: customize client!
	resp, err := http.Post(graphqlAddr(), jsonContentType, bytes.NewBufferString(q))

	Convey("Error should be nil", func() {
		So(err, ShouldBeNil)

		unmarshalBody(resp.Body, &res)

		Convey("HTTP status should be ok", func() {
			So(resp.StatusCode, ShouldEqual, http.StatusOK)

			checkResponse()
		})
	})
}

func unmarshalBody(body io.ReadCloser, res interface{}) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("couldn't read body", err)
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		log.Fatalf("couldn't unmarshal body: err: %v | body: `%s`", err, b)
	}
}

func httpAddr() string {
	return defaultHTTPAddr
}

func graphqlAddr() string {
	return httpAddr() + "query"
}

func prepareQuery(q string) string {
	q = strings.Replace(q, "\t", " ", -1)
	return strings.Replace(q, "\n", " ", -1)
}
