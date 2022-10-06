package examplemodules;

import(
  "sync"
)

func recv_likes_and_sum(second_channel chan int, wg *sync.WaitGroup){ 
  likes := 0
  for{
    incoming_like, ok := <- second_channel 
    if !ok{
      print("Likes counted : ",likes)
      wg.Done()
      return
    }
    likes+=incoming_like
  }
}

func rapid_like_sender(channel chan int, wg *sync.WaitGroup){
  i:=0
  for i<100{
    channel<-1
    i++
  }
  wg.Done() //Does -1 to wg

}

func Likes_unbuf_channel(){
  second_channel := make(chan int) //Unbuffered
  wg := new(sync.WaitGroup)
  wg2 := new(sync.WaitGroup)
  wg.Add(1)
  go recv_likes_and_sum(second_channel,wg)
  wg2.Add(100)
  //We are basically sendinf 2 requests simulatneoulsy...the case that we studied before.
  i:=1
  for i<=10{
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    go rapid_like_sender(second_channel, wg2)
    i++
  }
  wg2.Wait()
  close(second_channel)
  wg.Wait()
}
