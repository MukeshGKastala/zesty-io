PORT := 9000
BINARY := autocomplete

#Corpuses
SHAKESPEARECOMPLETE := shakespeare-complete.txt
SHAKESPEARESAMPLE := shakespeare-sample.txt

#Directories
TARGETDIR := bin
DATADIR := data

run:
	@go build -o $(TARGETDIR)/$(BINARY) main.go
	@./$(TARGETDIR)/$(BINARY) -p $(PORT) -f $(DATADIR)/$(SHAKESPEARECOMPLETE)

clean:
	@$(RM) -rf $(TARGETDIR)

test:
	@go test -race -v -cover ./...

examples:
	@chmod +x examples.sh
	@./examples.sh $(PORT)

.PHONY: run clean test examples