import { sendWebsocketMessage } from "@/api/websocket";
import type { ChatFormData } from "@/types/ChatFormData";
import type { ChatMessageToSend } from "@/types/ChatMessageToSend";
import { ChatMessageType } from "@/types/ChatMessageType";

export const handleChatFileDrop = (
  e: DragEvent,
  setFiles: (files: File[]) => void
) => {
  e.preventDefault();

  const files = e.dataTransfer?.files;
  if (!files) return;

  setFiles(Array.from(files));
}

export const handleChatFileAttachDrag = (
  e: DragEvent,
  setDragValue: (value: boolean) => void,
  value: boolean,
) => {
  e.preventDefault();
  setDragValue(true);
}

export const handleChatFileChange = (
  e: Event,
  setFiles: (files: File[]) => void
) => {
  if (!e.target) return;
  const target = e.target as HTMLInputElement;

  setFiles(Array.from(target.files || []));
}

export const handleChatFormSubmit = (
  data: ChatFormData, reset: () => void,
) => {
  reset();

  const message: ChatMessageToSend = {
    type: ChatMessageType.Sent,
    content: data.message,
    sentAt: Date.now(),
    chatPartnerId: 0,
  }

  sendWebsocketMessage<ChatMessageToSend>(message);
}
