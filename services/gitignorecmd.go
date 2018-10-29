// Code generated by go-github-cli/generator; DO NOT EDIT

package services

type GitignoreCmd struct {
	ListTemplates GitignoreListTemplatesCmd `cmd:"" help:"Listing available templates"`
	GetTemplate   GitignoreGetTemplateCmd   `cmd:"" help:"Get a single template"`
}

type GitignoreListTemplatesCmd struct {
	baseCmd
}

func (c *GitignoreListTemplatesCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/gitignore/templates"
	return c.doRequest("GET")
}

type GitignoreGetTemplateCmd struct {
	baseCmd
	Name string `required:"" name:"name"`
}

func (c *GitignoreGetTemplateCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/gitignore/templates/:name"
	c.updateURLPath("name", c.Name)
	return c.doRequest("GET")
}
