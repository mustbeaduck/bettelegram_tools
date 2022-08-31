package main

import "os"
import "io"
import "fmt"
import "net/url"
import "net/http"
import "github.com/Arman92/go-tdlib"


type JsonMap = map[string]interface{}
const gtmpath = "/tmp/bettelegramtemporaryvoicenotefile"

//5bit depth waveform of voice message
var gWaveform string = "ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"


//send a single text message
func sendMessage(client * tdlib.Client, chatId int64, text string) error {

    inputMsgTxt := tdlib.NewInputMessageText(tdlib.NewFormattedText(text, nil), true, true) 
    _, err := client.SendMessage(chatId, 0, 0, nil, nil, inputMsgTxt)
    
    return err
}

//send an audo file via voice message
func sendVoiceFile(client * tdlib.Client, chatId int64, path string) error {

    message := tdlib.UpdateData {
        "@type":"sendMessage",
        "chat_id": chatId,
        "input_message_content": JsonMap {
            "@type": "inputMessageVoiceNote",
            "duration": -1,
            "waveform": gWaveform,
            "voice_note": JsonMap {
                "@type": "inputFileLocal",
                "path": path,
            },
        },
    }

    result, err := client.SendAndCatch(message)

    if result.Data["@type"].(string) == "error" {
		return fmt.Errorf("error! code: %f msg: %s", result.Data["code"], result.Data["message"])
	}

    return err
}


func sendVoiceMessage(client *tdlib.Client, chatId int64, text, lang string) error {

    //load a tts from google translate api
    url := fmt.Sprintf("http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s", url.QueryEscape(text), lang)
    response, err := http.Get(url)

    if err != nil { return err }
    defer response.Body.Close()

    file, err := os.Create(gtmpath)
    if err != nil { return err }

    _, err = io.Copy(file, response.Body)
    file.Close()


    return sendVoiceFile(client, chatId, gtmpath)
}
