package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"

	conf "github.com/MockyBang/go-zero/tools/goctl/config"
	"github.com/MockyBang/go-zero/tools/goctl/rpc/parser"
	"github.com/MockyBang/go-zero/tools/goctl/util/format"
	"github.com/MockyBang/go-zero/tools/goctl/util/pathx"
)

const configTemplate = `package config

import "github.com/MockyBang/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
}
`

// GenConfig generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/MockyBang/go-zero/tree/master/tools/goctl/config/config.go
func (g *Generator) GenConfig(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetConfig()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "config")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, configTemplateFileFile, configTemplate)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
}
