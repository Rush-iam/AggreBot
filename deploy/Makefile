#     _                        ____        _
#    / \   __ _  __ _ _ __ ___| __ )  ___ | |_
#   / _ \ / _` |/ _` | '__/ _ \  _ \ / _ \| __|
#  / ___ \ (_| | (_| | | |  __/ |_) | (_) | |_
# /_/   \_\__, |\__, |_|  \___|____/ \___/ \__|
#         |___/ |___/      by nGragas/Rush-iam
include .env

NAME = AggreBot

all: up

up: | $(DB_DATA_HOST_DIR)
	docker compose -p "$(NAME)" up --build
$(DB_DATA_HOST_DIR):
	mkdir -p $@
stop:
	docker compose -p "$(NAME)" stop -t 5
clean:
	docker compose -p "$(NAME)" rm -f -s
fclean:
	docker compose -p "$(NAME)" down --rmi all -v -t 3
	sudo rm -rf $(DB_DATA_HOST_DIR)
.PHONY: up stop clean fclean
