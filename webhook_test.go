package starling

import (
	"bytes"
	"net/http"
	"testing"
)

var publicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgIdCVYnz6JOFT7GGtjrMg4uaPRGGs5VlglSSd9i2i73zRp7AwZm8O/3LM5kPuPONOysJpdVSz9x6VGsRcaKkvMaOfYWYa6fe4l5IFiM8Z+WaL0WjIebdJOOjWxH3q/kW6KclwKBW0+2iNZPcZocllCOjPn/swp2MdhKLJOQkdB/1Q8Emxr6tsOlJkc2lWpXdtPHWUbBp31eF5/eDmuVCCBhTL76UyogQNgRV5qH2g/a2bNcNgTThR0PntXJLy2HLi9cEfXepevpoJM8HXNdaFwZV4pQUEzm3/jG7zI3isXnvtffG4uTIR8Q35yDrYeN8pX+zOAcnJYNbr9xdFEv7JQIDAQAB"
var signature = "KDGgtd7VDeyvNdyafyXNVZM8l/0zohWze5UCt1N0mbzCZ1f23nYEgnLrFvTRYADnToat/axKOGeXjiOBWJh/FcPvcWParx8x5d35j2u76/UmRPKjo8jxtMspmN27WlPdtTRr9kqHdDHUg80/9z1qKuEcUfm4EQX52NOvozDMb4qyYorgxaFCwUwMdZNskArIBTeJBtULAOtJqnEGipKRtRjeU6j2xD2uNzc3Vcy3+tdImRfqbX6SkS44zgkcFua6xEc09qRnRvLd+bxjSIufQ/wU695Uej9AtFg7MlrRCUaEZ2SVkNcmOUdRP2q882Y9mWGDIXdk66QHCVfCVu7pog=="

var validateTestCases = []struct {
	body  []byte
	key   string
	sig   string
	valid bool
}{
	{
		body:  []byte(`{"one":"Value","two":"Other"}`),
		key:   publicKey,
		sig:   signature,
		valid: true,
	},
	{
		body:  nil,
		key:   publicKey,
		sig:   signature,
		valid: false,
	},
	{
		body:  []byte(`{"one":"Value","two":"Other"}`),
		key:   publicKey,
		sig:   "[invalid]signature",
		valid: false,
	},
	{
		body:  []byte("[invalid]this is the request body"),
		key:   publicKey,
		sig:   signature,
		valid: false,
	},
	{
		body:  []byte(`{"one":"Value","two":"Other"}`),
		key:   "[invalid]publicKey",
		sig:   signature,
		valid: false,
	},
}

func TestValidate_Valid(t *testing.T) {
	for _, tc := range validateTestCases {
		req, err := http.NewRequest("POST", "http://localhost/callback", bytes.NewBuffer(tc.body))
		if err != nil {
			t.Error("should create a request without error:", err)
		}

		req.Header.Set("X-Hook-Signature", tc.sig)
		valid, err := Validate(req, tc.key)
		if err != nil && valid != tc.valid {
			t.Error("should be able to perform validation without error:", err)
		}

	}
}
