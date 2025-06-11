import type { ChatPartnerPublicKeyData } from "./ChatPartnerPublicKeyData"
import type { UserData } from "./UserData"

export type SignInData = {
  userData: UserData,
  chatPartners: ChatPartnerPublicKeyData[]
}
