import type { ChatMessage } from "./ChatMessage"
import type { ChatPartner } from "./ChatPartner"

export type ChatData = {
  chatPartner: ChatPartner,
  messages: ChatMessage[],
}
