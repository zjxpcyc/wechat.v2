package utils

import (
	"time"

	"github.com/zjxpcyc/wechat.v2/declares"
)

// ScheduleTask 定时任务
type ScheduleTask struct {
	done chan bool
	task declares.Task
}

// NewScheduleTask 新建任务
func NewScheduleTask(task declares.Task) *ScheduleTask {
	return &ScheduleTask{
		done: make(chan bool),
		task: task,
	}
}

// Run 启动任务
// 如果需要延时启动, 可以设置 delay 参数
func (s *ScheduleTask) Run(delay ...time.Duration) {
	go s.start(delay...)
}

// Stop 停止任务
func (s *ScheduleTask) Stop() {
	s.done <- true
}

func (s *ScheduleTask) start(delay ...time.Duration) {
	go func() {
		if delay != nil && len(delay) > 0 {
			time.Sleep(delay[0])
		}

		d := s.task() // 上次任务, 返回下一次的执行周期
		for {
			time.Sleep(d)
			d = s.task()
		}
	}()

	for {
		select {
		case done := <-s.done:
			if done {
				return
			}
		}
	}
}
