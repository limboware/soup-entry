#pragma once

#include "shared.h"

extern void (*__configs_SetNull)(uintptr_t, String*);

static void SetNull(uintptr_t config, String* key) {
	__configs_SetNull(config, key);
}

extern void (*__configs_SetObject)(uintptr_t, String*);

static void SetObject(uintptr_t config, String* key) {
	__configs_SetObject(config, key);
}

extern void (*__configs_SetArray)(uintptr_t, String*);

static void SetArray(uintptr_t config, String* key) {
	__configs_SetArray(config, key);
}

extern void (*__configs_SetBool)(uintptr_t, String*, bool);

static void SetBool(uintptr_t config, String* key, bool value) {
	__configs_SetBool(config, key, value);
}

extern void (*__configs_SetInt)(uintptr_t, String*, int64_t);

static void SetInt(uintptr_t config, String* key, int64_t value) {
	__configs_SetInt(config, key, value);
}

extern void (*__configs_SetFloat)(uintptr_t, String*, double);

static void SetFloat(uintptr_t config, String* key, double value) {
	__configs_SetFloat(config, key, value);
}

extern void (*__configs_SetString)(uintptr_t, String*, String*);

static void SetString(uintptr_t config, String* key, String* value) {
	__configs_SetString(config, key, value);
}

