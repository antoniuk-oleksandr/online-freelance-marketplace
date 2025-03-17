import { errorStore } from "@/common-stores/error-store"
import type { MyProfileOrdersData } from "@/types/MyProfileOrdersData"
import type { MyProfileOrdersRequestResponse } from "@/types/MyProfileOrdersRequestResponse"
import { StatusEnum } from "@/types/StatusEnum"
import Cookies from "js-cookie"
import { SvelteURLSearchParams } from "svelte/reactivity"

export const makeMyProfileOrdersRequest = (currentPage: number): MyProfileOrdersRequestResponse => {
  return {
    data: {
      cursor: '',
      hasMore: true,
      totalPages: 100,
      orders: [
        {
          id: 12345,
          title: 'Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design Logo Design',
          status: StatusEnum.Completed,
          price: 50,
          date: 1742063416,
          image: 'http://localhost:8030/files/avatar_2.jpg',
        },
        {
          id: 12346,
          title: 'Website Development',
          status: StatusEnum.InProgress,
          price: 300,
          date: 1742063416,
          image: 'http://localhost:8030/files/avatar_2.jpg',
        },
        {
          id: 12347,
          title: 'SEO Optimization',
          status: StatusEnum.Pending,
          price: 120,
          date: 1742063416,
          image: 'http://localhost:8030/files/avatar_2.jpg',
        },
        {
          id: 12348,
          title: 'Mobile App UI/UX',
          status: StatusEnum.AwaitingAcceptance,
          price: 200,
          date: 1742063416,
          image: 'http://localhost:8030/files/avatar_2.jpg',
        },
      ],
    },
    status: 200,
  }
}

export const getMyProfileOrdersCurrentPage = () => {
  const params = new SvelteURLSearchParams(window.location.search);
  return parseInt(params.get('page') || '1');
}

export const fetchMyOrdersData = async (
  currentPage: number,
  setOrdersData: (data: MyProfileOrdersData) => void,
) => {
  const accessToken = Cookies.get('accessToken')

  const { data, status } = makeMyProfileOrdersRequest(currentPage)
  // const { data, status } = await request<MyProfileOrdersRequestResponse>(
  //   "/my-profile/orders", "GET", accessToken
  // );
  if (status === 200) setOrdersData(data)
  else errorStore.set({ shown: true, error: data.error })
}

export const buildMyProfileOrdersURL = (page: number, cursor: string | null | undefined) => {
  let url = "/my-profile/orders"
  url += `?page=${page}`
  if (cursor) url += `&cursor=${cursor}`
  return url
}

