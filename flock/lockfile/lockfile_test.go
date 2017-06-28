package lockfile

import (
	"os"
	"testing"
)

func TestFileLockUnlock(t *testing.T) {
	test_file_path, _ := os.Getwd()
	locked_file := test_file_path + string(os.PathSeparator) + "0"


	flock := NewFileLock(locked_file)

	err := flock.Lock()
	if err != nil {
		t.Fatal(err)
	}

	err = flock.Lock()
	if err == nil {
		t.Fatal("can not be locked")
	}


	flock2 := NewFileLock(locked_file)

	err = flock2.Lock()
	if err == nil {
		t.Fatal("can not be locked")
	}


	err = flock.Unlock()
	if err != nil {
		t.Fatal(err)
	}

	err = flock2.Lock()
	if err != nil {
		t.Fatal(err)
	}

}

func TestFileLock_Lock(t *testing.T) {
	test_file_path, _ := os.Getwd()
	locked_file := test_file_path + string(os.PathSeparator) + "1"


	flock := NewFileLock(locked_file)

	err := flock.Unlock()
	if err == nil {
		t.Fatal("can not be unlocked")
	}

	err = flock.Lock()
	if err != nil {
		t.Fatal(err)
	}
	err = flock.Lock()
	if err == nil {
		t.Fatal("can not be locked")
	}
	err = flock.Unlock()
	if err  != nil {
		t.Fatal(err)
	}

	err = flock.Lock()
	if err != nil {
		t.Fatal(err)
	}

}

