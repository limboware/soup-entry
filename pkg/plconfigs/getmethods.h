#pragma once

#include "shared.h"

extern bool (*__configs_GetBool)(uintptr_t, String*, bool);

static bool GetBool(uintptr_t config, String* key, bool defaultValue) {
	return __configs_GetBool(config, key, defaultValue);
}

extern int64_t (*__configs_GetInt)(uintptr_t, String*, int64_t);

static int64_t GetInt(uintptr_t config, String* key, int64_t defaultValue) {
	return __configs_GetInt(config, key, defaultValue);
}

extern double (*__configs_GetFloat)(uintptr_t, String*, double);

static double GetFloat(uintptr_t config, String* key, double defaultValue) {
	return __configs_GetFloat(config, key, defaultValue);
}

extern String (*__configs_GetString)(uintptr_t, String*, String*);

static String GetString(uintptr_t config, String* key, String* defaultValue) {
	return __configs_GetString(config, key, defaultValue);
}

