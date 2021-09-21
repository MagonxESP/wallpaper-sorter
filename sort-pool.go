package main

const MaxQueueLength = 10

type SortPool struct {
	queue [MaxQueueLength]string
	sorter *Sorter
}

func NewSortPool(sorter *Sorter) SortPool {
	return SortPool{
		queue: [MaxQueueLength]string{""},
		sorter: sorter,
	}
}

func (s *SortPool) Sort(filePath string) {
	s.AddToQueue(filePath)

	go s.StartQueue()
}

func (s *SortPool) StartQueue() {
	for _, queuePath := range s.queue {
		// TODO: sort wallpaper
	}
}

func (s *SortPool) FindAvailablePosition(position chan int) {
	for {
		for index, queuePath := range s.queue {
			if queuePath == "" {
				position <- index
			}
		}
	}
}

func (s *SortPool) AddToQueue(filePath string) {
	position := make(chan int)

	go s.FindAvailablePosition(position)

	s.queue[<-position] = filePath
}