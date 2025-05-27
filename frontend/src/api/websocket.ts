import { getHost, getWSHost } from '@/utils/utils';

let ws: WebSocket | null = null;

export const connectToWebsocket = async () => {
  const wsHost = getWSHost();
  const backendHost = getHost();

  // ws = new WebSocket(`${wsHost}/api/v1/ws`);

  // ws.onopen = () => {
  //   console.log('WebSocket connected');
  // };

  // ws.onmessage = (event) => {
  //   console.log('WebSocket message:', event.data);
  // };

  // ws.onerror = (error) => {
  //   console.error('WebSocket error:', error);
  // };

  // ws.onclose = () => {
  //   console.log('WebSocket connection closed');
  // };
}

export const sendWebsocketMessage = <T>(message: T) => {
  // if (!ws || ws.readyState !== WebSocket.OPEN) return;

  // ws.send(JSON.stringify(message));
}
