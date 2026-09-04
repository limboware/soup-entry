#pragma once

#include "shared.h"

extern uintptr_t (*__configs_Read)(String*);

static uintptr_t Read(String* path) {
	return __configs_Read(path);
}

extern uintptr_t (*__configs_ReadMultiple)(Vector*);

static uintptr_t ReadMultiple(Vector* paths) {
	return __configs_ReadMultiple(paths);
}

extern uintptr_t (*__configs_Make)();

static uintptr_t Make() {
	return __configs_Make();
}

extern void (*__configs_Delete)(uintptr_t);

static void Delete(uintptr_t config) {
	__configs_Delete(config);
}

