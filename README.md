# 说明
> 基于Docker与Go的OSS同步工具

# 功能
- 基于Docker运行
- 通过环境变量配置OSS信息
- 挂载要同步的目录到容器内/sync
- 自定义cron任务时间

# 下载
```
git clone https://github.com/mathcoder23/docker-oss-sync.git
# 加速
git clone https://hub.fastgit.org/mathcoder23/docker-oss-sync.git
```

# 编译
```shell script
docker built -t docker-oss-sync:1.0 .
```
```shell script
cd go 
chmod +x build.sh 
./build.sh
```
# 运行模板
```shell script
docker run -e TASK_CRON="0 2 * * *" \
-e OSS_ENDPOINT=oss-cn-beijing.aliyuncs.com \
-e OSS_KEY_ID=xxx \
-e OSS_KEY_SECRET=xxx \
-e OSS_BUCKET_NAME=xxx \
-e OSS_PREFIX=backups \
-v xxx:/sync \
docker-oss-sync:1.0
```

# 测试运行
先设置 TASK_CRON为`@every 10s` 每十秒触发一次

# 正式运行
- `0 * * * *` 一小时一次
- `0 2 * * *` 每天两点

# 脚本运行
- 修改脚本内的环境变量
```shell script
chmod +x start.sh
./start.sh
```

# 时间调度配置参考
## TASK_CRON
> 与Linux 中crontab命令相似，cron库支持用 5 个空格分隔的域来表示时间。这 5 个域含义依次为：

- Minutes：分钟，取值范围[0-59]，支持特殊字符* / , -；
- Hours：小时，取值范围[0-23]，支持特殊字符* / , -；
- Day of month：每月的第几天，取值范围[1-31]，支持特殊字符* / , - ?；
- Month：月，取值范围[1-12]或者使用月份名字缩写[JAN-DEC]，支持特殊字符* / , -；
- Day of week：周历，取值范围[0-6]或名字缩写[JUN-SAT]，支持特殊字符* / , - ?。
注意，月份和周历名称都是不区分大小写的，也就是说SUN/Sun/sun表示同样的含义（都是周日）。

特殊字符含义如下：

- *：使用*的域可以匹配任何值，例如将月份域（第 4 个）设置为*，表示每个月；
- /：用来指定范围的步长，例如将小时域（第 2 个）设置为3-59/15表示第 3 分钟触发，以后每隔 15 分钟触发一次，因此第 2 次触发为第 18 分钟，第 3 次为 33 分钟。。。直到分钟大于 59；
- ,：用来列举一些离散的值和多个范围，例如将周历的域（第 5 个）设置为MON,WED,FRI表示周一、三和五；
- -：用来表示范围，例如将小时的域（第 1 个）设置为9-17表示上午 9 点到下午 17 点（包括 9 和 17）；
- ?：只能用在月历和周历的域中，用来代替*，表示每月/周的任意一天。
了解规则之后，我们可以定义任意时间：

30 * * * *：分钟域为 30，其他域都是*表示任意。每小时的 30 分触发；
30 3-6,20-23 * * *：分钟域为 30，小时域的3-6,20-23表示 3 点到 6 点和 20 点到 23 点。3,4,5,6,20,21,22,23 时的 30 分触发；
0 0 1 1 *：1（第 4 个） 月 1（第 3 个） 号的 0（第 2 个） 时 0（第 1 个） 分触发。

> 定义调度任务的执行时间,具体参见go cron的语法格式
- @yearly：也可以写作@annually，表示每年第一天的 0 点。等价于0 0 1 1 *；
- @monthly：表示每月第一天的 0 点。等价于0 0 1 * *；
- @weekly：表示每周第一天的 0 点，注意第一天为周日，即周六结束，周日开始的那个 0 点。等价于0 0 * * 0；
- @daily：也可以写作@midnight，表示每天 0 点。等价于0 0 * * *；
- @hourly：表示每小时的开始。等价于0 * * * *。

## 挂载容器/sync
> 宿主机目录需要挂载到容器的/sync下

## OSS_PREFIX
> bucket的对象key前缀

`其它参数很明显，就不说明了。`
