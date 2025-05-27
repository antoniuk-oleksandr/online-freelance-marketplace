import type { UserKeys } from "@/types/UserKeys";

export const generateECDHKeyPair = async (): Promise<UserKeys> => {
  const keyPair = await window.crypto.subtle.generateKey(
    {
      name: "ECDH",
      namedCurve: "P-256",
    },
    true,
    ["deriveKey", "deriveBits"]
  );


  const publicKey = await window.crypto.subtle.exportKey("spki", keyPair.publicKey);
  const privateKey = await window.crypto.subtle.exportKey("pkcs8", keyPair.privateKey);

  return {
    privateKey: new Uint8Array(privateKey),
    publicKey: new Uint8Array(publicKey)
  }
}

export const deriveSharedSecret = async (
  userPrivateKeyBase64: string,
  partnerPublicKeyBase64: string
): Promise<ArrayBuffer> => {
  const privateKeyBuffer = Uint8Array.from(atob(userPrivateKeyBase64), c => c.charCodeAt(0));
  const publicKeyBuffer = Uint8Array.from(atob(partnerPublicKeyBase64), c => c.charCodeAt(0));

  const privateKey = await window.crypto.subtle.importKey(
    "pkcs8",
    privateKeyBuffer,
    { name: "ECDH", namedCurve: "P-256" },
    false,
    ["deriveKey", "deriveBits"]
  );

  // Import public key
  const publicKey = await window.crypto.subtle.importKey(
    "spki",
    publicKeyBuffer,
    { name: "ECDH", namedCurve: "P-256" },
    false,
    []
  );

  const sharedSecret = await window.crypto.subtle.deriveBits(
    {
      name: "ECDH",
      public: publicKey
    },
    privateKey,
    256 
  );

  return sharedSecret;
};
