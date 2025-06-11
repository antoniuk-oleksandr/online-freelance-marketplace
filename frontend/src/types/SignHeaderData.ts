export type SignHeaderData = {
  avatar: string | null,
  authenticated: boolean,
  userId: string | null,
  chatPartners?: {
    userId: number,
    publicKey: string | Uint8Array,
  }[],
}
