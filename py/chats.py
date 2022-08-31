from telegram.client import Telegram

def print_chats(tg, chat_IDs):

    for chat_id in chat_IDs:

        chat = tg.get_chat(chat_id)
        chat.wait()

        chat = chat.update

        print(f"[{chat['title']} | {chat['id']}] : ", end = "")

        if "last_message" in chat:

            msg = chat["last_message"]
            if msg["content"]["@type"] == "messageText":
                print(msg["content"]["text"]["text"])

            else:
                print(msg["content"]["@type"])

        else:
            print("[No messages here]")

        print("__________\n")

