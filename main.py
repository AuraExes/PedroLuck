import os
import random
import telebot

TOKEN = os.getenv("BOT_TOKEN")
bot = telebot.TeleBot(TOKEN)

@bot.message_handler(commands=['luck'])
def send_luck(message):
    luck_score = random.randint(0, 100)
    
    if luck_score == 100:
        res = "Ты выбил всё лучшее. Удача повышена на: 100%"
    elif luck_score > 50:
        res = f"✅ Повезло. Удача повышена на: {luck_score}%"
    elif luck_score > 10:
        res = f"Неплохо. Удача повышена на: {luck_score}%"
    else:
        res = f"❌ Сегодня не твой день. Удача повышена на: {luck_score}%"
        
    bot.reply_to(message, res)

if __name__ == "__main__":
    bot.infinity_polling()
