export async function addSalt(
    key: string | Uint8Array,
    salt: Uint8Array
): Promise<CryptoKey> {
    key = typeof key === "string" ? (new TextEncoder()).encode(key) : key;

    const keyMaterial = await crypto.subtle.importKey(
        "raw",
        key,
        { name: "PBKDF2" },
        false,
        ["deriveKey"]
    );

    return await crypto.subtle.deriveKey(
        {
            name: "PBKDF2",
            salt,
            iterations: 100_000,
            hash: "SHA-256",
        },
        keyMaterial,
        {
            name: "AES-GCM",
            length: 256,
        },
        true,
        ["encrypt", "decrypt"]
    );
}


export async function encryptWithKey(
    value: Uint8Array,
    key: string | Uint8Array,
    iv: Uint8Array,
    salt: Uint8Array
): Promise<Uint8Array> {
    const cryptoKey = await addSalt(key, salt)
    const encrypted = await crypto.subtle.encrypt(
        { name: "AES-GCM", iv },
        cryptoKey,
        value
    );

    return new Uint8Array(encrypted);
}

export async function decryptWithKey(
    cipher: Uint8Array,
    key: string |Uint8Array,
    iv: Uint8Array,
    salt: Uint8Array
): Promise<Uint8Array> {
    const cryptoKey = await addSalt(key, salt);
    const decrypted = await crypto.subtle.decrypt(
        { name: "AES-GCM", iv },
        cryptoKey,
        cipher
    );

    return new Uint8Array(decrypted);
}

export const generateAESIV = (): Uint8Array => {
    return crypto.getRandomValues(new Uint8Array(12));
}

export const generateAESSalt = (): Uint8Array => {
    return crypto.getRandomValues(new Uint8Array(16));
}
