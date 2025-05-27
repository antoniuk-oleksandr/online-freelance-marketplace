export const convertToBase64 = (bytes: Uint8Array): string => {
  let binary = '';
  const len = bytes.length;
  for (let i = 0; i < len; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  return btoa(binary);
};

export const convertBase64ToUint8Array = (base64: string): Uint8Array => {
  return Uint8Array.from(atob(base64), c => c.charCodeAt(0));
}
