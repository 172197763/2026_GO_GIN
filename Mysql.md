## Upsert
命令：INSERT INTO ... ON DUPLICATE KEY UPDATE ...<br>
效果：主键or唯一键冲突时，更新数据，不冲突时，插入数据<br>
注意：affected-rows：insert成功1、update成功2、没有更新0
