// © 2012 Steve McCoy. Available under the MIT License.

/*
topad pastes its standard input to itsapad.appspot.com and prints the URL
on standard output.

For example:
	echo Hello | topad
will paste "Hello" to itsapad.appspot.com and print out the URL of the paste.
*/
package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var serviceURL = flag.String("url", "https://pad.mccoy.space", "URL to pad service")

func main() {
	flag.Parse()

	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString("Failed to read anything: " + err.Error() + "\n")
		os.Exit(1)
	}

	vals := make(url.Values)
	vals.Set("body", string(text))

	skipRedirect := errors.New("I want the Location")

	c := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return skipRedirect
		},
	}

	resp, err := c.PostForm(*serviceURL, vals)
	if err != nil && err.(*url.Error).Err != skipRedirect {
		os.Stderr.WriteString("Request failed: " + err.Error() + "\n")
		os.Exit(1)
	}
	defer resp.Body.Close()

	loc, err := resp.Location()
	if err != nil {
		os.Stderr.WriteString("Response failed: " + err.Error() + "\n")
		os.Exit(1)
	}
	os.Stderr.WriteString(loc.String() + "\n")
}
