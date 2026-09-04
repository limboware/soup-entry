#include "shared.h"

PLUGIFY_EXPORT bool (*__configs_TrySetFromBool)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__configs_TrySetFromInt)(uintptr_t, String*, int64_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_TrySetFromFloat)(uintptr_t, String*, double) = NULL;


PLUGIFY_EXPORT bool (*__configs_TrySetFromString)(uintptr_t, String*, String*) = NULL;


