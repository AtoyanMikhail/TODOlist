# Используем официальный образ PostgreSQL
FROM postgres:15-alpine

# Копируем миграции в папку для автоматического выполнения
COPY migrations/*.sql /docker-entrypoint-initdb.d/

# Копируем скрипт-обертку для контроля выполнения
COPY entrypoint.sh /docker-entrypoint-initdb.d/

# Делаем скрипт исполняемым
RUN chmod +x /docker-entrypoint-initdb.d/entrypoint.sh

# Открываем порт PostgreSQL
EXPOSE 5432