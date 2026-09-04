#include "shared.h"

PLUGIFY_EXPORT bool (*__configs_HasKey)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT bool (*__configs_Empty)(uintptr_t) = NULL;


PLUGIFY_EXPORT int64_t (*__configs_GetSize)(uintptr_t) = NULL;


PLUGIFY_EXPORT String (*__configs_GetName)(uintptr_t) = NULL;


PLUGIFY_EXPORT String (*__configs_GetPath)(uintptr_t) = NULL;


