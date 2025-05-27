import type { ChatFile } from "./ChatFile";
import type { ChatMessageType } from "./ChatMessageType";

export type ChatMessageToSend = {
  sentAt: number,
  type: ChatMessageType,
  chatPartnerId: number,
  content?: string,
  files?: ChatFile[],
}
