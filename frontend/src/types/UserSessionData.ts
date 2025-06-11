import type { ChatPartnerPublicKeyData } from "./ChatPartnerPublicKeyData";

export type UserSessionData = {
  masterKey: string | Uint8Array,
  chatPartners: ChatPartnerPublicKeyData[],
  authenticated: boolean,
}
