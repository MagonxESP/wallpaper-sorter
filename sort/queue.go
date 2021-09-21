package sort

import (
	"fmt"
	"log"
)

const MaxQueueLength = 10

const (
	QueueRunning = "running"
	QueueStop = "stop"
	QueueFinished = "finished"
)

type Queue struct {
	queue  [MaxQueueLength]string
	sorter Sorter
	pending int
	status string
}

func NewQueue(sorter Sorter) Queue {
	return Queue{
		queue: [MaxQueueLength]string{""},
		sorter: sorter,
		pending: 0,
		status: QueueStop,
	}
}

func (s *Queue) ProcessQueue() {
	for {
		if s.status == QueueStop {
			return
		}

		for index, queuePath := range s.queue {
			if queuePath != "" {
				log.Println(fmt.Sprintf("Sorting: %s", queuePath))
				s.sorter.Sort(queuePath)
				s.queue[index] = ""
				s.pending--
			}
		}

		if s.pending == 0 {
			s.status = QueueFinished
			return
		}
	}
}

func (s *Queue) FindAvailablePosition(position chan int) {
	for {
		for index, queuePath := range s.queue {
			if queuePath == "" {
				position <- index
			}
		}
	}
}

func (s *Queue) Add(filePath string) {
	log.Println(fmt.Sprintf("Adding path to the queue: %s", filePath))

	position := make(chan int)

	go s.FindAvailablePosition(position)

	s.pending++
	s.queue[<-position] = filePath
}

// Start or resume the queue
func (s *Queue) Start() {
	if s.pending > 0 && s.status != QueueRunning {
		s.status = QueueRunning
		go s.ProcessQueue()
	}
}

// Stop the current queue
func (s *Queue) Stop() {
	s.status = QueueStop
}

func (s *Queue) WaitUntilFinished() {
	for {
		if s.status == QueueFinished {
			return
		}
	}
}