# xianzhi_crawl

# 介绍

该项目为学习Go 语言的练手项目，将先知社区的文章保存为md文件，并将阿里云的图片保存到本地直接进行替换。

效果如下：在 main 函数中获取页数

![](https://cdn.jsdelivr.net/gh/h0ld1rs/image/image202401251023319.png)

该页数指的是：/t 后面的部分

![](https://cdn.jsdelivr.net/gh/h0ld1rs/image/image202401251024748.png)

使用前请自行下载相对应版本的chromeDriver

目前效果如下：

执行之后，会在当前目录下生成doc文件夹，在doc文件夹下创建该文章标题的文件夹和md文件，在其文件夹下，将图片保存到assert 文件夹

![](https://cdn.jsdelivr.net/gh/h0ld1rs/image/image202401251027438.png)

![](https://cdn.jsdelivr.net/gh/h0ld1rs/image/image202401251025830.png)

效果如下：

![](https://cdn.jsdelivr.net/gh/h0ld1rs/image/image202401251026289.png)

![](https://cdn.jsdelivr.net/gh/h0ld1rs/image/image202401251026949.png)

# Todo

* 增加代理池，将所有先知的文章保存
* 增加对比列表，优化速率
