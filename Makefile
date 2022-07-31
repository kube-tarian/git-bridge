start-docker-compose:
	env GITLAB_HOME=./gitlab-repo-server docker compose up -d --no-recreate

stop-docker-compose:
	docker compose down -v

