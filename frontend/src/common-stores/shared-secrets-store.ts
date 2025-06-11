import type { ChatPartnerSecret } from "@/types/ChatPartnerSecret";

let sharedSecrets: Record<number, CryptoKey> = {};

export const getSharedSecret = (orderId: number): CryptoKey | undefined => {
  return sharedSecrets[orderId];
}

export const setSharedSecret = (orderId: number, secret: CryptoKey): void => {
  sharedSecrets[orderId] = secret;
};

export const setSharedSecrets = (secrets: ChatPartnerSecret[]): void => {
  sharedSecrets = secrets.reduce((acc, { userId, secret }) => {
    acc[userId] = secret;
    return acc;
  }, {} as Record<number, CryptoKey>);
}

export const clearSharedSecrets = (): void => {
  sharedSecrets = {};
}
