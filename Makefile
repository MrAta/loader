build:
	go build loader.go

run:
	chmod a+x run.sh
	./run.sh

clean:
	rm -rf loader
	rm -rf *.log
