import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { MyProfileRequestsData } from "@/types/MyProfileRequestsData";
import type { MyProfileRequestsRequestResponse } from "@/types/MyProfileRequestsRequestResponse";
import type { NewSelectItem } from "@/types/NewSelectItem";
import { StatusEnum } from "@/types/StatusEnum";
import { navigate } from "svelte-routing";
import { SvelteURLSearchParams } from "svelte/reactivity";

export const fetchMyRequestsData = async (
  currentPage: number,
  setRequestsData: (data: MyProfileRequestsData) => void,
  selectStatusValue: StatusEnum | 0
) => {
  const { data, status } = await request<MyProfileRequestsRequestResponse>(
    "GET",
    `/my-profile/requests?page=${currentPage}&status=${selectStatusValue}`,
    undefined,
    true
  );

  if (status === 200) setRequestsData(data)
  else errorStore.set({ shown: true, error: data.error })
};


export const getInitialRequestsStatusValues = (): NewSelectItem[] => [
  { title: 'All' },
  { title: 'In Progress' },
  { title: 'Completed' },
  { title: 'Cancelled' },
  { title: 'Awaiting Acceptance' },
]

const convertMyProfileRequestsSelectValue = (
  value: string | string[]
): StatusEnum | 0 => {
  if (Array.isArray(value)) return 0;

  if (value === 'All') return 0;
  value = value.replace(/\s+/, '');

  const tuple = Object.entries(StatusEnum).find(([key]) => key === value);
  return tuple ? tuple[1] as StatusEnum : 0;
}

export const handleMyProfileRequestsSelectChange = (
  selectStatusValue: string | string[]
) => {
  const status = convertMyProfileRequestsSelectValue(selectStatusValue);
  navigate(`/my-profile/requests?page=1&status=${status}`)
}

export const getMyProfileRequestsPageParam = () => {
  const params = new SvelteURLSearchParams(window.location.search);
  return {
    page: parseInt(params.get('page') || '1'),
    status: parseInt(params.get('status') || '0'),
  };
}

export const getMyProfileRequestsInitialSelectIndex = (value: number): number => {
  if (value === 0) return 0;

  const tuple = Object.entries(StatusEnum).find(([_, val]) => val === value);
  const selectItems = getInitialRequestsStatusValues();

  const index = selectItems.findIndex(({ title }) => title.replaceAll(/\s+/g, '') === tuple?.[0]);
  console.log(index);
  return index;
  
}
