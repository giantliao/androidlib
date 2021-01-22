BINDIR=bin

#.PHONY: pbs

all: android
#
#pbs:
#	cd pbs/ && $(MAKE)
#

tp:=./

# test:
# 	go build  -ldflags '-w -s' -o $(BINDIR)/ctest mac/*.go
# m:
# 	CGO_CFLAGS=-mmacosx-version-min=10.11 \
# 	CGO_LDFLAGS=-mmacosx-version-min=10.11 \
# 	GOARCH=amd64 GOOS=darwin go build  --buildmode=c-archive -o $(BINDIR)/dss.a mac/*.go
# 	cp mac/callback.h $(BINDIR)/
android:
	gomobile bind -v -x -o $(BINDIR)/dss.aar -target=android github.com/giantliao/androidlib
# i:
# 	gomobile bind -v -o $(BINDIR)/iosLib.framework -target=ios github.com/hyperorchidlab/go-lib/ios
# 	cp -rf bin/iosLib.framework $(tp)
# 	rm -rf bin/iosLib.framework

clean:
	gomobile clean
	rm $(BINDIR)/*
