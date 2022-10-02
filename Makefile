all:
	go build "-ldflags=-s -w -buildid=" -trimpath -o clouddns
