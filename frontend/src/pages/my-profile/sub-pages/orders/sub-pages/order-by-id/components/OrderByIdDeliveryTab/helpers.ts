import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { MyProfileDeliveryRequestResponse } from "@/types/MyProfileDeliveryRequestResponse";

export const makeMyProfileDeliveryRequest = async (orderId: string): Promise<MyProfileDeliveryRequestResponse> => {
  const url = `/my-profile/orders/${orderId}/delivery`;
  const response = await request<MyProfileDeliveryRequestResponse>('GET', url, undefined, true);

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error })
  }

  return response
}

export const downloadFile = async (file: string) => {
  const response = await fetch(file);
  const blob = await response.blob();

  const url = URL.createObjectURL(blob);
  const a = document.createElement('a');
  a.href = url;
  const fileName = file.split('/').pop();
  if (!fileName) return;

  a.download = fileName;
  document.body.appendChild(a);
  a.click();
  a.remove();
  URL.revokeObjectURL(url);
}

