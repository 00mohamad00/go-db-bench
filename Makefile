SRCS = $(patsubst ./%,%,$(shell find . -name "*.go" -not -path "*vendor*" -not -path "*.pb.go"))

bench: $(SRCS)
	go build -o $@ -ldflags="$(LD_FLAGS)" ./cmd

up-postgres:
	docker run --name postgres --env=POSTGRES_USER=test --env=POSTGRES_PASSWORD=test --env=POSTGRES_DB=test -p 5432:5432 -d postgres

down-postgres:
	docker rm -f postgres

up-mongo:
	docker run --name mongo -p 27018:27017 -d docker.repos.balad.ir/mongo

down-mongo:
	docker rm -f mongo