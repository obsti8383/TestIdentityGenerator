package main

import (
	"testing"
)

func TestStandardCaseWithManagers(t *testing.T) {
	testIdentityCount := 2000
	testManagerCount := 1000
	allIdentities := generate(testIdentityCount, "test.test", testManagerCount)

	// check counts and duplicates
	countIdentities := len(allIdentities)
	countManagers := 0
	countNonManagers := 0
	countDuplicateEmails := 0
	emailDuplicates := make(map[string]bool)
	for _, identity := range allIdentities {
		if identity.Manager == "" {
			countManagers++
		} else {
			countNonManagers++
		}
		if emailDuplicates[identity.Email] {
			countDuplicateEmails++
		}
		emailDuplicates[identity.Email] = true
	}
	if countDuplicateEmails > 0 {
		t.Error(countDuplicateEmails, "duplicate emails. this should only happen in very large datasets. check randum number generator")
	}
	if countIdentities != testIdentityCount {
		t.Error("mismatch in number of identities, should be", testIdentityCount, ", are:", countIdentities)
	}
	if countManagers != testManagerCount {
		t.Error("mismatch in number of managers, should be ", testManagerCount, ", are:", countManagers)
	}
	if countNonManagers != testIdentityCount-testManagerCount {
		t.Error("mismatch in number of non-managers, should be", testIdentityCount-testManagerCount, ", are:", countNonManagers)
	}
}

func TestStandardCaseWithoutManagers(t *testing.T) {
	testIdentityCount := 2000
	testManagerCount := 0
	allIdentities := generate(testIdentityCount, "test.test", testManagerCount)

	// check counts and duplicates
	countIdentities := len(allIdentities)
	countFilledManagerFields := 0
	countDuplicateEmails := 0
	emailDuplicates := make(map[string]bool)
	for _, identity := range allIdentities {
		if identity.Manager != "" {
			countFilledManagerFields++
		}
		if emailDuplicates[identity.Email] {
			countDuplicateEmails++
		}
		emailDuplicates[identity.Email] = true
	}
	if countDuplicateEmails > 0 {
		t.Error(countDuplicateEmails, "duplicate emails. this should only happen in very large datasets. check randum number generator")
	}
	if countIdentities != testIdentityCount {
		t.Error("mismatch in number of identities, should be", testIdentityCount, ", are:", countIdentities)
	}
	if countFilledManagerFields > 0 {
		t.Error("manager fields are not all empty,", countFilledManagerFields, " are filled")
	}
}
