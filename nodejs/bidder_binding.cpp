// TODO: autogenerate correct file
#include <nan.h>
#include <vector>
#include <boost/scoped_array.hpp>

using v8::FunctionTemplate;
using v8::Handle;
using v8::Object;
using v8::String;

extern int __main__(int, char**);

NAN_METHOD(RunBidder) {
    std::vector<std::string> args;
    for ( int i=0; i < info.Length(); ++i) {
       String::Utf8Value str(info[i]);
       args.push_back(*str);
    }

    std::vector<char> params(1024) ;
    boost::scoped_array<char *> argv ( new char *[args.size()+1] ) ;
    std::vector<std::string>::const_iterator itr = args.begin() ;
    std::vector<char>::iterator pitr = params.begin() ;
    for ( int i = 0; itr != args.end() ; ++itr , ++pitr) {
         argv[i++] = &*pitr ;
         pitr = std::copy(itr->begin(), itr->end(), pitr) ;
    }
    argv[args.size()] = nullptr ;
    __main__(args.size(), argv.get());
}

//void InitAll(Handle<Object> exports) {
//  exports->Set(Nan::New<String>("runBidder").ToLocalChecked(),
//    Nan::New<FunctionTemplate>(RunBidder)->GetFunction());
//}

NAN_MODULE_INIT(InitAll) {
    Nan::Export(target, "runBidder", RunBidder);
}

NODE_MODULE(addon, InitAll)

