package bitbucket

import (
	"encoding/json"
	"os"

	"github.com/k0kubun/pp"
)

type Branches struct {
	c *Client
}

func (b *Branches) Create(bo *BranchesOptions) (interface{}, error) {
	data := b.buildBranchesBody(bo)
	urlStr := b.c.requestUrl("/repositories/%s/%s/refs/branches", bo.Owner, bo.RepoSlug)
	return b.c.execute("POST", urlStr, data)
}

type branchesBody struct {
	Name   string `json:"name"`
	Target Target `json:"target"`
}

type Target struct {
	Hash string `json:"hash"`
}

func (b *Branches) buildBranchesBody(bo *BranchesOptions) string {
	body := branchesBody{
		Name: bo.Name,
		Target: Target{
			Hash: bo.Commit,
		},
	}
	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
		os.Exit(9)
	}
	return string(data)
}
