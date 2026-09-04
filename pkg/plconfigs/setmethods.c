#include "shared.h"

PLUGIFY_EXPORT void (*__configs_SetNull)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT void (*__configs_SetObject)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT void (*__configs_SetArray)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT void (*__configs_SetBool)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT void (*__configs_SetInt)(uintptr_t, String*, int64_t) = NULL;


PLUGIFY_EXPORT void (*__configs_SetFloat)(uintptr_t, String*, double) = NULL;


PLUGIFY_EXPORT void (*__configs_SetString)(uintptr_t, String*, String*) = NULL;


