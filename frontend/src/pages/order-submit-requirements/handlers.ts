import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { OrderSubmitRequirementsFormData } from "@/types/OrderSubmitRequirementsFormData";
import type { PostRequirementsByOrderIdRequestResponse } from "@/types/PostRequirementsByOrderIdRequestResponse";
import { navigate } from "svelte-routing";

export const handleOrderSubmitRequirementsFormSubmit = async (
  data: OrderSubmitRequirementsFormData
) => {
  const formData = new FormData();
  formData.append("answers", JSON.stringify(data.answers));
  formData.append("customerMessage", JSON.stringify(data.customerMessage));
  formData.append("orderId", data.orderId);
  data.files.forEach((file) => {
    formData.append("files", file);
  })

  const response = await request<PostRequirementsByOrderIdRequestResponse>("POST", `/orders/${data.orderId}/requirements`, formData, true)
  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error });
  } else navigate(`/my-profile/orders/${data.orderId}`);
}
