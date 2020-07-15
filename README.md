# goexamer
记忆小工具

* 基于walk库的简单记忆工具
* 可以根据用户脚本帮助用户记忆内容

# 命令
* @deduct:item name item的测试次数-1 
* @mark:item name item的测试次数+1 
* @jmp:item name 当前结束后强制跳转并执行指定的item 
* @link:item name 当前结束后进入指定的item, 如果该item已经执行过, 就不执行 
* @set:count 设置当前item的测试次数为大于0的数 
* @showImg:image name 显示当前img文件夹下的一张图片, 在交互之后进行 
* @execute:item name 改变当前action的执行对象为指定对象, 只针对在交互之后执行的action有效 
* @img:image name 显示当前img文件夹下的一张图片, 在交互之前进行 
* @help[:action name] 显示帮助 
