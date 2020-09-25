
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO jobs (id, job_title, recruitment_numbers, dev_start_date, dev_end_date, job_status_id, job_description, user_id) 
VALUES (1, 'テストタイトル', 2, '2020-07-01 12:00:00', '2020-08-01 12:00:00', 1, '案件詳細だーーーーーーーーーーーーーーーーーーーーーー！', 2);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM jobs;