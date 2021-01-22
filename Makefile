BINDIR=bin

all: android

android:
	gomobile bind -v -x -o $(BINDIR)/dss.aar -target=android github.com/giantliao/androidlib

clean:
	gomobile clean
	rm $(BINDIR)/*
