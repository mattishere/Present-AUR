CMD=go
BINARY=presaur
PREFIX ?= /usr

install:
	@$(CMD) build -o $(BINARY) main.go
	@install -Dm755 presaur $(DESTDIR)$(PREFIX)/bin/$(BINARY)

uninstall:
	@rm -rf $(DESTDIR)$(PREFIX)/bin/$(BINARY)
