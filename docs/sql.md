


```sql
主键 自增
CREATE SEQUENCE gc_goods_level_price_id START 1000;
```
在 表时
```sql
nextval('gc_goods_level_price_id::regclass')
```