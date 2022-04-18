package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	conf "github.com/MockyBang/go-zero/tools/goctl/config"
	"github.com/MockyBang/go-zero/tools/goctl/rpc/parser"
	"github.com/MockyBang/go-zero/tools/goctl/util"
	"github.com/MockyBang/go-zero/tools/goctl/util/format"
	"github.com/MockyBang/go-zero/tools/goctl/util/pathx"
	"github.com/MockyBang/go-zero/tools/goctl/util/stringx"
)

const etcTemplate = `Name: {{.serviceName}}.rpc
ListenOn: 127.0.0.1:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc
`

// GenEtc generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenEtc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetEtc()
	etcFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, fmt.Sprintf("%v.yaml", etcFilename))

	text, err := pathx.LoadTemplate(category, etcTemplateFileFile, etcTemplate)
	if err != nil {
		return err
	}

	return util.With("etc").Parse(text).SaveTo(map[string]interface{}{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
	}, fileName, false)
}
