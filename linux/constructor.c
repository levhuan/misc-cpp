#include <stdio.h>
#include <stdlib.h>

void constructor_one(void) __attribute__((constructor));
void constructor_two(void) __attribute__((constructor(1000)));
void constructor_three(void) __attribute__((constructor(1001)));
void constructor_four(void) __attribute__((constructor));

void constructor_one(void) {
		    printf("%s\n", __FUNCTION__);
}

void constructor_two(void) {
		    printf("%s\n", __FUNCTION__);
}

void constructor_three(void) {
		    printf("%s\n", __FUNCTION__);
}

void constructor_four(void) {
		    printf("%s\n", __FUNCTION__);
}

int main(char **argv, int argc) {
		    printf("%s\n", __FUNCTION__);
}
