include .env

run up:
	./script/app.sh run

stop down:
	./script/app.sh stop

deps:
	docker compose -f ./tools/compose-tools.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-openobserve/api" \
		-w "/go/src/github.com/hyorimitsu/hello-openobserve/api" \
		go-mod

logs-%:
	./script/app.sh logs $*

dashboard:
	./script/app.sh dashboard

destroy:
	./script/app.sh destroy
