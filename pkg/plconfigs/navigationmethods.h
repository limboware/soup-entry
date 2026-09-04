#pragma once

#include "shared.h"

extern bool (*__configs_JumpFirst)(uintptr_t);

static bool JumpFirst(uintptr_t config) {
	return __configs_JumpFirst(config);
}

extern bool (*__configs_JumpLast)(uintptr_t);

static bool JumpLast(uintptr_t config) {
	return __configs_JumpLast(config);
}

extern bool (*__configs_JumpNext)(uintptr_t);

static bool JumpNext(uintptr_t config) {
	return __configs_JumpNext(config);
}

extern bool (*__configs_JumpPrev)(uintptr_t);

static bool JumpPrev(uintptr_t config) {
	return __configs_JumpPrev(config);
}

extern bool (*__configs_JumpKey)(uintptr_t, String*, bool);

static bool JumpKey(uintptr_t config, String* key, bool create) {
	return __configs_JumpKey(config, key, create);
}

extern bool (*__configs_JumpN)(uintptr_t, int32_t);

static bool JumpN(uintptr_t config, int32_t n) {
	return __configs_JumpN(config, n);
}

extern bool (*__configs_JumpBack)(uintptr_t);

static bool JumpBack(uintptr_t config) {
	return __configs_JumpBack(config);
}

extern void (*__configs_JumpRoot)(uintptr_t);

static void JumpRoot(uintptr_t config) {
	__configs_JumpRoot(config);
}

