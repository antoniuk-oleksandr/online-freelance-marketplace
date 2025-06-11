import { sendWebsocketMessage } from "@/api/websocket";
import type { ChatFormData } from "@/types/ChatFormData";
import type { ChatMessageToSend } from "@/types/ChatMessageToSend";
import { ChatMessageType } from "@/types/ChatMessageType";
import type { ChatMessage } from "@/types/ChatMessage";
import { get } from "svelte/store";
import { signDataStore } from "@/common-components/Header/components/HeaderProfileBlock/sign-data-store";
import type { WebSocketRequest } from "@/types/WebSocketRequest";
import { WebSocketRequestType } from "@/types/WebSocketRequestType";
import { getSharedSecret } from "@/common-stores/shared-secrets-store";
import { errorStore } from "@/common-stores/error-store";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import { encryptWithECDHKey } from "@/utils/ecdh-utils";
import { convertToBase64 } from "@/utils/base64-utils";

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

export const handleChatFormSubmit = async (
  data: ChatFormData,
  reset: () => void,
) => {
  reset();
  if (!signDataStore) return;

  const senderId = get(signDataStore)!.userId;
  if (!senderId) return;

  const { chatParentId } = data;
  const sharedSecret = getSharedSecret(chatParentId);
  if (sharedSecret === undefined) {
    errorStore.set({ shown: true, error: ResponseErrorEnum.ErrInvalidSharedSecret });
    return;
  }

  const { encrypted, iv } = await encryptWithECDHKey(data.message, sharedSecret)
  const base64Content = convertToBase64(encrypted);
  const base64ContentIV = convertToBase64(iv);

  const message: ChatMessage = {
    type: ChatMessageType.Sent,
    content: base64Content,
    contentIV: base64ContentIV,
    sentAt: Date.now(),
    chatPartnerId: data.chatParentId,
    senderId: parseInt(senderId),
    orderId: data.orderId
  }

  const websocketRequestData: WebSocketRequest = {
    data: message,
    type: WebSocketRequestType.ChatMessage,
  }

  sendWebsocketMessage(websocketRequestData);
}

export const handleChatKeyDown = (event: KeyboardEvent, handleSubmit: () => void) => {
  const isInForm = (event.target as HTMLElement)?.closest('form') !== null

  if (!isInForm) return

  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    handleSubmit()
  }
}
