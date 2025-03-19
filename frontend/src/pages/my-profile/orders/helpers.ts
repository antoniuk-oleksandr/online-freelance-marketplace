import { request } from "@/api/request"
import { errorStore } from "@/common-stores/error-store"
import type { MyProfileOrdersData } from "@/types/MyProfileOrdersData"
import type { MyProfileOrdersRequestResponse } from "@/types/MyProfileOrdersRequestResponse"

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
