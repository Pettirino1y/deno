package main

import (
	"io/ioutil"
	"testing"
)

func SetCompileDirForTest(prefix string) {
	dir, err := ioutil.TempDir("", prefix)
	if err != nil {
		panic(err)
	}
	CompileDir = dir
}

func TestLoadOutputCodeCache(t *testing.T) {
	SetCompileDirForTest("TestLoadOutputCodeCache")

	filename := "Hello.ts"
	sourceCodeBuf := []byte("1+2")

	cacheFn := CacheFileName(filename, sourceCodeBuf)

	outputCode, err := LoadOutputCodeCache(filename, sourceCodeBuf)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if outputCode != "" {
		t.Fatalf("Expected empty outputCode but got <<%s>>", outputCode)
	}

	// Now let's write to the cache file
	err = ioutil.WriteFile(cacheFn, []byte("blah"), 0700)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Try it again.
	outputCode, err = LoadOutputCodeCache(filename, sourceCodeBuf)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if outputCode != "blah" {
		t.Fatalf("Bad outputCode but got <<%s>>", outputCode)
	}
}
