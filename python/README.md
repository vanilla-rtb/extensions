# python 2.7 bindings - Work in progress 

The bindings based on [Boost.Python](https://www.boost.org/doc/libs/1_66_0/libs/python/doc/html/index.html)

### Build pyBidder module and Run

```
npm install
python bidder.py  --config config.cfg
```

#### Warning !!!
cmake script is automatically placing ${HOME}/user-config.jam for boost python libraries build

```
#fixes for boost_python
if(PYTHONLIBS_FOUND)
   find_package(PythonInterp)
   if(PYTHONINTERP_FOUND)
       set(user_jam "using python : : ${PYTHON_EXECUTABLE} : ${PYTHON_INCLUDE_PATH} \;")
       message(INFO ${user_jam})
       message(INFO "creating file $ENV{HOME}/user-config.jam")
       file(WRITE "$ENV{HOME}/user-config.jam"  ${user_jam})
   endif()
endif()
```

The solution  provides limited support for Python 2.x and will not work with Python 3.x without work arounds 
```
.cmake-js/boost/1.66.0/stage/lib/libboost_python3-mt.a
```
You may need to change few  lines 
```
find_package( PythonLibs 3.6 REQUIRED )
```
Then this might help
```
list(APPEND REQUIRED_BOOST_LIBRARIES "log" "program_options" "system" "serialization" "date_time" "regex" "python3")
```

The current build  will also break in Boost 1.67.0 because of 27 is appended to python library name starting with boost 1.67
```
.cmake-js/boost/1.67.0/stage/lib/libboost_python27-mt.a
```

If you need to run it with Boost >= 1.67 try to append id to python library name
```
list(APPEND REQUIRED_BOOST_LIBRARIES "log" "program_options" "system" "serialization" "date_time" "regex" "python27")
require_boost_libs("1.67.0" "${REQUIRED_BOOST_LIBRARIES}")
``` 

