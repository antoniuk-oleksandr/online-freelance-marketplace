let privateKey: Uint8Array = new Uint8Array(0);

export const setPrivateKey = (key: Uint8Array) => {
  privateKey = key;
}

export const getPrivateKey = (): Uint8Array => privateKey;
