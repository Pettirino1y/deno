# Deno Standard Modules

[![Build Status](https://dev.azure.com/denoland/deno_std/_apis/build/status/denoland.deno_std?branchName=master)](https://dev.azure.com/denoland/deno_std/_build/latest?definitionId=2?branchName=master)

These modules do not have external dependencies and they are reviewed by the
Deno core team. The intention is to have a standard set of high quality code
that all Deno projects can use fearlessly.

Contributions are welcome!

## How to use

These modules are tagged in accordance with Deno releases. So, for example, the
v0.2.6 tag is guaranteed to work with deno v0.2.6.
You can link to v0.2.6 using the URL `https://deno.land/x/std@v0.2.6/`

It's strongly recommended that you link to tagged releases rather than the
master branch. The project is still young and we expect disruptive renames in
the future.

## Style Guide

### Use TypeScript

### Use the term "module" instead of "library" or "package"

For clarity and consistency avoid the terms "library" and "package". Instead use
"module" to refer to a single JS or TS file and also to refer to a directory of
TS/JS code.

### Do not use the filename `index.ts` nor `index.js`

Deno does not treat "index.js" or "index.ts" in a special way. By using these
filenames, it suggests that they can be left out of the module specifier when
they cannot. This is confusing.

If a directory of code needs a default entry point, use the filename `mod.ts`.
The filename `mod.ts` follows Rust’s convention, is shorter than `index.ts`, and
doesn’t come with any preconceived notions about how it might work.

### Within `deno_std`, do not depend on external code

`deno_std` is intended to be baseline functionality that all Deno programs can
rely on. We want to guarantee to users that this code does not include
potentially unreviewed third party code.

### Within `deno_std`, minimize dependencies; do not make circular imports.

Although `deno_std` is a standalone codebase, we must still be careful to keep
the internal dependencies simple and manageable. In particular, be careful to
not to introduce circular imports.

### For consistency, use underscores, not dashes in filenames.

Example: Instead of `file-server.ts` use `file_server.ts`.

### Format code according using prettier.

More specifically, code should be wrapped at 80 columns and use 2-space
indentation and use camel-case. Use `//format.ts` to invoke prettier.

### Exported functions: max 2 args, put the rest into an options object.

When designing function interfaces, stick to the following rules.

1. A function that is part of the public API takes 0-2 required arguments,
   plus (if necessary) an options object (so max 3 total).

2. Optional parameters should generally go into the options object.

   An optional parameter that's not in an options object might be acceptable
   if there is only one, and it seems inconceivable that we would add more
   optional parameters in the future.

3. The 'options' argument is the only argument that is a regular 'Object'.

   Other arguments can be objects, but they must be distinguishable from a
   'plain' Object runtime, by having either:

   - a distinguishing prototype (e.g. `Array`, `Map`, `Date`, `class MyThing`)
   - a well-known symbol property (e.g. an iterable with `Symbol.iterator`).

   This allows the API to evolve in a backwards compatible way, even when the
   position of the options object changes.

```ts
// BAD: optional parameters not part of options object. (#2)
export function resolve(
  hostname: string,
  family?: "ipv4" | "ipv6",
  timeout?: number
): IPAddress[] {}

// GOOD.
export interface ResolveOptions {
  family?: "ipv4" | "ipv6";
  timeout?: number;
}
export function resolve(
  hostname: string,
  options: ResolveOptions = {}
): IPAddress[] {}
```

```ts
export interface Environment {
  [key: string]: string;
}

// BAD: `env` could be a regular Object and is therefore indistinguishable
// from an options object. (#3)
export function runShellWithEnv(cmdline: string, env: Environment): string {}

// GOOD.
export interface RunShellOptions {
  env: Environment;
}
export function runShellWithEnv(
  cmdline: string,
  options: RunShellOptions
): string {}
```

```ts
// BAD: more than 3 arguments (#1), multiple optional parameters (#2).
export function renameSync(
  oldname: string,
  newname: string,
  replaceExisting?: boolean,
  followLinks?: boolean
) {}

// GOOD.
interface RenameOptions {
  replaceExisting?: boolean;
  followLinks?: boolean;
}
export function renameSync(
  oldname: string,
  newname: string,
  options: RenameOptions = {}
) {}
```

```ts
// BAD: too many arguments. (#1)
export function pwrite(
  fd: number,
  buffer: TypedArray,
  offset: number,
  length: number,
  position: number
) {}

// BETTER.
export interface PWrite {
  fd: number;
  buffer: TypedArray;
  offset: number;
  length: number;
  position: number;
}
export function pwrite(options: PWrite) {}
```

### Use JS Doc to document exported machinery

We strive for complete documentation. Every exported symbol ideally should have
a documentation line.

If possible, use a single line for the JS Doc. Example:

```ts
/** foo does bar. */
export function foo() {
  // ...
}
```

See [CONTRIBUTING.md](https://github.com/denoland/deno/blob/master/.github/CONTRIBUTING.md#documenting-apis)
for more details.

### TODO Comments

TODO comments should be include an issue or the author's github username in
parentheses. Example:

```
// TODO(ry) Add tests.
// TODO(#123) Support Windows.
```

### Copyright headers

Most files in `deno_std` should have the following copyright header:

```
// Copyright 2018-2019 the Deno authors. All rights reserved. MIT license.
```

If the code originates elsewhere, ensure that the file has the proper copyright
headers. We only allow MIT, BSD, and Apache licensed code in `deno_std`.

### Top level functions should not use arrow syntax

Top level functions should use the `function` keyword. Arrow syntax should be
limited to closures.

Bad

```
export const foo(): string => {
  return "bar";
}
```

Good

```
export function foo(): string {
  return "bar";
}
```

### Meta-programming is discouraged. Including the use of Proxy.

Be explicit even when it means more code.

There are some situations where it may make sense to use such techniques, but in
the vast majority of cases it does not.

### If a filename starts with underscore, do not link to it: `_foo.ts`

Sometimes there maybe situations where an internal module is necessary but its
API is not meant to be stable or linked to. In this case prefix it with an
underscore. By convention, only files in its own directory should import it.

### When referencing Deno online, use the #denoland tag.

The name "deno" unfortunately is not especially unique on the internet. In order
to centralize the community, please tag github project, tweet, and other content
with `#denoland`.

---

Copyright 2018-2019 the Deno authors. All rights reserved. MIT license.
