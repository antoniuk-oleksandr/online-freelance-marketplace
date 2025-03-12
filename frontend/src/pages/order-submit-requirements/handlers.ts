import { request } from "@/api/request";
import type { OrderSubmitRequirementsFormData } from "@/types/OrderSubmitRequirementsFormData";
import Cookies from "js-cookie";

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

  const token = Cookies.get("accessToken");
  const response = await request(`/orders/${data.orderId}/requirements`, "POST", token, formData)
  console.log(response);
}
