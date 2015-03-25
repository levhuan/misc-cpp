RM=rm
CC=c++
CFLAGS=-g 
DFLAGS=-DDEBUG_PRINT

all: test.cc
	$(CC) $(CFLAGS) test.cc -o test

debug: test.cc
	$(CC) $(CFLAGS) $(DFLAGS) test.cc -o test_debug

clean:
	$(RM) -fr *dSYM test_debug test > /dev/null
