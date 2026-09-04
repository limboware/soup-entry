#pragma once

#include "shared.h"

extern void (*__configs_SetError)(String*);

static void SetError(String* error_) {
	__configs_SetError(error_);
}

extern String (*__configs_GetError)();

static String GetError() {
	return __configs_GetError();
}

