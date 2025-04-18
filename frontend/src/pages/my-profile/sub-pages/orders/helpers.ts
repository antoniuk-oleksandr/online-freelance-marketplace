import { request } from "@/api/request"
import { errorStore } from "@/common-stores/error-store"
import type { MyProfileOrdersData } from "@/types/MyProfileOrdersData"
import type { MyProfileOrdersRequestResponse } from "@/types/MyProfileOrdersRequestResponse"
import type { OrderTableData } from "@/types/OrderTableData"
import { StatusEnum } from "@/types/StatusEnum"

export const fetchMyOrdersData = async (
  currentPage: number,
  setOrdersData: (data: MyProfileOrdersData) => void,
) => {
  const { data, status } = await request<MyProfileOrdersRequestResponse>(
    "GET",
    `/my-profile/orders?page=${currentPage}`,
    undefined,
    true
  );

  if (status === 200) setOrdersData(data)
  else errorStore.set({ shown: true, error: data.error })
}

export const defineOrderTableLink = (
  order: OrderTableData, isLargeScreen?: boolean
): string | undefined => {
  if (isLargeScreen === false) return undefined;

  if (order.status === StatusEnum.Incomplete) return `/orders/${order.id}/submit-requirements`
  else return `/my-profile/orders/${order.id}`
}

export const defineOrderTableButtonText = (order: OrderTableData) => {
  if (order.status === StatusEnum.Incomplete) return "Submit Requirements"
  else return "View Order"
}


export const formatDeliveryTime = (deliveryTime: number) => {
  if (deliveryTime === 1) return "1 day"
  else return `${deliveryTime} days`
}
