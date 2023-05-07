# mysql

## SQL 分类
* DDL 数据定义语言 
  * CREATE
  * DROP
  * ALERT
* DML 数据操作语言
  * INSERT
  * DELETE
  * UPDATE
  * SELECT
* DCL 数据控制语言 
  * GRANT
  * REVOKE
  * COMMIT
  * ROLLBACK
  * SAVEPOINT 

## 14 视图
### 数据库对象
* 表
* 数据字典
* 约束 
  * 执行数据校验的规则，用于保证数据完整性的规则
* 视图
 * 一个或多个数据表里的数据的逻辑显示，视图并不存储数据
* 索引
  * 提高查取性能
* 存储过程
* 存储函数
* 触发器
  * 事件监听器，当数据发生特定事件后，触发器被触发，完成相应处理

### 视图
#### 1.为什么使用视图
视图使用的是表的一部分，针对不同的用户展示不同的数据

#### 2.视图的理解
* 视图是虚拟表，占内存小
* 视图建立在数据表的基础上，此时数据表叫做基表
* 视图不保存数据，但对视图中数据进行增删改查操作时，对应数据表中的数据也会发生变化
#### 3. 视图的使用 
* 创建视图
```sql
-- 查询语句中的字段别名会作为视图中字段的名称出现
CREATE VIEW 视图名称[(字段列表)] AS 查询语句 

-- 视图中的字段不是表中的字段时，要给字段起个别名
```
```sql
--利用视图对数据进行格式化
CREATE VIEW emp_depart
AS
SELECT CONCAT(last_name,"(",department_name,")") AS emp_dept
FROM employees e JOIN department d 
WHERE e.depart  ment_id = d.deparment_id
```
* 基于视图创建视图
```sql
CREATE VIEW  emp 
AS
SELECT employee_id,last_name
FROM vu_empl 
```
* 查看视图
```sql
-- 查看数据库 表对象 视图对象
SHOW TABLES

-- 查看视图结构
DESCRIBE 视图名称
    
-- 查看视图属性信息
SHOW TABLE STATUS LIKE '视图名称' \G
    
-- 查看视图的详细定义信息
SHOW CREATE VIEW '视图名称'
```
#### 4.不可更新的视图
只有视图中的字段和基表中的字段一一对应的时候,才能对视图进行更新或者删除操作

**注意**
视图一般用来查询数据

#### 5.修改视图
```sql
CREATE [OR REPALCE ] VIEW 视图名称[(字段列表)] AS 查询语句

ALERT VIEW 视图名称 AS 查询语句
```
#### 6.删除视图
```sql
DROP VIEW 视图名称
```
#### 7.总结
* 优点
  * 操作简单
  * 减少数据冗余
  * 数据安全
  * 适应灵活多变的需求
  * 能够分解复杂的查询逻辑
* 不足
  * 实际数据结构发生变化,我们要对试图进行维护,如果视图含有复杂逻辑,会造成难以维护问题

## 15 存储过程与存储函数
### 1 存储过程
 - - -
#### 1.1 理解
* 思想:一组经过**预先编译**的SQL语句的封装
* 执行过程：存储过程预先存储在MySQL服务器上，需要执行的时候，客户端只需要向服务端发送命令，服务器端
就可以把预先存储好的一系列SQL语句全部执行
* 优点
  * 简化操作，提高SQL语句的复用性
  * 减少操作中的失误，提高效率
  * 减少网络传输量（客户端不需要把所有的SQL语句通过网络发给服务器）
  * 减少SQL语句暴露的风险，提高安全性
* 和视图，存储函数的不同
  * 视图是虚拟表，不对底层数据直接操作，存储过程直接操作底层数据，能够处理一些发杂的逻辑
  * 和函数一样可以直接调用，但相比较于函数 没有返回值
#### 1.2分类
存储过程的参数可以是IN，OUT，和INOUT 
* 没有参数
* 仅带IN类型
* 仅带OUT类型
* 即带IN又带OUT
* 带INOUT
### 2 创建存储过程
- - - 
#### 2.1 语法
```sql
CREATE PROCEDURE 存储过程名(IN|OUT|INOUT 参数名 参数类型, ...)
[characteristics...]
BEGIN 
    存储过程体
END
```
### 3 存储函数
-- -
#### 3.1 语法分析
```sql
CREATE FUNCTION 函数名(参数名 参数类型,...)
RETURNS 返回值类型
[characteristics...]
BEGIN 
 函数体 函数体中包含RETURN 语句
END
```
#### 3.2 调用存储函数
```sql
SELECT 函数名(参数列表)
```

#### 3.3 存储过程  存储函数查看
* 查看创建信息
```sql
SHOW CREATE [ PROCEDURE | FUNCTION ] 存储过程/存储函数名
```
* 查看状态信息
```sql
SHOW  [ PROCEDURE | FUNCTION ]  STATUS
SHOW  [ PROCEDURE | FUNCTION ]  STATUS  LIKE 存储过程/存储函数名
```






















