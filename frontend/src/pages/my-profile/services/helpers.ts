import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { MyProfileServicesData } from "@/types/MyProfileServicesData";
import type { MyProfileServicesRequestResponse } from "@/types/MyProfileServicesRequestResponse";

export const fetchMyServicesData = async (
  currentPage: number,
  setServicesData: (data: MyProfileServicesData) => void,
) => {
  const { data, status } = await request<MyProfileServicesRequestResponse>(
    "GET",
    `/my-profile/services?page=${currentPage}`,
    undefined,
    true
  );


  if (status === 200) setServicesData(data)
  else errorStore.set({ shown: true, error: data.error })
}
