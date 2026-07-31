// Copyright (c) Microsoft Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import "testing"

func TestMicrosoftArchiveValidate(t *testing.T) {
	for _, test := range []struct {
		name    string
		archive MicrosoftArchive
		wantErr bool
	}{
		{name: "valid", archive: MicrosoftArchive{URL: "https://example.test/go.tar.gz", SHA256: "checksum"}},
		{name: "missing URL", archive: MicrosoftArchive{SHA256: "checksum"}, wantErr: true},
		{name: "missing checksum", archive: MicrosoftArchive{URL: "https://example.test/go.tar.gz"}, wantErr: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := test.archive.validate("msgo1.27.1-2")
			if (err != nil) != test.wantErr {
				t.Fatalf("validate error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
