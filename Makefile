all:
	go build "-ldflags=-s -w -buildid=" -trimpath -o ${HOME}/bin/clouddns
