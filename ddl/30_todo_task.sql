/*
### SCHEMA
タスク管理 / todo
### TABLE
タスク / task
*/

DROP TABLE IF EXISTS todo.task CASCADE;

CREATE TABLE todo.task
(
  task_id serial NOT NULL,
  title varchar(120) NOT NULL,
  assignee varchar(3),
  status todo.task_status NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_user varchar(3) NOT NULL,
  modified_user varchar(3) NOT NULL,
  primary key (task_id)
);

COMMENT ON TABLE todo.task IS 'タスク';

COMMENT ON COLUMN todo.task.task_id IS 'タスクID';
COMMENT ON COLUMN todo.task.title IS 'タイトル';
COMMENT ON COLUMN todo.task.assignee IS '担当者';
COMMENT ON COLUMN todo.task.status IS 'ステータス';
COMMENT ON COLUMN todo.task.created_at IS '登録日時';
COMMENT ON COLUMN todo.task.modified_at IS '更新日時';
COMMENT ON COLUMN todo.task.created_user IS '登録ユーザー';
COMMENT ON COLUMN todo.task.modified_user IS '更新ユーザー';

CREATE TRIGGER task_created BEFORE INSERT ON todo.task FOR EACH ROW EXECUTE PROCEDURE set_created_at();
CREATE TRIGGER task_modified BEFORE UPDATE ON todo.task FOR EACH ROW EXECUTE PROCEDURE set_modified_at();
