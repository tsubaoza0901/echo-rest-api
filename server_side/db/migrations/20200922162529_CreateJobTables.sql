-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS jobs (
    id int(11) NOT NULL AUTO_INCREMENT,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL,
    job_title varchar(128) NOT NULL COMMENT "案件タイトル",
    recruitment_numbers int(11) NOT NULL COMMENT "案件募集人数",
    dev_start_date datetime NOT NULL COMMENT "案件募集開始日",
    dev_end_date datetime NOT NULL COMMENT "案件募集終了日",
    job_status_id int(11) NOT NULL COMMENT "募集ステータス管理ID",
    job_description text NOT NULL COMMENT "案件募集詳細",
    user_id int(11) NOT NULL COMMENT "案件募集ユーザーID",
    PRIMARY KEY(id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE jobs;
