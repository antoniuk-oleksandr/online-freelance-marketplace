import type { ChatFile } from "./ChatFile"
import type { ChatMessageType } from "./ChatMessageType"

export type ChatMessage = {
  id: number,
  senderId: number,
  content: string,
  sentAt: number,
  type: ChatMessageType,
  files?: ChatFile[],
}
