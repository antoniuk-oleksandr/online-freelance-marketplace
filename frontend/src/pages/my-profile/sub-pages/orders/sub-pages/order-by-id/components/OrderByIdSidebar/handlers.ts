import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { PutUpdateOrderStatusByOrderIdRequestResponse } from "@/types/PutUpdateOrderStatusByOrderIdRequestResponse";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import { StatusEnum } from "@/types/StatusEnum"
import { orderByIdSidebarStore } from "../../stores/order-by-id-sidebar-store";

export const handleOrderStatusChange = async (
  value: string | string[],
  orderId: string
) => {
  if (Array.isArray(value)) return;

  const formattedValue = value.replaceAll(/\s/g, "");
  const entries = Object.entries(StatusEnum)
  const selectedStatus = entries.reduce<StatusEnum | undefined>((acc, [key, val]) => {
    if (key === formattedValue) {
      acc = val as StatusEnum;
    }
    return acc
  }, undefined);

  if (!selectedStatus) {
    errorStore.set({ shown: true, error: ResponseErrorEnum.UnexpectedError })
    return;
  }

  const url = `/my-profile/orders/${orderId}/status/${selectedStatus}`;
  console.log(url)
  const response = await request<PutUpdateOrderStatusByOrderIdRequestResponse>(
    'PUT', url, undefined, true
  );

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error });
    return;
  }

  orderByIdSidebarStore.update((prev) => {
    if (!prev) return {
      status: selectedStatus,
      endedAt: undefined,
    };
    return { ...prev, status: selectedStatus }
  })
}

