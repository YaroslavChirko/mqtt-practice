package main

import (
    "fmt"
    "encoding/json"
    "crypto/sha1"
    "os"
    "bufio"
    "encoding/hex"
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

type payload struct{
    Secret string `json:"secret"` 
    Wisdom string `json:"wisdom"` 
    Team string `json:"team"` 
}

func main() {
    opts := mqtt.NewClientOptions().
    
        AddBroker(fmt.Sprintf("tcp://%s:1883", "localhost")) //%s
    client := mqtt.NewClient(opts)
    client.Connect().Wait()
    client.Subscribe("/test/inception", 0,
       func(client mqtt.Client, message mqtt.Message) {
       var dat payload
        fmt.Printf("* [%s] %s\n", message.Topic(), string(message.Payload()))
        if err := json.Unmarshal(message.Payload(), &dat);
           err != nil {
            return
        }
secret,_ := json.Marshal(dat.Secret)

dst := make([]byte, hex.EncodedLen(len(secret)))
	hex.Encode(dst, secret)
    data, _ := json.Marshal(&payload{Secret : hex.EncodeToString(sha1.New().Sum([]byte(dat.Secret + "mmys"))), Wisdom : dat.Wisdom,  Team : "mmys"})
   client.Publish("/test/result", 0, false, data)
   err := json.Unmarshal(data, &dat)
if err!=nil {
panic(err)
}
fmt.Println(dat)
   if err := json.Unmarshal(data, client);
           err != nil {
            return
        }
 fmt.Println(string(data))
})
input := bufio.NewScanner(os.Stdin)
    input.Scan()

}

