Так как у нашего контейнера нету почти ничего изначально в образе, приходиться делать так.
Меняем DNS
```
echo "nameserver 127.0.0.1" > /etc/resolv.conf
```
Перезаписываем nsswitch
```
sed -i 's/^hosts:.*/hosts: files/' /etc/nsswitch.conf
```