all: cmd gui

cmd: build/musicurator-cli.exe
	go build -o build/musicurator-cli.exe ./cmd/musicurator/

gui: build/musicurator-gui.exe
	go build -o build/musicurator-gui.exe ./gui/musicurator/

run-gui: gui
	./build/musicurator-gui.exe