RM=rm
CC=c++
CFLAGS=-g 
DFLAGS=-DDEBUG_PRINT

all: test.cc solution.cc
	$(CC) $(CFLAGS) test.cc -o test
	$(CC) $(CFLAGS) solution.cc -o solution

debug: test.cc
	$(CC) $(CFLAGS) $(DFLAGS) test.cc -o test_debug

clean:
	$(RM) -fr a.out *dSYM test_debug test solution > /dev/null
