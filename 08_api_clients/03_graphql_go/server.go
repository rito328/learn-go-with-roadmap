package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// User 型の定義
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
			"age":  &graphql.Field{Type: graphql.Int},
		},
	},
)

// データソース（サンプル）
var users = []map[string]interface{}{
	{"id": 1, "name": "Alice", "age": 30},
	{"id": 2, "name": "Bob", "age": 25},
}

// クエリの定義
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					for _, user := range users {
						if user["id"] == id {
							return user, nil
						}
					}
					return nil, nil
				},
			},
		},
	},
)

// スキーマの作成
func initSchema() (graphql.Schema, error) {
	var schema graphql.Schema
	var err error

	schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		},
	)
	if err != nil {
		return graphql.Schema{}, fmt.Errorf("❌ Failed to create GraphQL schema: %v", err)
	}

	return schema, nil
}

// GraphQL ハンドラー（GET & POST 両対応）
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	var query string

	if r.Method == http.MethodGet {
		// GET クエリパラメータから取得
		query = r.URL.Query().Get("query")
	} else if r.Method == http.MethodPost {
		// POST の場合、リクエストボディを読み取る
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var req struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		query = req.Query
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// GraphQL の実行
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// JSON レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

var schema graphql.Schema

// メイン関数
func main() {
	var err error
	schema, err = initSchema()
	if err != nil {
		log.Fatalf("❌ Failed to init schema: %v", err)
	}

	http.HandleFunc("/graphql", graphqlHandler)

	port := 8080
	fmt.Printf("🚀 Server is running at http://localhost:%d/graphql\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
