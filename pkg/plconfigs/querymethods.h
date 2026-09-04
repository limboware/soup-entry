#pragma once

#include "shared.h"

extern bool (*__configs_HasKey)(uintptr_t, String*);

static bool HasKey(uintptr_t config, String* key) {
	return __configs_HasKey(config, key);
}

extern bool (*__configs_Empty)(uintptr_t);

static bool Empty(uintptr_t config) {
	return __configs_Empty(config);
}

extern int64_t (*__configs_GetSize)(uintptr_t);

static int64_t GetSize(uintptr_t config) {
	return __configs_GetSize(config);
}

extern String (*__configs_GetName)(uintptr_t);

static String GetName(uintptr_t config) {
	return __configs_GetName(config);
}

extern String (*__configs_GetPath)(uintptr_t);

static String GetPath(uintptr_t config) {
	return __configs_GetPath(config);
}

