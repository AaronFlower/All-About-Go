package thumbnail

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

var files = []string{"timg.jpeg", "timg2.jpeg"}

// generates thumbnails without routines
func BenchmarkMakeFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			if _, err := ImageFile(f); err != nil {
				log.Println(err)
			}
		}
	}
}

// generates thumbnails with routines
// but this is not corrected
// 在这里开启的进程没有完成就结束了，因为外部没有等待他们。
// This makeThumbnails returns before it has finished doing what it was supposed to do.
// It starts all the goroutines, one per file, but doesn't wait for them to finish.
func BenchmarkMakeFilesWithGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			go ImageFile(f) // ignoring errors
		}
	}
}

// makes thumbnails of the specified files in parallel
// 使用匿名函数来启动 goroutine, 我们就可以获取文件处理完成的状态。注意匿名函数的传递值。
// 但是这个函数没捕获函数的返回值，如查想获取函数的返回值怎么办那？看下面的函数
func BenchmarkMakeWithChannels(b *testing.B) {
	ch := make(chan struct{})
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			go func(f string) {
				ImageFile(f)
				ch <- struct{}{}
			}(f)
		}
		for range files {
			<-ch
		}
	}
}

// This function has a subtle bug. When it encounters the first non-nil error, it returns the error
// to the caller, leaving no goroutine draining the errors channel. Each remaining worker goroutine
// will block forever when it tries to send a value on that channel, and will never terminate. This
// situation, a goroutine leak, may cause the whole program to get stuck or to run out of memory.
func BenchmarkMakeWithChannelsAndReturns(b *testing.B) {
	ch := make(chan error)
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			go func(f string) {
				_, err := ImageFile(f)
				ch <- err
			}(f)
		}
		for range files {
			if err := <-ch; err != nil {
				log.Print(err)
				return // Note: 在这里直接 return 的话会造成 goroutine leak!
			}
		}
	}
}

// 解决上面 goroutine leak 的方法是使用一个 buffered channel。The simplest solution is to use a
// buffered channel with sufficient capacity that no worker goroutine will block when it sends a
// message.
// 这里写的 benchmark 的方法，把没有返回值，实际的函数是应该返回的。

func BenchmarkMake5(b *testing.B) {
	type item struct {
		file string
		err  error
	}

	type ret struct {
		files []string
		err   error
	}
	var r ret
	var outfiles []string
	ch := make(chan item, len(files))
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			go func(f string) {
				outfile, err := ImageFile(f)
				ch <- item{outfile, err}
			}(f)
		}

		for range files {
			it := <-ch
			if it.err != nil {
				r = ret{nil, it.err}
				// return r // 这里写的 benchmark 的方法，把没有返回值，实际的函数是应该返回的。
			}
			outfiles = append(outfiles, it.file)
		}
	}

	r = ret{outfiles, nil}
	fmt.Println(r)
	// return r // 这里写的 benchmark 的方法，把没有返回值，实际的函数是应该返回的。
}

// 最后一个例子，统计生成的 thumbnails 的所有文件的大小。
func BenchmarkTotalSize(b *testing.B) {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		for _, f := range files {
			wg.Add(1)
			// start worker
			go func(f string) {
				defer wg.Done()
				thumbfile, err := ImageFile(f)
				if err != nil {
					log.Fatal(err)
					return
				}
				info, _ := os.Stat(thumbfile)
				sizes <- info.Size()
			}(f)
		}
	}

	// wait and close
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Println("total size: ", total)
	// return total
}
