RM=rm
CPP=c++ -std=c++11
CFLAGS=-g -Wall
DFLAGS=-DDEBUG

%.o: %.cc
	$(CPP) $(DFLAGS) $(CFLAGS) -c $^ -o $@

broker: broker.o main.o
	$(CPP) $(DFLAGS) $(CFLAGS) broker.o main.o -o broker

dcenter: dcenter.o main.o
	$(CPP) $(DFLAGS) $(CFLAGS) dcenter.o main.o -o dcenter

all: dcenter broker

clean:
	$(RM) -f *.o broker dcenter >& /dev/null
