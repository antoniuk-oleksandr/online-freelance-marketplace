export type SignInEncryptionData = {
  privateKey: Uint8Array,
  privateKeyIV: Uint8Array,
  privateKeySalt: Uint8Array,
  masterKey: Uint8Array,
}
