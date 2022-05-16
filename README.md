# RSS AggreBot - RSS Aggregator Telegram Bot

## Что это?
**RSS AggreBot** - бот для Telegram, формирующий персональную новостную ленту на
основе различных RSS источников.


## Как им пользоваться?
### Простой способ
Запустить бот в Telegram и следовать инструкциям.

Через интерфейс можно настроить свою ленту - добавить/удалить источники и 
настроить [RegExp](https://ru.wikipedia.org/wiki/Регулярные_выражения) фильтр.

### Для продвинутых пользователей
С приложением можно общаться через HTTP API - запросами gRPC или REST 
_(прокси к gRPC)_.


## Как устроен внутри?
### Схема работы приложения
```mermaid
flowchart LR

subgraph External
    DIRECT{{Direct Requests}}
    TGUSER{{Telegram User}}
    RSS{{RSS/Atom Source}}
end
subgraph Docker Compose Application
        subgraph BOTSERVICE [Bot UI Service]
            BOT>Telegram<br/>chat handling]
        end
        subgraph API [Backend API Service]
          CONFIGURER((Sources<br/>Configurator))
        end
    subgraph Database [Database Container]
        DB[(PostgreSQL)]
    end
    subgraph Courier Service
        READER[\Reader/]
        SENDER[/Sender\]
    end
end

DIRECT  ---|gRPC / REST| CONFIGURER
TGUSER  ===|config<br/>menu| BOT ---|gRPC| CONFIGURER---|SQL| DB
RSS -.->|feeds| READER

DB -->|get<br/>active sources| READER 
READER -->|Go<br/>queue| SENDER
SENDER ---|update<br/>entries log| DB
TGUSER ===|newsfeed<br/>messages| SENDER

```
Приложение бота разворачивается через **Docker Compose** на четырёх 
контейнерах:
1. Bot UI Service
2. Backend API Service
3. Courier Service
4. Database Container

#### Bot UI Service
Обрабатывает сообщения/команды пользователей, транслируя запросы к
Backend API через gRPC.

#### Backend API Service
Имеет открытые HTTP порты, принимающие запросы через gRPC или REST _(прокси к 
gRPC)_.
Работает с базой данных, где хранит конфигурации пользователя (источники).

#### Courier Service
Состоит из двух компонентов:
- **Reader** - читает все активные источники из БД и скачивает их записи, 
  которые передаёт для обработки Sender.
- **Sender** - рассылает новые записи пользователю (с учётом пользовательского 
  фильтра), обновляя в БД информацию об отправленных (лог хешей записей).

#### Database Container
База данных PostgreSQL.

## Как запустить самостоятельно?

### Локальный запуск
**Требования:**
- **Go 1.18+**
- **PostgreSQL 14+**
- [**goose**](https://github.com/pressly/goose)

**Установка:**
1) В файле `.env_local` указать параметры подключения к Базе 
Данных и Токен бота.
2) Выполнить `make goose` для создания таблиц БД.

**Запуск**.
Выполнить в отдельных оболочках:
1) `make run_backend`
2) `make run_bot_ui`
3) `make run_courier`

### Запуск в Docker Compose
**Требования:**
- **UNIX**-совместимая ОС
- **Docker Compose**

Параметры создаваемого приложения Docker Compose задаются через файл
`/deploy/.env`.

В нём необходимо как минимум указать Токен бота - `TG_TOKEN`.

Через переменную `DB_DATA_HOST_DIR` можно изменить место хранения БД.

**Управление.** Перейти в директорию `/deploy/`:
- `make` - сборка и запуск
- `make stop` - остановка
- `make clean` - удаление контейнеров
- `make fclean` - полное удаление, включая файлы БД

---
_Артем **nGragas** Корников. Учебный проект для Ozon Route 256._