#pragma once

#include "shared.h"

extern int32_t (*__configs_GetType)(uintptr_t);

static int32_t GetType(uintptr_t config) {
	return __configs_GetType(config);
}

extern bool (*__configs_IsNull)(uintptr_t);

static bool IsNull(uintptr_t config) {
	return __configs_IsNull(config);
}

extern bool (*__configs_IsBool)(uintptr_t);

static bool IsBool(uintptr_t config) {
	return __configs_IsBool(config);
}

extern bool (*__configs_IsInt)(uintptr_t);

static bool IsInt(uintptr_t config) {
	return __configs_IsInt(config);
}

extern bool (*__configs_IsFloat)(uintptr_t);

static bool IsFloat(uintptr_t config) {
	return __configs_IsFloat(config);
}

extern bool (*__configs_IsString)(uintptr_t);

static bool IsString(uintptr_t config) {
	return __configs_IsString(config);
}

extern bool (*__configs_IsObject)(uintptr_t);

static bool IsObject(uintptr_t config) {
	return __configs_IsObject(config);
}

extern bool (*__configs_IsArray)(uintptr_t);

static bool IsArray(uintptr_t config) {
	return __configs_IsArray(config);
}

