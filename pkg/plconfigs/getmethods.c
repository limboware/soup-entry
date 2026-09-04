#include "shared.h"

PLUGIFY_EXPORT bool (*__configs_GetBool)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT int64_t (*__configs_GetInt)(uintptr_t, String*, int64_t) = NULL;


PLUGIFY_EXPORT double (*__configs_GetFloat)(uintptr_t, String*, double) = NULL;


PLUGIFY_EXPORT String (*__configs_GetString)(uintptr_t, String*, String*) = NULL;


