# PilotJobService
单机版定时任务框架

## 使用方法
在job目录下新建自定义job类继承AbstractJob类，实现GetName方法（定义任务名），实现Do方法（任务的执行逻辑），然后在init函数中将任务添加到Jobs切片中。

配置文件使用yaml格式，在Job项下新增任务配置，格式如下：
```yaml
Job:
  JobName:
    Enabled: true
    Cron: "0 * * * * "
```

其中，JobName对应GetName方法中定义的任务名，Enabled为true时任务启用，Cron为任务执行的时间表达式。
项目使用viper监听了配置文件变更，对Enabled和Cron的修改会实时生效。

