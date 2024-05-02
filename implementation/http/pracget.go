package noteshttp

import (
	"fmt"
	"net/http"
)

func ContentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	} //If there is an error, we will return an empty string, which would be the empty response, and the error

	defer resp.Body.Close() //Make sure we close the body

	ctype := resp.Header.Get("Content-Type")

	if ctype == "" { //Return error if Content-Type header is not found
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, nil //If everythin is ok, we will return the contentype and nil for the error, since there is no error.
}
