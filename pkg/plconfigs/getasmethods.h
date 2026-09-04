#pragma once

#include "shared.h"

extern bool (*__configs_GetAsBool)(uintptr_t, String*, bool);

static bool GetAsBool(uintptr_t config, String* key, bool defaultValue) {
	return __configs_GetAsBool(config, key, defaultValue);
}

extern int64_t (*__configs_GetAsInt)(uintptr_t, String*, int64_t);

static int64_t GetAsInt(uintptr_t config, String* key, int64_t defaultValue) {
	return __configs_GetAsInt(config, key, defaultValue);
}

extern double (*__configs_GetAsFloat)(uintptr_t, String*, double);

static double GetAsFloat(uintptr_t config, String* key, double defaultValue) {
	return __configs_GetAsFloat(config, key, defaultValue);
}

extern String (*__configs_GetAsString)(uintptr_t, String*, String*);

static String GetAsString(uintptr_t config, String* key, String* defaultValue) {
	return __configs_GetAsString(config, key, defaultValue);
}

