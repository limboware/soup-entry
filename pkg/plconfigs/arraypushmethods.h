#pragma once

#include "shared.h"

extern void (*__configs_PushNull)(uintptr_t);

static void PushNull(uintptr_t config) {
	__configs_PushNull(config);
}

extern void (*__configs_PushBool)(uintptr_t, bool);

static void PushBool(uintptr_t config, bool value) {
	__configs_PushBool(config, value);
}

extern void (*__configs_PushInt)(uintptr_t, int64_t);

static void PushInt(uintptr_t config, int64_t value) {
	__configs_PushInt(config, value);
}

extern void (*__configs_PushFloat)(uintptr_t, double);

static void PushFloat(uintptr_t config, double value) {
	__configs_PushFloat(config, value);
}

extern void (*__configs_PushString)(uintptr_t, String*);

static void PushString(uintptr_t config, String* value) {
	__configs_PushString(config, value);
}

extern void (*__configs_PushObject)(uintptr_t);

static void PushObject(uintptr_t config) {
	__configs_PushObject(config);
}

extern void (*__configs_PushArray)(uintptr_t);

static void PushArray(uintptr_t config) {
	__configs_PushArray(config);
}

