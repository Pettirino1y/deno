// Copyright 2018-2019 the Deno authors. All rights reserved. MIT license.
import { test, assert } from "./test_util.ts";

test(function platformTransform() {
  // deno.platform is injected by rollup at compile time. Here
  // we check it has been properly transformed.
  const { arch, os } = Deno.platform;
  assert(arch === "x64");
  assert(os === "mac" || os === "win" || os === "linux");
});
