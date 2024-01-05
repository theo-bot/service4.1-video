# Check to see if we can use ash in Alpine images of default to BASH
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATHH)),/bin/ash,/bin/bash)

tidy:
	go mod tidy
	go mod vendor