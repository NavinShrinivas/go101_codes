package examplemodules

import(
  "fmt"
)

func Simple_buf_channel(){
  // first_channel := make(chan string) //This alone will cause runtime errors! As capacity is 0
  first_channel := make(chan string,1)
  fmt.Println("Before queuing, len :",len(first_channel)) //Shows number of queued messages in the buffer
  fmt.Println("Before queuing, cap :",cap(first_channel)) //Shows the total capacity of buffer
  first_channel <- "Hello, World!"
  fmt.Println("After queuing one string, len :",len(first_channel))
  fmt.Println("After queuing one string, cap :",cap(first_channel))
  // first_channel <- "Second message :(" //Will cause runtime
  our_message := <- first_channel 
  //We have to use := for "our_message" as the variables needs to figure what it stored by itself, we can also declare its type manually
  fmt.Println(our_message)

  first_channel <- "Second message :(" 
  var second_message string
  second_message = <- first_channel 
  fmt.Println(second_message)
}
