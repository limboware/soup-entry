#include "shared.h"

PLUGIFY_EXPORT bool (*__configs_GetAsBool)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT int64_t (*__configs_GetAsInt)(uintptr_t, String*, int64_t) = NULL;


PLUGIFY_EXPORT double (*__configs_GetAsFloat)(uintptr_t, String*, double) = NULL;


PLUGIFY_EXPORT String (*__configs_GetAsString)(uintptr_t, String*, String*) = NULL;


