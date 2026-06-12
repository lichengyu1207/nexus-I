package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Iichengyu1207/nexus-I/backend/internal/config"
	"github.com/Iichengyu1207/nexus-I/backend/pkg/db"
)

type ParseRequest struct {
	VoiceText string `json:"voice_text"`
	UserID    string `json:"user_id,omitempty"`
}

type ParseResponse struct {
	Intent    string            `json:"intent"`
	Params    map[string]string `json:"params"`
	Confident float64           `json:"confident"`
}

var dbConn *sql.DB

func parseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ParseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := req.UserID
	if userID == "" {
		userID = "anonymous"
	}

	// 第一周最简单的规则匹配
	intent := "unknown"
	params := make(map[string]string)

	if contains(req.VoiceText, "剪辑") {
		intent = "clip"
		params["source"] = "gallery"
	} else if contains(req.VoiceText, "封面") {
		intent = "cover"
	} else if contains(req.VoiceText, "发布") {
		intent = "publish"
	} else if contains(req.VoiceText, "查询") || contains(req.VoiceText, "看看") {
		intent = "query"
	}

	resp := ParseResponse{
		Intent:    intent,
		Params:    params,
		Confident: 0.85,
	}

	// 写入解析记录到 SQLite
	if dbConn != nil {
		_, err := dbConn.ExecContext(r.Context(),
			"INSERT INTO command_logs (user_id, raw_voice_text, parsed_intent) VALUES (?, ?, ?)",
			userID, req.VoiceText, intent,
		)
		if err != nil {
			log.Printf("写入 command_log 失败: %v", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "nexus-gateway",
	})
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func main() {
	// 加载配置
	cfg := config.Load()
	log.Printf("环境: %s | 数据库: %s | 路径: %s", cfg.AppEnv, cfg.DBType, cfg.SQLitePath)

	// 初始化数据库
	var err error
	dbConn, err = db.NewDB(cfg)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	if dbConn != nil {
		defer dbConn.Close()
		log.Println("SQLite 连接成功:", cfg.SQLitePath)
	} else {
		log.Println("未配置数据库，将跳过日志存储")
	}

	// 注册路由
	http.HandleFunc("/v1/command/parse", parseHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Gateway started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
