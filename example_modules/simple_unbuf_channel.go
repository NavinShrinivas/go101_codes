package examplemodules

import(
  "sync"
  "fmt"
)

func unbuf_sender(channel chan string, wg *sync.WaitGroup){
  defer wg.Done()
  defer close(channel)
  i := 0
  for i<5{
    channel <- "Hello World "+fmt.Sprint(i)
    i++
  }
}

func unbuf_reader(channel chan string, wg *sync.WaitGroup){
  defer wg.Done()
  for{
    res, ok := <- channel
    if !ok{
      fmt.Println("We done, adios")
      return
    }
    fmt.Println(res)
  }
  
}

func Simple_unbuf_channel(){
  unbuf_channels := make(chan string)
  wg := new(sync.WaitGroup)
  wg.Add(1)
  go unbuf_sender(unbuf_channels, wg)
  wg.Add(1)
  go unbuf_reader(unbuf_channels, wg) //Code will panic if this receiver isnt active
  wg.Wait()
}
