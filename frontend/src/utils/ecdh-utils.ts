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

export async function deriveSharedSecret(
  privateKeyBuffer: Uint8Array,
  publicKeyBuffer: Uint8Array
): Promise<CryptoKey> {
  const privateKey = await crypto.subtle.importKey(
    "pkcs8",
    privateKeyBuffer,
    { name: "ECDH", namedCurve: "P-256" },
    false,
    ["deriveBits"]
  );

  const publicKey = await crypto.subtle.importKey(
    "spki",
    publicKeyBuffer,
    { name: "ECDH", namedCurve: "P-256" },
    false,
    []
  );

  const bits = await crypto.subtle.deriveBits(
    { name: "ECDH", public: publicKey },
    privateKey,
    256
  );

  return crypto.subtle.importKey(
    "raw",
    bits,
    { name: "AES-GCM" },
    false,
    ["encrypt", "decrypt"]
  );
}


export const encryptWithECDHKey = async (
  value: string,
  aesKey: CryptoKey
): Promise<{ encrypted: Uint8Array; iv: Uint8Array }> => {
  const uint8Array = new TextEncoder().encode(value);
  const iv = crypto.getRandomValues(new Uint8Array(12)); // 12 bytes for AES-GCM

  const encrypted = await crypto.subtle.encrypt(
      { name: "AES-GCM", iv },
      aesKey,
      uint8Array
  );

  return {
      encrypted: new Uint8Array(encrypted),
      iv
  };
}

export const decryptWithECDHKey = async (
  encrypted: Uint8Array,
  iv: Uint8Array,
  aesKey: CryptoKey
): Promise<string> => {
  try {
    const decrypted = await crypto.subtle.decrypt(
      {
        name: "AES-GCM",
        iv,
      },
      aesKey,
      encrypted
    );

    return new TextDecoder().decode(decrypted);
  } catch (e) {
    console.error("Decryption failed:", e);
    throw new Error("Invalid decryption or tampered data.");
  }
};
