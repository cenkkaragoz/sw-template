install:
	go install -v

validate:
	swagger validate ./swagger/swagger.yml

gen: validate 
	swagger generate server \
		--target=./swagger \
		--spec=./swagger/swagger.yml \
		--exclude-main \
		--name=hello

.PHONY: install gen validate