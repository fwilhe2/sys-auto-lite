all:
	docker build -t sys-auto-lite .
	docker run -it --rm sys-auto-lite bash