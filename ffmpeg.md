# FFmpeg
A complete, cross-platform solution to record, convert and stream audio and video.  
[Documentation](https://www.ffmpeg.org)  

## Примеры команд
ffmpeg [параметры входного кодирования] -i 1.mov [параметры выходного кодирования] 1.mp4  
ffmpeg -i 1.mov -preset veryslow 1.mp4  
ffmpeg -i 1.mov -preset slow 1.mkv  

## Наиболее популярные контейнеры
mp4  
avi  
mpeg  
mov  
flv  
webm  
mkv   
3gp  

## Вспомогательные инструменты
FFplay - a simple media player based on SDL and the FFmpeg libraries.  
FFprobe - a simple multimedia stream analyzer.  

Видеофайл - контейнер, в котором содержится видео-поток и аудио-поток, закодированные в определенный стандарт. Например, контейнер mp4, в нем с помощью кодека H264 закодировано видео и аудио.

## Указание вида параметров, на которое дествует настройка
-параметр:v - для видео  
-параметр:a - для аудио  

## Настройки
кодек:  
-codec:v libx264 -preset:v slow  
  
профиль:  
-profile:v high  
  
пресет:  
-preset veryslow  
все пресеты:  
ultrafast - размер наибольший, качество плохое, скорость быстрая  
superfast  
veryfast  
faster  
fast  
medium - gefault  
slow  
slower  
veryslow - размер наименьший, качество хорошее, скорость медленная  
  
постоянное визуальное качество (crf)  
в случае с НЕ H264 - это битрейт (размер видеопотока в секунду)  
а в случае H264 - параметр квантования (18 - хорошо, но тяжело, 28 - плохо но легко)  
-crf:v 20   
  
группа изображений (gop)  
настройки ключевых кадров и т.п.  
