import type { TabType } from "@/types/TabType"
import { navigate } from "svelte-routing"
// @ts-ignore
import { orderByIdChatTabComponent } from './components/OrderByIdChatTab/OrderByIdChatTab.svelte'
// @ts-ignore
import { orderByIdDiaryTabComponent } from './components/OrderByIdDiaryTab/OrderByIdDiaryTab.svelte'
// @ts-ignore
import { orderByIdOverviewTabComponent } from './components/OrderByIdOverviewTab/OrderByIdOverviewTab.svelte'
// @ts-ignore
import { orderByIdRequirementsTabComponent } from './components/OrderByIdRequirementsTab/OrderByIdRequirementsTab.svelte'
// @ts-ignore
import { orderByIdDeliveryTabComponent } from './components/OrderByIdDeliveryTab/OrderByIdDeliveryTab.svelte'
// @ts-ignore
import { orderByIdReviewTabComponent } from './components/OrderByIdReviewTab/OrderByIdReviewTab.svelte'

import { request } from "@/api/request"
import { errorStore } from "@/common-stores/error-store"
import type { MyProfileOrderByIdData } from "@/types/MyProfileOrderByIdData"
import type { MyProfileOverviewRequestResponse } from "@/types/MyProfileOverviewRequestResponse"
import type { RouterBase } from "svelte-routing/types/RouterContext"
import { SvelteURLSearchParams } from "svelte/reactivity"
import { makeMyProfileChatRequest } from "./components/OrderByIdChatTab/helpers"
import { makeMyProfileDeliveryRequest } from "./components/OrderByIdDeliveryTab/helpers"
import { makeMyProfileDiaryRequest } from "./components/OrderByIdDiaryTab/helpers"
import { makeMyProfileRequirementsRequest } from "./components/OrderByIdRequirementsTab/helpers"
import type { Writable } from "svelte/store"
import { makeMyProfileReviewRequest } from "../../../requests/sub-pages/request-by-id/components/RequestByIdReviewTab/helpers"

export const getInitialTabComponentsData = (): MyProfileOrderByIdData => [
    undefined,
    undefined,
    undefined,
    undefined,
    undefined,
    undefined,
]

export const handleOrderByIdTabChange = (orderId: string, tabIndex: number) => {
    navigate(`/my-profile/orders/${orderId}?tabIndex=${tabIndex}`)
}

export const getOrderByIdTabData = (tabComponentsData: MyProfileOrderByIdData): TabType[] => [
    {
        title: 'Overview',
        //@ts-ignore
        component: orderByIdOverviewTabComponent,
        icon: 'hugeicons:globe-02'
    },
    {
        title: 'Requirements',
        //@ts-ignore
        component: orderByIdRequirementsTabComponent,
        icon: 'hugeicons:task-01'
    },
    {
        title: 'Chat',
        //@ts-ignore
        component: orderByIdChatTabComponent,
        icon: 'hugeicons:bubble-chat'
    },
    {
        title: 'Diary',
        //@ts-ignore
        component: orderByIdDiaryTabComponent,
        icon: 'hugeicons:book-02'
    },
    {
        title: 'Delivery',
        //@ts-ignore
        component: orderByIdDeliveryTabComponent,
        icon: 'hugeicons:package'
    },
    {
        title: 'Review',
        //@ts-ignore
        component: orderByIdReviewTabComponent,
        icon: 'hugeicons:star'
    },
]

export const makeMyProfileOverviewRequest = async (id: string): Promise<MyProfileOverviewRequestResponse> => {
    const url = `/my-profile/orders/${id}/overview`
    const response = await request<MyProfileOverviewRequestResponse>('GET', url, undefined, true);
    if (response.status !== 200) {
        errorStore.set({ shown: true, error: response.data.error });
    }

    return response;
}

export const fetchOrderTabData = (
    route: RouterBase,
    orderId: string,
    tabIndex: Writable<number>,
    tabComponentsData: MyProfileOrderByIdData,
    setTabComponentsData: (value: MyProfileOrderByIdData) => void,
) => {
    if (!route.path.includes('my-profile/orders/:orderId')) return

    const params = new SvelteURLSearchParams(window.location.search)
    const paramTabIndex = params.get('tabIndex') ? parseInt(params.get('tabIndex')!) : 0
    tabIndex.set(paramTabIndex)

    const requestFunctions = [
        makeMyProfileOverviewRequest,
        makeMyProfileRequirementsRequest,
        makeMyProfileChatRequest,
        makeMyProfileDiaryRequest,
        makeMyProfileDeliveryRequest,
        makeMyProfileReviewRequest,
    ]

    const fetchData = async () => {
        const requestFn = requestFunctions[paramTabIndex]
        if (!requestFn || tabComponentsData[paramTabIndex] !== undefined) return

        tabComponentsData[paramTabIndex] = await requestFn(orderId)
        setTabComponentsData(tabComponentsData)
    }

    fetchData()
}
