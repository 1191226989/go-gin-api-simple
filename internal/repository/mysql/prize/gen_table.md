#### go_gin_api_simple.prize 
奖品表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int unsigned | PRI | NO | auto_increment |  |
| 2 | name | 奖品名称 | varchar(100) | UNI | YES |  |  |
| 3 | image | 奖品图片 | varchar(100) |  | YES |  |  |
| 4 | worth | 奖品价值 | decimal(10,0) |  | YES |  |  |
| 5 | content | 奖品描述 | varchar(100) |  | YES |  |  |
| 6 | is_used | 是否启用 1:是  -1:否 | tinyint(1) |  | NO |  | 1 |
| 7 | created_at | 创建时间 | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 8 | updated_at | 更新时间 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
