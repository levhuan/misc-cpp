#include <string>

namespace mobilecloud {
namespace api {

class AppInterface {
public:
    static AppInterface * create();

    virtual void onStartup() = 0;
    virtual int  onProcessCmd(std::string input) = 0;
    virtual int  onExit() = 0;
};

}
}
