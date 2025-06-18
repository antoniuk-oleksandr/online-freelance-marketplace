import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import { orderByIdSidebarStore } from "@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/stores/order-by-id-sidebar-store";
import type { DeliveryData } from "@/types/DeliveryData";
import type { DeliveryFormData } from "@/types/DeliveryFormData"
import type { MyProfileDeliveryRequestResponse } from "@/types/MyProfileDeliveryRequestResponse";

export const handleDeliveryFormSubmit = async (
  data: DeliveryFormData,
  orderId: string,
  setNewDeliveryData: (data: DeliveryData) => void
) => {
  const formData = new FormData()
  formData.append("message", data.message)

  data.files.forEach(element => {
    formData.append("files", element);
  });

  const url = `/my-profile/orders/${orderId}/delivery`;
  const response = await request<MyProfileDeliveryRequestResponse>('POST', url, formData, true);
  if (response.status === 200 && response.data.delivery) {
    orderByIdSidebarStore.set({
      endedAt: response.data.delivery.date,
      status: response.data.status,
    })
    setNewDeliveryData(response.data)
  } else if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error })
  }
}
