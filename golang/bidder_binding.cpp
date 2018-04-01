
extern int __main__(int, char**);

#ifdef __cplusplus
extern "C" {
#endif
void RunBidder(int argc, char **argv) {
    __main__(argc, argv);
}
#ifdef __cplusplus
}
#endif

