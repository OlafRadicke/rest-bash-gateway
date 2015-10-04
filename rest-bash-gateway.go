package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "os/exec"
  "strings"
  "encoding/json"
  "io/ioutil"
)

var CammandMatch map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
  ressource := r.URL.Path[1:len(r.URL.Path)]
  result := "command not supported"
  if command, ok := CammandMatch[ressource]; ok {
//    command := string(CammandMatch[ressource])
    result = rest2bash( command )
  }
  fmt.Fprintf(w, "%s", result)
}

func main() {
  fmt.Printf("Search config.json ...\n")
  readConfig()
  http.HandleFunc("/", handler)
  fmt.Printf("Listen on http://localhost%s/ \n", string(CammandMatch["restbashgateway_port"]))
  log.Fatal(http.ListenAndServe(CammandMatch["restbashgateway_port"], nil))
//  http.ListenAndServe(CammandMatch["restbashgateway_port"], nil)
}

func readConfig(){
  file, e := ioutil.ReadFile("./config.json")
  if e != nil {
      fmt.Printf("File error: %v\n", e)
      os.Exit(1)
  }
  json.Unmarshal(file, &CammandMatch)
}

func rest2bash(command string) string {
  splitCommand := strings.Split(command, " ")
  commandHead := splitCommand[0]
  params := splitCommand[1:len(splitCommand)]
  out, err := exec.Command( commandHead, params... ).Output()
  if err != nil {
    log.Fatal(err)
    }
      return string(out)
}
