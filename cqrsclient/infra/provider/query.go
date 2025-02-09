package provider

import (
	"cqrsclient/infra/connect"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
)

// Queryクライアントプロバイダ
type QueryClientProvider struct {
	Category v1.CategoryQueryClient // カテゴリ用クライアント
	Product  v1.ProductQueryClient  // 商品用クライアント
}

// コンストラクタ
func NewQueryClientProvider(connector connect.ServerConnector) (*QueryClientProvider, error) {
	if connect, err := connector.Connect(); err != nil {
		return nil, err
	} else {
		// カテゴリ用クライアントを生成する
		category := v1.NewCategoryQueryClient(connect)
		// 商品用クライアントを生成する
		product := v1.NewProductQueryClient(connect)
		// プロバイダの生成
		return &QueryClientProvider{Category: category, Product: product}, nil
	}
}
