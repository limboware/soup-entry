#include "shared.h"

PLUGIFY_EXPORT bool (*__configs_JumpFirst)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_JumpLast)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_JumpNext)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_JumpPrev)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_JumpKey)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__configs_JumpN)(uintptr_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__configs_JumpBack)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__configs_JumpRoot)(uintptr_t) = NULL;


