package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
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
	// TODO: customize client!
	resp, err := http.Post(graphqlAddr(), jsonContentType, prepareQuery(q))

	Convey("_ http post error should be nil", func() {
		So(err, ShouldBeNil)

		Convey("_ response body should be read and unmarshalled without errors", func() {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Println("couldn't read body, err: ", err, "query: ", q)
			}
			So(err, ShouldBeNil)

			if err = json.Unmarshal(b, &res); err != nil {
				Printf("couldn't unmarshal body: err: %v | body: `%s` | query: `%s", err, b, q)
			}
			So(err, ShouldBeNil)

			Convey("_ response status should be ok", func() {
				So(resp.StatusCode, ShouldEqual, http.StatusOK)

				checkResponse()
			})
		})

	})
}

func httpAddr() string {
	return defaultHTTPAddr
}

func graphqlAddr() string {
	return httpAddr() + "query"
}

func prepareQuery(q string) io.Reader {
	q = strings.Replace(q, "\t", " ", -1)
	return bytes.NewBufferString(strings.Replace(q, "\n", " ", -1))
}
