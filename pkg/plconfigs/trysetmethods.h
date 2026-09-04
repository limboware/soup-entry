#pragma once

#include "shared.h"

extern bool (*__configs_TrySetFromBool)(uintptr_t, String*, bool);

static bool TrySetFromBool(uintptr_t config, String* key, bool value) {
	return __configs_TrySetFromBool(config, key, value);
}

extern bool (*__configs_TrySetFromInt)(uintptr_t, String*, int64_t);

static bool TrySetFromInt(uintptr_t config, String* key, int64_t value) {
	return __configs_TrySetFromInt(config, key, value);
}

extern bool (*__configs_TrySetFromFloat)(uintptr_t, String*, double);

static bool TrySetFromFloat(uintptr_t config, String* key, double value) {
	return __configs_TrySetFromFloat(config, key, value);
}

extern bool (*__configs_TrySetFromString)(uintptr_t, String*, String*);

static bool TrySetFromString(uintptr_t config, String* key, String* value) {
	return __configs_TrySetFromString(config, key, value);
}

