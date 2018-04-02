cmake_minimum_required(VERSION 3.2)

# Name of the project (will be the name of the plugin)
project (bidder-cpp CXX)

if ( NOT DEFINED VANILLA_RTB_ROOT )
message(FATAL_ERROR "Please setup VANILLA_RTB_ROOT a path  to vanilla-rtb sources")
endif()

SET(CMAKE_CXX_STANDARD 14)
SET(CMAKE_CXX_STANDARD_REQUIRED ON)
SET(CMAKE_CXX_EXTENSIONS OFF)
SET(CMAKE_POSITION_INDEPENDENT_CODE ON)

# Boost dependency
set(Boost_USE_STATIC_LIBS   ON)
set(Boost_USE_MULTITHREADED ON)
find_package(Threads)

# Include BoostLib module
file(GLOB_RECURSE boostlib_cmake_path "${CMAKE_CURRENT_SOURCE_DIR}/node_modules" "BoostLib.cmake")
list(GET boostlib_cmake_path 0 boostlib_cmake_path)
get_filename_component(boostlib_cmake_path "${boostlib_cmake_path}" DIRECTORY)
SET(CMAKE_MODULE_PATH "${CMAKE_MODULE_PATH};${boostlib_cmake_path}")
include(BoostLib)

# Locate/Download Boost (semver)
list(APPEND REQUIRED_BOOST_LIBRARIES "log" "program_options" "system" "serialization" "date_time" "regex")
require_boost_libs("1.66.0" "${REQUIRED_BOOST_LIBRARIES}")

include_directories(${Boost_INCLUDE_DIRS})

# Essential include files to build a node addon,
# you should add this line in every CMake.js based project.
set(JSONV_ROOT  ${VANILLA_RTB_ROOT}/jsonv/include)
set(RTB_ROOT   ${VANILLA_RTB_ROOT}/rtb)
include_directories(${CMAKE_JS_INC} ${VANILLA_RTB_ROOT} ${JSONV_ROOT} ${RTB_ROOT})

# Declare the location of the source files
file(GLOB SOURCE_FILES  RELATIVE_PATH "." "*.cpp")
file(GLOB_RECURSE CRUD_SOURCE_FILES ${VANILLA_RTB_ROOT}/CRUD/service/*.cpp)
list(APPEND SOURCE_FILES "${RTB_ROOT}/core/logger.cpp" "${CRUD_SOURCE_FILES}")
file(GLOB ARCHIVE_FILES  RELATIVE_PATH "." "*.a")
file(GLOB_RECURSE BOOST_ARCHIVE_LIST ${Boost_LIBRARY_DIR}*.a)
string(REPLACE ";" " " BOOST_ARCHIVES "${BOOST_ARCHIVE_LIST}")


IF (CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)
    SET ( CMAKE_INSTALL_PREFIX ${CMAKE_SOURCE_DIR} )
ENDIF()

set(PROGRAM_PERMISSIONS_DEFAULT
    OWNER_WRITE OWNER_READ OWNER_EXECUTE
    GROUP_READ GROUP_EXECUTE
    WORLD_READ WORLD_EXECUTE)

if (NOT WIN32 AND NOT APPLE)
set(RT_LIB rt)
SET(CMAKE_SHARED_LIBRARY_LINK_${lang}_FLAGS "")
endif()

IF ( GO2CPP ) 
    message("Building GO2CPP interfaces...")
    add_library(${PROJECT_NAME} STATIC ${SOURCE_FILES})
    #list(REMOVE_ITEM ARCHIVE_FILES ${CMAKE_BINARY_DIR}/bid_handler.a)
    #set( ENV{CGO_LDFLAGS} "${ARCHIVE_FILES} ${CMAKE_BINARY_DIR}/lib${PROJECT_NAME}.a ${BOOST_ARCHIVES}")
    set( ENV{CGO_LDFLAGS} "${CMAKE_BINARY_DIR}/lib${PROJECT_NAME}.a ${BOOST_ARCHIVES}")
    #message("CGO_LDFLAGS=" $ENV{CGO_LDFLAGS})
    add_custom_command(
        TARGET
        ${PROJECT_NAME}
        POST_BUILD
        COMMAND echo $ENV{CGO_LDFLAGS} | ${CMAKE_SOURCE_DIR}/go2cpp.sh ${CMAKE_BINARY_DIR}/bidder-go ${CMAKE_SOURCE_DIR}/bidder.go
        #COMMAND go build  -buildmode=exe -o ${CMAKE_BINARY_DIR}/bidder ${CMAKE_SOURCE_DIR}/bidder.go
    )
    install(FILES ${CMAKE_BINARY_DIR}/bidder-go DESTINATION ${CMAKE_INSTALL_PREFIX} PERMISSIONS ${PROGRAM_PERMISSIONS_DEFAULT})
ELSE ()
    message("Building CPP2GO interfaces...")
    list(REMOVE_ITEM SOURCE_FILES ${CMAKE_CURRENT_SOURCE_DIR}/bidder_binding.cpp)
    add_executable(${PROJECT_NAME} ${SOURCE_FILES})
    target_link_libraries(${PROJECT_NAME} ${CMAKE_JS_LIB} ${Boost_LIBRARIES} ${CMAKE_THREAD_LIBS_INIT} ${ARCHIVE_FILES} ${RT_LIB})
    install( TARGETS ${PROJECT_NAME}
        RUNTIME DESTINATION ${CMAKE_INSTALL_PREFIX}
        PERMISSIONS ${PROGRAM_PERMISSIONS_DEFAULT}
    )
ENDIF()


message(${Boost_LIBRARIES})
#message(${BOOST_ARCHIVES})

