#pragma once

#include "shared.h"

extern void (*__configs_Merge)(uintptr_t, uintptr_t);

static void Merge(uintptr_t config, uintptr_t other) {
	__configs_Merge(config, other);
}

extern void (*__configs_MergeMove)(uintptr_t, uintptr_t);

static void MergeMove(uintptr_t config, uintptr_t other) {
	__configs_MergeMove(config, other);
}

