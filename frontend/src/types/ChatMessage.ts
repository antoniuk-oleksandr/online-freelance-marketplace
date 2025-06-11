import type { ChatFile } from "./ChatFile"
import type { ChatMessageType } from "./ChatMessageType"

export type ChatMessage = {
  content: string,
  contentIV: string,
  sentAt: number,
  type: ChatMessageType,
  orderId?: number,
  id?: number,
  senderId?: number,
  files?: ChatFile[],
  chatPartnerId?: number,
}
