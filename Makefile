dev:
	@$(MAKE) -j server web

server:
	go run main.go

web:
	cd ~/Dev/nitro-react && yarn start --host