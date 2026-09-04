#pragma once

#include "shared.h"

extern String (*__configs_NodeToJsonString)(uintptr_t);

static String NodeToJsonString(uintptr_t config) {
	return __configs_NodeToJsonString(config);
}

extern String (*__configs_RootToJsonString)(uintptr_t);

static String RootToJsonString(uintptr_t config) {
	return __configs_RootToJsonString(config);
}

