import { errorStore } from '@/common-stores/error-store';
import { getSharedSecret } from '@/common-stores/shared-secrets-store';
import { chatBodyStore } from '@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/components/OrderByIdChatTab/stores/chat-body-store';
import { messageStore } from '@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/components/OrderByIdChatTab/stores/message-store';
import type { ChatMessage } from '@/types/ChatMessage';
import { ResponseErrorEnum } from '@/types/ResponseErrorEnum';
import type { WebSocketRequest } from '@/types/WebSocketRequest';
import { WebSocketRequestType } from '@/types/WebSocketRequestType';
import { convertBase64ToUint8Array } from '@/utils/base64-utils';
import { decryptWithECDHKey } from '@/utils/ecdh-utils';
import { getWSHost } from '@/utils/utils';
import { tick } from 'svelte';
import { get } from 'svelte/store';

let ws: WebSocket | null = null;

export const connectToWebsocket = async () => {
  const wsHost = getWSHost();

  ws = new WebSocket(`${wsHost}/api/v1/ws`);

  // ws.onopen = () => {
  //   console.log('WebSocket connected');
  // };

  ws.onmessage = (event) => {
    const jsonData: WebSocketRequest = JSON.parse(event.data);
    switch (jsonData.type) {
      case WebSocketRequestType.ChatMessage:
        handleWebsocketChatMessage(jsonData.data);
        break;
    }
  };

  ws.onerror = (error) => {
    console.error('WebSocket error:', error);
  };

  // ws.onclose = () => {
  //   console.log('WebSocket connection closed');
  // };
}

export const sendWebsocketMessage = (requestData: WebSocketRequest) => {
  if (!ws || ws.readyState !== WebSocket.OPEN) return;

  ws.send(JSON.stringify(requestData));
}

const handleWebsocketChatMessage = async (message: ChatMessage) => {
  console.log('Received chat message:', message);
  
  const chatBodyElement = get(chatBodyStore);
  if (!chatBodyElement || !message.chatPartnerId) return;

  const messageContentBytes = convertBase64ToUint8Array(message.content)
  const messageContentIVBytes = convertBase64ToUint8Array(message.contentIV)
  const sharedSecret = getSharedSecret(message.chatPartnerId);
  if (!sharedSecret) {
    errorStore.set({ shown: true, error: ResponseErrorEnum.ErrInvalidSharedSecret })
    return;
  }

  message.content = await decryptWithECDHKey(messageContentBytes, messageContentIVBytes, sharedSecret)

  messageStore.update((prev) => [...prev, message])
  tick().then(() => {
    chatBodyElement.scrollTop = chatBodyElement.scrollHeight;
  });
}
