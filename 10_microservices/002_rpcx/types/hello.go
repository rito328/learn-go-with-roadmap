package types

// HelloRequest はクライアントが送るリクエストの構造体
type HelloRequest struct {
	Name string
}

// HelloResponse はサーバーが返すレスポンスの構造体
type HelloResponse struct {
	Message string
}
