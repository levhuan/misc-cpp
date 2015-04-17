#include <iostream>
#include <signal.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include "appInterface.h"


using namespace std;
using namespace mobilecloud::api;

const char * g_named_pipe = "/tmp/broker_info";

bool process_usr1 = false;
bool process_usr2 = false;
bool process_hup = false;

const int MAX_BUFFER = 1024;
char g_buffer[1024];

extern "C" {
typedef void (*sig_handler_fn)(int sig, struct __siginfo *, void *);

void sigusr1_handler (int sig, struct __siginfo * siginfo, void *ctx) {
    process_usr1 = true;
}

void sigusr2_handler (int sig, struct __siginfo * siginfo, void *ctx) {
    process_usr2 = true;
}

void sigquit_handler (int sig, struct __siginfo * siginfo, void *ctx) {
}

void sighup_handler (int sig, struct __siginfo * siginfo, void *ctx) {
    process_hup = true;
}

void setup_signal_handler(int signum, sig_handler_fn handler) {
    struct sigaction new_action = {{0}};
    /*! initialize SA_SIGINFO new_action handler */
    new_action.sa_flags = (SA_SIGINFO | SA_RESTART);
    new_action.sa_sigaction = handler;

    if (sigaction(signum, &new_action, NULL)) {
        perror("sigaction: error not handled\n");
        exit(1);
    }
}
}

void exit_cleanup () {
    unlink(g_named_pipe);
}

int main (int argc, char **argv) {
    
    mkfifo(g_named_pipe, 0666);
    atexit(exit_cleanup);

    setup_signal_handler(SIGUSR1, sigusr1_handler);
    setup_signal_handler(SIGUSR2, sigusr2_handler);
    setup_signal_handler(SIGHUP,  sighup_handler);
    setup_signal_handler(SIGQUIT, sigquit_handler);

    int fd = open(g_named_pipe, O_RDONLY);
    ssize_t c;

    AppInterface *app = AppInterface::create();

    app->onStartup();

    while (true) {
        /**
         * Command line structure:
         * sender:<id>|cmd:<verb>|num_arg:<num>[|arg_name:<arg>]+$
         */
        c = read(fd, g_buffer, MAX_BUFFER);
        if (c > 0) {
            g_buffer[c] = 0;
            std::cout << "Get from pipe: " << g_buffer << endl;
        }

        if (process_usr1 || process_usr2 || process_hup) {
            std::cout << "Get USR1 signal: " << process_usr1 << endl
                 << "    USR2 signal: " << process_usr2 << endl
                 << "    HUP  signal: " << process_hup 
                 << endl;

            process_usr1 = process_usr2 = process_hup = false;
        }
    }

    return 0;
}
