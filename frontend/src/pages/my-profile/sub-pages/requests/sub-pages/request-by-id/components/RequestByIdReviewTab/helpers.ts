import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { MyProfileReviewRequestResponse } from "@/types/MyProfileReviewRequestResponse";

export const makeMyProfileReviewRequest = async (
  orderId: string
): Promise<MyProfileReviewRequestResponse> => {
  const url = `/my-profile/orders/${orderId}/review`;
  const response = await request<MyProfileReviewRequestResponse>('GET', url, undefined);
  if (response.status !== 200 && response.status !== 404) {
    errorStore.set({ shown: true, error: response.data.error })
  }
  
  return response;
}
