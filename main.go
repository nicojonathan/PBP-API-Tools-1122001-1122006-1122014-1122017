package main

import (
	"fmt"
	"sync"
	"time"
	m "tools-api/models"
)

// TODO: cari kegunaan dari sync.WaitGroup

func main() {
	var wg sync.WaitGroup

	// Contoh dummy data tasks list
	tasks := []m.Task{
		{ID: 1, Title: "Meeting", DueDate: time.Now().Add(5 * time.Second), Email: "user@example.com", Notified: false},
		{ID: 2, Title: "Project Deadline", DueDate: time.Now().Add(10 * time.Second), Email: "user@example.com", Notified: false},
		{ID: 3, Title: "Presentation", DueDate: time.Now().Add(15 * time.Second), Email: "user@example.com", Notified: false},
	}

	// Bikin GoRoutine untuk setiap task
	for _, task := range tasks {
		wg.Add(1)
		go func(t m.Task) {
			defer wg.Done()
			for {
				// Ini kalau kita mau ngirim emailnya setelah si tasknya udah lewat dari due date
				if time.Now().After(t.DueDate) && !t.Notified {
					fmt.Printf("Task %s is due!\n", t.Title)
					// Kirim emailnya disini
					// sendEmail(string recEmail, string subject, string content)
					// Contoh: sendEmail(t.Email, "Task Due", "Your task "+t.Title+" is due now!")
					t.Notified = true
					break
				}
				time.Sleep(1 * time.Second)
				// Ini refresh rate dari checking si tasknya
				// Jadi kalau diatas dia akan ngecheck task tiap 1 detik
			}
		}(task)
	}

	wg.Wait()
}
