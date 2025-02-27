package main

import (
	"blog-back/database"
	"blog-back/router"
)

func main() {

	// 初始化数据库
	db := database.InitDB() //初始化数据库
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to close database" + err.Error())
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
	}()

	// 初始化定时任务
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	crontask.InitBirthdayTask(c)
	c.Start()
	defer c.Stop()

	// 初始化redis
	// RedisUtil.InitRedis()
	// 初始化router
	r := router.Routers()

	// 设置监听端口
	panic(r.Run(":9090"))
}
