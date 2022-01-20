package template

var (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/MockyBang/go-zero/core/stores/builder"
	"github.com/MockyBang/go-zero/core/stores/cache"
	"github.com/MockyBang/go-zero/core/stores/sqlc"
	"github.com/MockyBang/go-zero/core/stores/sqlx"
	"github.com/MockyBang/go-zero/core/stringx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/MockyBang/go-zero/core/stores/builder"
	"github.com/MockyBang/go-zero/core/stores/sqlc"
	"github.com/MockyBang/go-zero/core/stores/sqlx"
	"github.com/MockyBang/go-zero/core/stringx"
)
`
)
