#!/bin/bash

TARGET_DIR="./brabus_linux"
ARCHIVE_NAME="brabus_package_$(date +%Y%m%d_%H%M%S).tar.gz"

# Создаем целевую директорию
mkdir -p "$TARGET_DIR"

# Функция для копирования с заменой debug значения
safe_copy_config() {
    local src=$1
    local dest=$2

    if [ -f "$dest" ]; then
        echo "Удаляю старую версию $dest"
        rm -f "$dest"
    fi

    if [ -f "$src" ]; then
        sed 's/debug: true/debug: false/g' "$src" > "$dest"
        echo "Скопировано и изменено: $src -> $dest (debug установлен в false)"
    else
        echo "Ошибка: исходный файл $src не найден"
        exit 1
    fi
}

# Функция для обычного копирования
safe_copy() {
    local src=$1
    local dest=$2

    if [ -f "$dest" ]; then
        echo "Удаляю старую версию $dest"
        rm -f "$dest"
    fi

    if [ -f "$src" ]; then
        cp "$src" "$dest"
        echo "Скопировано: $src -> $dest"
    else
        echo "Ошибка: исходный файл $src не найден"
        exit 1
    fi
}

# Копируем файлы
safe_copy "./brabus/bin/brabus" "$TARGET_DIR/brabus"
safe_copy "./brabus/bin/banana" "$TARGET_DIR/banana"
safe_copy "./brabus/docker-compose.yml" "$TARGET_DIR/docker-compose.yml"
safe_copy_config "./brabus/config.yaml" "$TARGET_DIR/config.yaml"

# Создаем архив
echo "Создаю архив $ARCHIVE_NAME..."
tar -czvf "$ARCHIVE_NAME" "$TARGET_DIR"

# Проверяем что архив создан
if [ -f "$ARCHIVE_NAME" ]; then
    echo "Готово! Создан архив:"
    ls -lh "$ARCHIVE_NAME"
else
    echo "Ошибка при создании архива"
    exit 1
fi
