package main

import "fmt"
import "math"
import "github.com/Arman92/go-tdlib"


//print out a list of chats
func printChats(chats []*tdlib.Chat) {
    
    for _, chat := range chats {
        
        fmt.Printf("[%s | %d]: ", (*chat).Title, (*chat).ID)

        if (*chat).LastMessage != nil {
            
            msg := (*chat.LastMessage)
            if msg.Content.GetMessageContentEnum() == "messageText" {

                textMsg := msg.Content.(*tdlib.MessageText) 
                fmt.Println(textMsg.Text.Text)

            } else {
                fmt.Println(msg.Content.GetMessageContentEnum())
            }

        } else {
            fmt.Println("[No messages here]")
        }

        fmt.Println("__________")
    }
}


//get list of chats
func getChats(client *tdlib.Client, chatLimit int32) ([]*tdlib.Chat, error) {
    
    chatList := tdlib.NewChatListMain()

    chats, err := client.GetChats(chatList, tdlib.JSONInt64(math.MaxInt64), 0, chatLimit)
    if err != nil { return nil, err }

    res := make([]*tdlib.Chat, 0, chatLimit)
    
    for _, chatId := range chats.ChatIDs {
        
        chat, err := client.GetChat(chatId)
        if err != nil { return nil, err }
        
        res = append(res, chat)
    }

    return res, nil
}
