package bitbucket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type BitbucketError struct {
	Message string
	Fields  map[string][]string
}

func DecodeError(e map[string]interface{}) error {
	var bitbucketError BitbucketError
	err := mapstructure.Decode(e["error"], &bitbucketError)
	if err != nil {
		return err
	}

	return errors.New(bitbucketError.Message)
}

func httpResponseErr(resp *http.Response) error {
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		if resp.Body != nil {
			buff := new(bytes.Buffer)
			io.Copy(buff, resp.Body)
			var errorStr ResponseError
			err := json.Unmarshal(buff.Bytes(), &errorStr)
			if err == nil {
				return fmt.Errorf("%d: %s %v", resp.StatusCode, errorStr.Error.Message, errorStr.Error.Fields)
			}
		}
		return fmt.Errorf(resp.Status)
	}
	return nil
}

type ResponseError struct {
	Type  string `json:"type"`
	Error Error  `json:"error"`
}

type Error struct {
	Fields  Fields `json:"fields"`
	Message string `json:"message"`
}

type Fields struct {
	TargetHash string `json:"target.hash"`
}
