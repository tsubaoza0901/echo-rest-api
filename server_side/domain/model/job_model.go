package model

import "time"

// Job Job構造体
// 役割：案件情報のモデル
type Job struct {
	ID        uint
	JobTitle  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// JobSearchRequest JobSearchRequest構造体
// 役割：
type JobSearchRequest struct {
	JobTitle string `query:"job_title"`
}