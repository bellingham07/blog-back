package crontask

import (
	"blog-back/database"
	"blog-back/model"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// 初始化生日定时任务
func InitBirthdayTask(c *cron.Cron) {
	_, err := c.AddFunc("0 8 * * *", checkBirthdays) // 每天北京时间8点执行
	if err != nil {
		log.Fatalf("创建生日定时任务失败: %v", err)
	}
}

func checkBirthdays() {
	db := database.GetDB()

	now := time.Now()
	today := fmt.Sprintf("%02d-%02d", now.Month(), now.Day())

	var birthdays []model.Birthday
	if err := db.Where("DATE_FORMAT(birthday, '%m-%d') = ?", today).Find(&birthdays).Error; err != nil {
		log.Printf("查询生日数据失败: %v", err)
		return
	}

	for _, user := range birthdays {
		if err := sendBirthdayEmail(user.Email, user.Name); err != nil {
			log.Printf("邮件发送失败至 %s: %v", user.Email, err)
			continue
		}
		log.Printf("已发送祝福邮件至 %s", user.Email)
	}
}

// 邮件发送函数（需实现具体邮件服务）
func sendBirthdayEmail(to, name string) error {
	// 示例使用SMTP发送（需配置邮件服务）
	// 实际应替换为您的邮件服务商配置
	emailContent := fmt.Sprintf(`
		<h1>亲爱的 %s：</h1>
		<p>今天是您的生日，祝您生日快乐！</p>
	`, name)

	return SendEmail(to, "生日祝福", emailContent)
}
