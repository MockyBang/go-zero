package javagen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MockyBang/go-zero/core/logx"
	"github.com/MockyBang/go-zero/tools/goctl/api/parser"
	"github.com/MockyBang/go-zero/tools/goctl/util/pathx"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

// JavaCommand the generate java code command entrance
func JavaCommand(c *cli.Context) error {
	apiFile := c.String("api")
	dir := c.String("dir")
	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	api, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}

	api.Service = api.Service.JoinPrefix()
	packetName := strings.TrimSuffix(api.Service.Name, "-api")
	logx.Must(pathx.MkdirIfNotExist(dir))
	logx.Must(genPacket(dir, packetName, api))
	logx.Must(genComponents(dir, packetName, api))

	fmt.Println(aurora.Green("Done."))
	return nil
}
