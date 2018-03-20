#include <vector>
#include <string>
#include <boost/scoped_array.hpp>
#include "bid_handler.h"

extern int __main__(int, char**);

#ifdef __cplusplus
extern "C" {
#endif
void RunBidder(char * _0, char *_1, char *_2) {
    std::vector<std::string> args = { _0, _1 , _2 } ;
    //for(int i=0; i< slice.len; ++i) {
    //    args.push_back((char *)slice.data);
    //}
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
#ifdef __cplusplus
}
#endif

