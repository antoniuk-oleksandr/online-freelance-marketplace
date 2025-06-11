import type { ChatMessage } from "./ChatMessage";
import type { WebSocketRequestType } from "./WebSocketRequestType";

export type WebSocketRequest = {
  type: WebSocketRequestType.ChatMessage,
  data: ChatMessage,
}
