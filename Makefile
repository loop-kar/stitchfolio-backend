build:
	go build -tags netgo -ldflags '-s -w' -o app/app


#Run Scripts
run:
	go run main.go --configFile config/prod.yaml

run-dev:
	go run main.go --configFile config/dev.yaml

run-qa:
	go run main.go --configFile config/qa.yaml


launch:	
	./app/app --configFile=config/prod.yaml

launch-qa:	
	./app/app --configFile=config/qa.yaml


serve:
	make build
	make launch

#Migrate
migrate:
	go run main.go --configFile config/prod.yaml --migrate true

	
migrate-qa:
	go run main.go --configFile config/qa.yaml --migrate true

	
migrate-dev:
	go run main.go --configFile config/dev.yaml --migrate true


#Dev Tools
wire:
	cd internal/di && wire

swagger:
	swag fmt
	swag init --parseDependency --parseInternal	