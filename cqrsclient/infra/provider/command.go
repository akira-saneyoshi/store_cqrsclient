package provider

import (
	"cqrsclient/infra/connect"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
)

// Commandクライアントプロバイダ
type CommandClientProvider struct {
	Category v1.CategoryCommandClient
	Product  v1.ProductCommandClient
}

// コンストラクタ
func NewCommandClientProvider(connector connect.ServerConnector) (*CommandClientProvider, error) {
	// Command Serviceの接続
	if connect, err := connector.Connect(); err != nil {
		return nil, err
	} else {
		// カテゴリ用クライアントを生成する
		category := v1.NewCategoryCommandClient(connect)
		// 商品用クライアントを生成する
		product := v1.NewProductCommandClient(connect)
		// プロバイダの生成
		return &CommandClientProvider{Category: category, Product: product}, nil
	}
}
