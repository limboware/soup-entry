#pragma once

#include "shared.h"

extern int32_t (*__configs_Remove)(uintptr_t);

static int32_t Remove(uintptr_t config) {
	return __configs_Remove(config);
}

extern bool (*__configs_RemoveKey)(uintptr_t, String*);

static bool RemoveKey(uintptr_t config, String* key) {
	return __configs_RemoveKey(config, key);
}

extern void (*__configs_Clear)(uintptr_t);

static void Clear(uintptr_t config) {
	__configs_Clear(config);
}

