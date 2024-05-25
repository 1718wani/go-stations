package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
    mux := http.NewServeMux()
    todoService := service.NewTODOService(todoDB)
    mux.Handle("/todos", handler.NewTODOHandler(todoService))
    return mux
}

// 全体の流れ
// 1. 郵便局の仕分け機（ルーター）を作る: http.NewServeMuxで新しいルーターを作成。
// 2. 道具箱（サービス）を作る: service.NewTODOServiceでTODOリストを管理するサービスを作成。
// 3. 郵便受け（ハンドラー）を設定: mux.Handleで、/todosに来たリクエストをTODOHandlerに送るように設定。
// 4. 仕分け機を返す: 最後に、作ったルーターを返す。