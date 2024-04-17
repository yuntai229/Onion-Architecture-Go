table "threads" {
  schema = schema.example
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "content" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = bigint
  }
  column "updated_at" {
    null = true
    type = datetime
  }
  column "created_at" {
    null = true
    type = datetime
  }
  column "deleted_at" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "threads_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = CASCADE
    on_delete   = NO_ACTION
  }
  index "user_id" {
    columns = [column.user_id]
  }
}
table "users" {
  schema = schema.example
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "name" {
    null = true
    type = varchar(255)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "hash_password" {
    null = true
    type = varchar(255)
  }
  column "updated_at" {
    null = true
    type = datetime
  }
  column "created_at" {
    null = true
    type = datetime
  }
  column "deleted_at" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
}
schema "example" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
