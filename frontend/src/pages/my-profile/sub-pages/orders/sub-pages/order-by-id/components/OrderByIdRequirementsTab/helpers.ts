import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store"
import type { MyProfileRequirementsRequestResponse } from "@/types/MyProfileRequirementsRequestResponse"

export const makeMyProfileRequirementsRequest = async (
  orderId: string
): Promise<MyProfileRequirementsRequestResponse> => {
  const url = `/my-profile/orders/${orderId}/requirements`;
  const response = await request<MyProfileRequirementsRequestResponse>('GET', url, undefined, true);

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error });
    return response
  }

  return response;
}

