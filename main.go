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
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString("Failed to read anything: " + err.Error() + "\n")
		os.Exit(1)
	}

	vals := make(url.Values)
	vals.Set("body", string(text))

	resp, err := http.PostForm("http://itsapad.appspot.com", vals)
	if err != nil {
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
