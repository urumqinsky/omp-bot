.PHONY: run build clean

SRCS = cmd/bot/main.go

NAME = bot

run:
	go run $(SRCS)

build:
	go build -o $(NAME) $(SRCS)

clean:
	rm -f $(NAME)