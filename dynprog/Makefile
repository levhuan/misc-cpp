RM=rm
CC=c++
CFLAGS=-g 
DFLAGS=-DDEBUG_PRINT

all: test.cc solution.cc
	$(CC) $(CFLAGS) test.cc -o tst
	$(CC) $(CFLAGS) solution.cc -o sol

debug: test.cc solution.cc
	$(CC) $(CFLAGS) $(DFLAGS) test.cc -o tst_debug
	$(CC) $(CFLAGS) $(DFLAGS) solution.cc -o sol_debug

clean:
	$(RM) -fr a.out *dSYM {sol,tst}_debug tst sol > /dev/null
