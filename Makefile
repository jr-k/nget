.PHONY: dev dev-build dev-run

dev: dev-run

dev-build:
	docker build -t nget .

dev-run: dev-build
	docker run -it --rm -p 33000-33100:33000-33100 -v .:/data nget

dev-test: dev-build
	docker run -it --rm -p 33000-33100:33000-33100 -v .:/data nget sh -c "./nget README.md ; ./nget README.md ; ./nget Dockerfile ; ./nget ls ; ./nget kill 1 ; ./nget ls ; exec sh"