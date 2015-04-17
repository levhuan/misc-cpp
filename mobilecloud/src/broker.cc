#include "appInterface.h"

using namespace mobilecloud::api;

class Broker : public AppInterface {
public:
    ~Broker() {}
    Broker() {}

    void onStartup() {
    }    

    int onProcessCmd(std::string input) {
        return 0;
    }

    int  onExit() {
        return 0;
    }
};

AppInterface * AppInterface::create() {
    return new Broker;
}
