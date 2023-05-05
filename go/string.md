## string

在go语言中，str中只能读，不能修改

让字符串修改的方法：
* 重新赋值
* 转化为slice类型 []byte(str) 


## slice

由三个部分组成 
* data 元素存储位置
* len 存了多少元素
* cap 可以存多少个元素

slice 公用底层数组地址

扩容规则
* 预估扩容后的容量
  * 如果oldCap*2 < cap  ->  newCap=cap
  * 否则
    * oldCap < 1024  -> newCap=oldCap*2
    * oldCap >= 1024  -> newCap=oldCap*1.25
* newCap个元素需要的内存大小
  * 预估容量*元素类型大小
* 匹配合适的内存规格

## 