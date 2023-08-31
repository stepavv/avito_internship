#!/bin/bash

go_path="/home/steven/internship/project/app/app.go"

# Проверка наличия исполняемого файла Go
if [ -f "$go_path" ]; then
  # Команда для компиляции и запуска файла Go
  go run "$go_path"
else
  echo "Ошибка: Файл $go_path не найден"
fi
