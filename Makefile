.PHONY: help
help:
	@echo "List of commands:"
	@echo " - dev    development runner"
	@echo " - css    build tailwind stylesheet"

.PHONY: dev
dev: css
	go run .

.PHONY: css
css:
	npm run css
