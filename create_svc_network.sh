#!/bin/sh
# Имя сети и подсеть /64
NET_NAME="svc-net"
SUBNET_IPV6="fd98:2dd6:8f48:1d99::/64"
MTU_VALUE="65535"

echo "=== Проверка Docker сети: ${NET_NAME} ==="

# Проверяем, существует ли сеть с таким именем
if docker network inspect "$NET_NAME" >/dev/null 2>&1; then
    echo "Предупреждение: Сеть '${NET_NAME}' уже существует."
    echo "Удаляем старую сеть..."
    
    # Force удаление (если к ней привязаны остановленные контейнеры)
    docker network rm "$NET_NAME" >/dev/null 2>&1
    
    if [ $? -ne 0 ]; then
        echo "Ошибка: Не удалось удалить сеть. Возможно, она используется запущенными контейнерами!"
        echo "Остановите контейнеры и попробуйте снова."
        exit 1
    fi
    echo "Старая сеть успешно удалена."
else
    echo "Старой сети с именем '${NET_NAME}' не обнаружено. Отлично."
fi

echo "Создаем новую IPv6-only сеть..."

# Команда создания сети со всеми твоими параметрами
docker network create \
  --ipv6 \
  --subnet="$SUBNET_IPV6" \
  --opt com.docker.network.driver.mtu="$MTU_VALUE" \
  "$NET_NAME"

# Проверка статуса создания
if [ $? -eq 0 ]; then
    echo "=== Успех! Сеть '${NET_NAME}' создана ==="
    echo "Подсеть: $SUBNET_IPV6"
    echo "MTU: $MTU_VALUE"
else
    echo "Ошибка: Что-то пошло не так при создании сети."
    exit 1
fi
