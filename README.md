# goexamer
记忆小工具

* 基于walk库的简单记忆工具
* 可以根据用户脚本帮助用户记忆内容

## 使用方法: 
0. 用户编写一个小脚本, 包括所有要记忆的项和命令
1. 使用 goexamer -i "脚本名称" 执行脚本
2. 系统会每次随机抽取一项，显示问题, 之后用户选择是否记得
3. 如果该项选择为"不记得"，会在之后的复习阶段重新出现以帮助用户巩固复习

### 标题
```
title: titlevalue
\NextLine
```

### 批次(batch)
```
[BatchName]
\NextLine
\@action
item...
```

### 项(item)
```
#itemName: value
\NextLine
\@action
```

### 动作命令(action)
* @deduct:itemName item的测试次数-1 
* @mark:itemName item的测试次数+1 
* @jmp:itemName 当前结束后强制跳转并执行指定的item 
* @link:itemName 当前结束后进入指定的item, 如果该item已经执行过, 就不执行 
* @set:count, @set:itemName:count 设置item的测试次数为大于0的数
* @showImg:imageName 显示当前img文件夹下的一张图片, 在交互之后进行
* @execute:itemName 改变当前action的执行对象为指定对象, 只针对在交互之后执行的action有效 
* @img:imageName 显示当前img文件夹下的一张图片, 在交互之前进行
* @help[:actionName] 显示帮助

