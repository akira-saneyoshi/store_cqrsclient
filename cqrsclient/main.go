package main

import (
	"cqrsclient/presen"

	"go.uber.org/fx"
)

// @Title gRPCサンプルプログラム(CQRSサービス)
// @Version 1.0
// @Description カテゴリ情報と商品情報を管理するAPIサービス
// @TermOfService http://localhost:8081/
// @Licence.name Apache 2.0
// @Licence.url https://www.apache.org/licenses/LICENSE-2.0
// @Host localhost:8081
// @Basepath /
func main() {
	fx.New(
		presen.PresonDepend,
	).Run()
}
