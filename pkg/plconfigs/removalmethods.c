#include "shared.h"

PLUGIFY_EXPORT int32_t (*__configs_Remove)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_RemoveKey)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT void (*__configs_Clear)(uintptr_t) = NULL;


