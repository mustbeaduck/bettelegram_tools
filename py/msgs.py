from telegram.client import Telegram
import requests as req


def send_file_voice_note(tg, chat_id, path):

    data = {
        "@type":"sendMessage",
        "chat_id": chat_id,
        "input_message_content": {
            "@type": "inputMessageVoiceNote",
            "duration": -1,
            #wont parse json request unless the padding is correct lmao
            "waveform": ("n"*136),
            "voice_note": {
                "@type": "inputFileLocal",
                "path": path,
            },
        },
    }
     
    try:

        tg._send_data(data, block=True)

    except Exception as e:

        print(f"!\t!\tfailed to send {path} via voice message \t!\t!")
        print(e)


def send_ttm_voice_note(tg, chat_id, text, lang):
    #disguising myself as a windows user :)
    headers = {"User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64)"}

    text = req.utils.quote(text) #URIEncode
    resp = req.get("http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q="+text+"&tl="+ lang, headers = headers)

    if resp.status_code == 200:

        path = "/tmp/_bettelegramvoicenote.mp3"

        # i wish tdlib had a way to send loaded file from memory yet i didnt find one
        file = open(path, "wb")
        file.write(resp.content)
        file.close()

        send_file_voice_note(tg, chat_id, path)

    else:
        print("failed to acces http://translate.google.com/translate_tts")
        print(f"http response code: {resp.status_code}")

