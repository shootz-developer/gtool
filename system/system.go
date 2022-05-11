package system

import (
	"log"
	"runtime"
	"time"

	"github.com/robfig/cron"
)

// StartTask 封装cron包，启动定时任务 6个字段 [second minute hour day month weekday]
// 如 "0 */5 * * * ?" 每隔5分钟
func StartTask(spec string, cmd func()) {
	f := func() {
		defer func() {
			if e := recover(); e != nil {
				buf := make([]byte, 16*1024*1024)
				buf = buf[:runtime.Stack(buf, false)]
				log.Printf("[PANIC]%v\n%s\n", e, buf)
			}
		}()
		log.Printf("start cron task")
		cmd()
		log.Printf("end cron task")
	}

	c := cron.New()
	e := c.AddFunc(spec, f)
	if e != nil {
		panic(e)
	}

	c.Start()
}

// StartTaskAtOnce 启动时先执行一次，然后调用 StartTask
func StartTaskAtOnce(spec string, timeout time.Duration, cmd func()) {
	timeStart := time.Now().UnixNano()
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16*1024*1024)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("[PANIC]%v\n%s", e, buf)
		}
		if time.Now().UnixNano() > timeout.Nanoseconds()+timeStart { // 超时
			log.Printf("process exceed timeLimit %v", timeout)
		}
	}()

	log.Printf("start cron task")
	cmd()
	log.Printf("end cron task")
	StartTask(spec, cmd)
}
