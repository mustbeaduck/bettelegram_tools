from telegram.client import Telegram
from chats import print_chats
from msgs import *

#tts language
LANG = "uk"

#phone_number = ""
phone_number = input("phone: ")

"""
                AUTH
"""

tg = Telegram(
    api_id=11280565,
    api_hash='5561e9f775bae9639412059f1cd4345d',
    phone = phone_number,
    use_message_database = False,
    database_encryption_key='',
)

tg.login()

run = True

while run:

    chat_list = tg.get_chats()
    chat_list.wait()

    print_chats(tg, chat_list.update["chat_ids"])
    chat_id = 0

    try:
        chat_id = int(input("select a chat: "))
    except ValueError:
        print("something is wrong with your integer, buddy")

    if chat_id in chat_list.update["chat_ids"]:
        
        chat_loop = True
        while chat_loop:

            user_inp = input("?: ")

            if len(user_inp) > 1:
                #process a command
                if user_inp[0] == "\\":
                    
                    if len(user_inp) > 3:

                        if user_inp[1] == "v":
                            send_ttm_voice_note(tg, chat_id, user_inp[3:], LANG)

                        elif user_inp[1] == "s":
                            send_file_voice_note(tg, chat_id, user_inp[3:])

                    #select another chat
                    elif user_inp[1] == "q": chat_loop = False
                    
                    #exit
                    if user_inp[1] == "e": 

                        chat_loop = False
                        run = False

                    else:
                        print("uknown command")
                else:
                    tg.send_message(chat_id = chat_id, text = user_inp)
            else:
                tg.send_message(chat_id = chat_id, text = user_inp)

    else:
        print("cant seem to find specified chat")


tg.stop()
