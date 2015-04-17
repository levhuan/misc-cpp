#include <iostream>
#include "appInterface.h"

using namespace mobilecloud::api;
using namespace std;

class Datacenter: public AppInterface {
public:
    Datacenter() {}
    ~Datacenter() {}
    void onStartup() {
    }

    int onProcessCmd(std::string input) {
        return 0;
    }

    int  onExit() {
        return 0;
    }
};

AppInterface *AppInterface::create() {
    return new Datacenter;
}
