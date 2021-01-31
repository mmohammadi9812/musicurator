clifiles := $(wildcard ./cmd/musicurator/*)
guifiles := $(wildcard ./gui/musicurator/*)
corefiles := $(wildcard ./core/*)

all: cmd gui

build/musicurator-cli.exe: $(corefiles) $(clifiles)
	go build -o build/musicurator-cli.exe ./cmd/musicurator/

build/musicurator-gui.exe: $(corefiles) $(guifiles)
	go build -o build/musicurator-gui.exe ./gui/musicurator/

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