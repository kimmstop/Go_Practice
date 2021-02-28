CC=go build
TARGET=wiki.out
OBJS=Handler Edit Save View Page Template

all: $(TARGET)

$(TARGET): $(OBJS)


Handler: Handler/makeHandler.go Handler/go.mod
	go build Handler/makeHandler.go

Edit: Edit/editHandler.go Edit/go.mod
	go build Edit/editHandler.go

Save: Save/saveHandler.go Save/savePage.go Save/go.mod
	go build Save/saveHandler.go Save/savePage.go

View: View/viewHandler.go View/go.mod
	go build View/viewHandler.go

Page: Page/Page.go Page/go.mod
	go build Page/Page.go

Template: Template/renderTemplate.go Template/go.mod
	go build Template/renderTemplate.go

wiki.out: wiki.go go.mod
	go build wiki.go



clean:
	rm -rf wiki
	
