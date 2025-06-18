import { modalStore } from "@/common-stores/modal-store";
import type { ReviewTabFormData } from "@/types/ReviewTabFormData";
import type { Writable } from "svelte/store";
//@ts-ignore
import { orderByIdReviewTabSuccessModelContent } from "./components/OrderByIdReviewTabSuccessModelContent/OrderByIdReviewTabSuccessModelContent.svelte";
import { request } from "@/api/request";
import type { PostOrderReviewRequestResponse } from "@/types/PostOrderReviewRequestResponse";
import { errorStore } from "@/common-stores/error-store";
import { orderReviewStore } from "../../stores/order-review-store";

export const handleReviewTabFormSubmit = async (
  data: ReviewTabFormData,
  orderId: string
) => {
  const url = `/my-profile/orders/${orderId}/review`;
  const response = await request<PostOrderReviewRequestResponse>('POST', url, data, true);
  if (response.status !== 201) {
    errorStore.set({ shown: true, error: response.data.error });
    return;
  }

  modalStore.set({ isOpened: true, renderContent: orderByIdReviewTabSuccessModelContent })
  orderReviewStore.set(response.data)
}

export const handleReviewRatingChange = (
  formDataStore: Writable<ReviewTabFormData>, value: number
) => {
  formDataStore.update((prev) => ({
    ...prev,
    rating: value,
  }))
}
