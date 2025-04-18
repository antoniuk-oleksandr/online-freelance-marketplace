import { getAccessToken } from '@/utils/token-utils';
import { getHost, getWSHost } from '@/utils/utils';

let ws: WebSocket | null = null;

export const connectToWebsocket = async () => {
  const wsHost = getWSHost();
  const backendHost = getHost();

  let accessToken;
  const response = await getAccessToken(backendHost);
  if (response.status === 200) accessToken = response.data.accessToken;
  else throw new Error('Failed to get access token');

  ws = new WebSocket(`${wsHost}/api/v1/ws?token=${accessToken}`);

  ws.onopen = () => {
    console.log('WebSocket connected');
  };

  ws.onmessage = (event) => {
    console.log('WebSocket message:', event.data);
  };

  ws.onerror = (error) => {
    console.error('WebSocket error:', error);
  };

  ws.onclose = () => {
    console.log('WebSocket connection closed');
  };
}

export const sendWebsocketMessage = (message: string) => {
  if (ws) {
    ws.send(message);
  }
}
