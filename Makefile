install: create-folder build-go

create-folder: 
	if [ ! -d "./tmp" ]; then\
		mkdir tmp;\
	fi 

build-go: 
	go build
