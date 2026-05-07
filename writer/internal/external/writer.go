package external

import (
	"context"
	"fmt"
	"io"
	h "net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Writer struct {
	url      string
	cooldown int
}

func NewWriter(url string, cooldown int) *Writer {
	return &Writer{
		url:      url,
		cooldown: cooldown,
	}
}

func (w *Writer) Ask() error {
	file, err := os.Create("/data/system_info.log")
	if err != nil {
		return err
	}
	defer file.Close()

	client := h.Client{}
	req, err := h.NewRequest("GET", w.url, nil)
	if err != nil {
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(time.Second * time.Duration(w.cooldown))
	defer ticker.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				resp, _ := client.Do(req)
				defer resp.Body.Close()
				body := resp.Body
				data, _ := io.ReadAll(body)
				file.Write(data)
				fmt.Println(string(data))
			case <-ctx.Done():
				return
			}
		}
	}()
	<-stop
	cancel()
	wg.Wait()
	return nil
}
