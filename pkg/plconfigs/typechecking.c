#include "shared.h"

PLUGIFY_EXPORT int32_t (*__configs_GetType)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsNull)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsBool)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsInt)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsFloat)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsString)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsObject)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_IsArray)(uintptr_t) = NULL;


