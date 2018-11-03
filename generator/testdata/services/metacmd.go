// Code generated by octo-cli/generator; DO NOT EDIT.

package services

import "github.com/octo-cli/octo-cli/internal"

type MetaCmd struct {
	Get MetaGetCmd `cmd:"" help:"Get"`
}

type MetaGetCmd struct {
	internal.BaseCmd
}

func (c *MetaGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/meta")
	return c.DoRequest("GET")
}
