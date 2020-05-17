package main
import (
  "github.com/apache/openwhisk-client-go/whisk"
  "fmt"
  "net/http"
)
func do_invoke(done chan bool) {
  client, err := whisk.NewClient(http.DefaultClient, nil)
  if err != nil {
      done <- false
  }
  m := make(map[string]string)
  m["name"] = "anh dinh"
  if res, _, err := client.Actions.Invoke("hello6", m, false, false); err == nil {
    done <- true
    fmt.Printf("Result %v\n", res["activationId"])
  } else {
    fmt.Println(err)
    done <- false
  }

}

func main() {
  n := 2
  done := make(chan bool, n)
  for i:=0; i<n; i++ {
    go do_invoke(done)
  }

  success := 0
  for n !=0 {
    if x := <- done; x {
      success+=1
    }
    n--
  }
  fmt.Printf("Done! #successful requests %v\n", success)
}
