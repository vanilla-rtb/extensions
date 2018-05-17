// TODO: autogenerate correct file
#include <boost/python.hpp>
#include <vector>
#include <boost/scoped_array.hpp>

extern int __main__(int, char**);


void RunBidder (int argc, boost::python::list __argv) {
    std::vector<std::string> args;
    for ( int i=0; i < argc; ++i) {
       args.push_back(boost::python::extract<std::string>(__argv[i]));
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

BOOST_PYTHON_MODULE(pyBidder)
{
    using namespace boost::python;
    def("RunBidder" , RunBidder); 
}

