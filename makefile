
# ----------------------------
.PHONY: stop
stop:
	docker-compose stop


# ----------------------------
.PHONY: removeall
removeall:
	docker-compose stop
	docker system prune
	sudo rm -rf db/engine/mysql/var_lib_mysql/
	sudo rm -rf db/engine/mysql/data/
	sudo rm -rf db/engine/postgres/data
	sudo rm -rf db/tool/phpmyadmin/sessions/
	sudo rm -rf db/tool/pgadmin/root/


# ----------------------------
.PHONY: build
build:
	docker-compose build
	# docker-compose build --no-cache


# ----------------------------
.PHONY: up
up:
	docker-compose up

# ----------------------------
.PHONY: reup
reup: build up

# ----------------------------
.PHONY: resetall
resetall: removeall build up
