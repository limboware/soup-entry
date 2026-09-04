#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__configs_Read)(String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__configs_ReadMultiple)(Vector*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__configs_Make)() = NULL;


PLUGIFY_EXPORT void (*__configs_Delete)(uintptr_t) = NULL;


