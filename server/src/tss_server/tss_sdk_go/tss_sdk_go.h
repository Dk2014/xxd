#ifdef __cplusplus
extern "C"
{
#endif

#include "tss_sdk.h"
#include "tss_sdk_anti.h"
#include <stdio.h>

#if defined(WIN32) || defined(_WIN64)

int tss_sdk_load_win(const char *shared_lib_file);
void *tss_sdk_get_syml_win(const char *syml_name);

#else

static int tss_sdk_load_linux(const char *shared_lib_file);
void *tss_sdk_get_syml_linux(const char *syml_name);

#endif

//custom
const TssSdkAntiInterfEx *M_TSS_SDK_GET_ANTI_INTERF_EX(const TssSdkAntiInitInfoEx *init_data);
int call_proc(TssSdkBusiInterf *busi_interf_);


int tss_sdk_load_impl(const char *shared_lib_file);
void *tss_sdk_get_syml_impl(const char *syml_name);
static const TssSdkBusiInterf *tss_sdk_get_comm_interf(const TssSdkInitInfo* init_info);
const TssSdkBusiInterf *tss_sdk_load(const char *shared_lib_dir, const TssSdkInitInfo *init_info);
int tss_sdk_unload();
// typedef const void *(*TssSdkGetInterf)(const void *init_data);
const void *tss_sdk_get_busi_interf(const char *syml_name, const void *data);

#ifdef __cplusplus
} /* end of extern "C" */
#endif