.PHONY: help clean test clean_test build run up

# set default goal to help
.DEFAULT_GOAL := help

# set application name
APPLICATION_NAME = positive_sum

### * make help                                            Print this help
help: Makefile
	@sed -n 's/^###//p' $<

### * make clean                                           Remove all artifacts
clean: 
	@ rm *.out || true

### * make test                                            test the application
test:
	@ gcc -c $(APPLICATION_NAME).c -o $(APPLICATION_NAME).out
	@ gcc -c $(APPLICATION_NAME)_test.c -o $(APPLICATION_NAME)_test.out
	@ gcc $(APPLICATION_NAME).out $(APPLICATION_NAME)_test.out -o test.out `pkg-config --cflags --libs check`
	@ ./test.out

### * make clean_test                                      test the application and clean
clean_test: test clean

### * make build                                           build the application
build:
	@ gcc -c $(APPLICATION_NAME).c -o $(APPLICATION_NAME).out
	@ gcc -c main.c -o main.out
	@ gcc $(APPLICATION_NAME).out main.out

### * make up                                             	run the application without build
up: 
	@ ./a.out

### * make run                                            	build and run the application
run: build up clean

