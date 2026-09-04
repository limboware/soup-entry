#include "shared.h"

PLUGIFY_EXPORT void (*__configs_PushNull)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__configs_PushBool)(uintptr_t, bool) = NULL;


PLUGIFY_EXPORT void (*__configs_PushInt)(uintptr_t, int64_t) = NULL;


PLUGIFY_EXPORT void (*__configs_PushFloat)(uintptr_t, double) = NULL;


PLUGIFY_EXPORT void (*__configs_PushString)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT void (*__configs_PushObject)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__configs_PushArray)(uintptr_t) = NULL;


