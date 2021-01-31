clifiles := $(wildcard ./cmd/musicurator/*)
guifiles := $(wildcard ./gui/musicurator/*)
corefiles := $(wildcard ./core/*)

all: cmd gui

build:
	mkdir -p build

build/musicurator-cli.exe: $(corefiles) $(clifiles) build
	go build -o build/musicurator-cli.exe ./cmd/musicurator/

build/musicurator-gui.exe: $(corefiles) $(guifiles) build
	go build -ldflags "-H=windowsgui" -o build/musicurator-gui.exe ./gui/musicurator/

.PHONY: cmd
cmd: build/musicurator-cli.exe

.PHONY: gui
gui: build/musicurator-gui.exe

run-gui: gui
	./build/musicurator-gui.exe

test-cmd: cmd
	./build/musicurator-cli.exe -n -s "C:\Users\Mohammad\Documents\Code\GO\musicurator\testdata" -d "C:\Users\Mohammad\Documents\Code\GO\musicurator\testoutput" -t '$$artist - $$title.$$ext'

.PHONY: testdata
testdata:
	rm ./testdata/dst/*
	cp -t ./testdata/src/ ./testdata/src.bak/*

.PHONY: package-windows
package-windows: gui
	fyne package -os windows -icon ./build/cassette.png -release -sourceDir ./gui/musicurator/ -executable ./build/musicurator-gui.exe