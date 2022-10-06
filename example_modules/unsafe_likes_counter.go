package examplemodules

import(
  "sync"
)

func rapid_likes_sender2(likes *int, wg *sync.WaitGroup){
  defer wg.Done()
  i := 0
  for i<100{
    *likes+=1
    i++
  }
}

func Unsafe_likes_counter(){
  wg := new(sync.WaitGroup)
  likes := 0;
  wg.Add(100)
  i:=1
  for i<=10{
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    go rapid_likes_sender2(&likes, wg)
    i++
  }
  wg.Wait()
  print("Likes counted :",likes)
}
