package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// QueryServerと接続する
type queryConnector struct{}

// コンストラクタ
func NewqueryConnector() ServerConnector {
	return &queryConnector{}
}

// QueryServiceとの接続を確立する
func (ins *queryConnector) Connect() (*grpc.ClientConn, error) {
	server := "queryservice:8083" // サーバ名とポート番号
	if conn, err := grpc.Dial(    // gRPCサーバとの接続を確立する
		server, // 接続文字列
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 接続が確立するまで待つか、失敗した場合にエラーを即座に取得する
		grpc.WithBlock(),
	); err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}
