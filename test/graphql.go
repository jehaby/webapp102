package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service/auth"
	"github.com/jehaby/webapp102/test/data"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	defaultHTTPAddr = "http://localhost:8899/"

	jsonContentType = "application/json"
)

var httpClient *http.Client

type graphErr struct {
	Message   string
	Locations []struct {
		Line, Column int
	}
}

func initHTTPClient(cfg config.C) {
	url := &url.URL{Scheme: "http", Host: cfg.HTTP.Addr}
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(url, []*http.Cookie{
		&http.Cookie{
			Name:  "jwt",
			Value: mustTokenFromUser(auth.New(cfg.Auth), &data.TestUser),
		},
	})

	httpClient = &http.Client{
		Jar: jar,
	}
}

func queryGraphql(q string, res interface{}, checkResponse func()) {

	resp, err := httpClient.Post(graphqlAddr(), jsonContentType, prepareQuery(q))

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

func escapeQuotes(q string) string {
	return strings.Replace(q, `"`, `\"`, -1)
}

func mustTokenFromUser(jwtAuth *auth.JwtAuth, u *entity.User) string {
	token, err := jwtAuth.TokenFromUser(&data.TestUser, time.Hour)
	if err != nil {
		log.Panicf("token from user (%v) returned error: %v", u, err)
	}
	return token
}
