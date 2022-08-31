package main

import "os"
import "fmt"
import "bufio"
import "strconv"


var gLang string  = "uk"
var gChatLimit int32 = 15


func main() {

    //auth
    client := getClient() 
    defer client.Close()
    authLoop(client)
   
    var userInput string
    stdin := bufio.NewReader(os.Stdin)

    run := true

    for run {

        chats, err := getChats(client, gChatLimit)
        if err != nil { panic(err) }

        printChats(chats)

        fmt.Print("Select a chat: ")
        userInput, _ = stdin.ReadString('\n')
        userInput = userInput[:len(userInput)-1]

        chatId, err := strconv.ParseInt(userInput, 10, 64)
        
        if err != nil {
            fmt.Fprintf(os.Stderr, "failed to parse '%s' to a number\n", userInput)
        } else {

            fmt.Println("...chat selected...")
            chatLoop := true

            for chatLoop {
                
                fmt.Print("?: ")
                userInput, _ = stdin.ReadString('\n')
                userInput = userInput[:len(userInput)-1]

                //if its a command
                if len(userInput) > 1 {
                    if userInput[0] == '\\' {
                        if len(userInput) > 3 {
                            if userInput[1] == 'v' {
                                sendVoiceMessage(client, chatId, userInput[3:], gLang)
                            } else if userInput[1] == 's' {
                                sendVoiceFile(client, chatId, userInput[3:]) 
                            }
                        } else if userInput[1] == 'q' {
                            chatLoop = false
                        } else if userInput[1] == 'e' { 
                            chatLoop = false 
                            run = false
                        } else {
                            fmt.Println("uknown command")
                        }
                    } else {
                        sendMessage(client, chatId, userInput)
                    }
                } else {
                    sendMessage(client, chatId, userInput)
                } 

            }
        }
    }
}
