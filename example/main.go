package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	boot "github.com/kysion/oss-library/example/internal"

	_ "github.com/kysion/oss-library/internal/logic"
)

func main() {
	boot.Main.Run(gctx.New())
}
