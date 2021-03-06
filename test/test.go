package main

import (
    "fmt"
    "github.com/hawell/workerpool"
    "time"
    "math/rand"
)

func getHandler() workerpool.JobHandler {
    x := rand.Int()
    return func(worker*workerpool.Worker, job workerpool.Job) {
    fmt.Println(worker.Id, " ", job, " rand = ", x)
    }
}

func main() {
    dispatcher := workerpool.NewDispatcher(10, 10)
    dispatcher.AddWorker(getHandler())
    dispatcher.AddWorker(getHandler())
    dispatcher.AddWorker(getHandler())
    dispatcher.Run()
    for i := 0; i < 100; i++ {
        job := fmt.Sprintf("job %d", i)
        dispatcher.Queue(job)
    }
    time.Sleep(time.Second * 3)
}