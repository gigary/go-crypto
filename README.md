# INTRO
A set of helper functions to encode/decode using Base64 scheme and encrypt/decrypt using AES algorithms

# USAGE

## Download package
`go get github.com/gigary/crypto`

## Import to project
`import "github.com/gigary/crypto"`

## Helper functions
1. `crypto.Encrypt(plain, key, crypto.IV)` to encrypt a plain text using a secret key
2. `crypto.Decrypt(cipher, key, crypto.IV)` to decrypt a cipher text using a secret key
3. `crypto.Encode64(text)` to encode a string using Base64 scheme
4. `crypto.Decode64(text)` to decode a string using Base64 scheme

# CREDITS
 - [Michael Bui](mailto:mf.michaelbui@gmail.com) from [Gigary](https://gigary.com)