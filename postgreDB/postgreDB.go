package postgreDB

import (
	"container/heap"
	"database/sql"
	"fmt"
	"time"
)

type dbItem struct {
	id          int
	createdTime time.Time
}

type PriorityQueue []dbItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].createdTime.Before(pq[j].createdTime)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(dbItem)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func queryData(db *sql.DB, startTime time.Time, endTime time.Time) (images []dbItem, posts []dbItem) {
	imagesRow, err := db.Query("SELECT id, createdTime FROM IMAGES WHERE createdTime >= $1 AND createdTime < $2", startTime, endTime)
	if err != nil {
		fmt.Println("Error querying images:", err)
		return
	}
	defer imagesRow.Close()

	for imagesRow.Next() {
		var idFromRow int
		var createdTimeFromRow time.Time
		err = imagesRow.Scan(&idFromRow, &createdTimeFromRow)
		if err != nil {
			fmt.Println("Error scanning image row:", err)
			return
		}
		images = append(images, dbItem{idFromRow, createdTimeFromRow})
	}

	postsRow, err := db.Query("SELECT id, createdTime FROM POSTS WHERE createdTime >= $1 AND createdTime < $2", startTime, endTime)
	if err != nil {
		fmt.Println("Error querying posts:", err)
		return
	}
	defer postsRow.Close()

	for postsRow.Next() {
		var idFromRow int
		var createdTimeFromRow time.Time
		err = postsRow.Scan(&idFromRow, &createdTimeFromRow)
		if err != nil {
			fmt.Println("Error scanning post row:", err)
			return
		}
		posts = append(posts, dbItem{idFromRow, createdTimeFromRow})
	}
	return
}

func prioritizeData(items []dbItem) (priorityItems []dbItem) {
	priorityQueue := &PriorityQueue{}
	heap.Init(priorityQueue)

	for _, item := range items {
		heap.Push(priorityQueue, item)
	}

	for priorityQueue.Len() > 0 {
		priorityItems = append(priorityItems, heap.Pop(priorityQueue).(dbItem))
	}

	return
}

func main() {
	// Assume DB is already connected
	// db, err := sql.Open("postgres", "...")
	// defer db.Close()
	var db sql.DB

	timeNow := time.Now()
	images, posts := queryData(&db, timeNow.AddDate(0, 0, -1), timeNow)

	// Combine two and prioritize them
	all := append(images, posts...)
	all = prioritizeData(all)

	// Print prioritized items
	for _, item := range all {
		fmt.Printf("ID: %d, CreatedTime: %s\n", item.id, item.createdTime)
	}
}
