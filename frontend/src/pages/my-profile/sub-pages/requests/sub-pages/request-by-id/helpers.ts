import type { TabType } from "@/types/TabType"
import { navigate } from "svelte-routing"
// @ts-ignore
import { orderByIdChatTabComponent } from '../../../orders/sub-pages/order-by-id/components/OrderByIdChatTab/OrderByIdChatTab.svelte'
// @ts-ignore
import { orderByIdDiaryTabComponent } from '../../../orders/sub-pages/order-by-id/components/OrderByIdDiaryTab/OrderByIdDiaryTab.svelte'
// @ts-ignore
import { orderByIdOverviewTabComponent } from '../../../orders/sub-pages/order-by-id/components/OrderByIdOverviewTab/OrderByIdOverviewTab.svelte'
// @ts-ignore
import { orderByIdRequirementsTabComponent } from '../../../orders/sub-pages/order-by-id/components/OrderByIdRequirementsTab/OrderByIdRequirementsTab.svelte'
// @ts-ignore
import { requestByIdDeliveryTabComponent } from './components/RequestByIdDeliveryTab/RequestByIdDeliveryTab.svelte'
// @ts-ignore
import { requestByIdReviewTabComponent } from './components/RequestByIdReviewTab/RequestByIdReviewTab.svelte'
import type { MyProfileOrderByIdData } from "@/types/MyProfileOrderByIdData"
import type { RouterBase } from "svelte-routing/types/RouterContext"
import { SvelteURLSearchParams } from "svelte/reactivity"
import { makeMyProfileChatRequest } from "../../../orders/sub-pages/order-by-id/components/OrderByIdChatTab/helpers"
import { makeMyProfileDeliveryRequest } from "../../../orders/sub-pages/order-by-id/components/OrderByIdDeliveryTab/helpers"
import { makeMyProfileDiaryRequest } from "../../../orders/sub-pages/order-by-id/components/OrderByIdDiaryTab/helpers"
import { makeMyProfileRequirementsRequest } from "../../../orders/sub-pages/order-by-id/components/OrderByIdRequirementsTab/helpers"
import { makeMyProfileOverviewRequest } from "../../../orders/sub-pages/order-by-id/helpers"
import type { OrderByIdSidebarStore } from "@/types/OrderByIdSidebarStore"
import { makeMyProfileReviewRequest } from "./components/RequestByIdReviewTab/helpers"

export const getInitialTabComponentsData = (): MyProfileOrderByIdData => [
    undefined,
    undefined,
    undefined,
    undefined,
    undefined,
    undefined,
]

export const handleOrderByIdTabChange = (orderId: string, tabIndex: number) => {
    navigate(`/my-profile/requests/${orderId}?tabIndex=${tabIndex}`)
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
        component: requestByIdDeliveryTabComponent,
        icon: 'hugeicons:package',
    },
    {
        title: 'Review',
        //@ts-ignore
        component: requestByIdReviewTabComponent,
        icon: 'hugeicons:star'
    }
]

export const fetchOrderTabData = (
    route: RouterBase,
    orderId: string,
    setTabIndex: (index: number) => void,
    tabComponentsData: MyProfileOrderByIdData,
    setTabComponentsData: (value: MyProfileOrderByIdData) => void,
) => {
    if (
        !route.path.includes('my-profile/orders/:orderId') &&
        !route.path.includes('my-profile/requests/:orderId')
    ) return

    const params = new SvelteURLSearchParams(window.location.search)
    const tabIndex = params.get('tabIndex') ? parseInt(params.get('tabIndex')!) : 0
    setTabIndex(tabIndex)

    const requestFunctions = [
        makeMyProfileOverviewRequest,
        makeMyProfileRequirementsRequest,
        makeMyProfileChatRequest,
        makeMyProfileDiaryRequest,
        makeMyProfileDeliveryRequest,
        makeMyProfileReviewRequest,
    ]

    const fetchData = async () => {
        const requestFn = requestFunctions[tabIndex]
        if (!requestFn || tabComponentsData[tabIndex] !== undefined) return

        tabComponentsData[tabIndex] = await requestFn(orderId)
        setTabComponentsData(tabComponentsData)
    }

    fetchData()
}

export const processOrderTabComponentsData = (
    tabComponentsData: MyProfileOrderByIdData,
    storeData: OrderByIdSidebarStore | undefined
): MyProfileOrderByIdData => {
    if (!storeData) return tabComponentsData;

    let orderData = tabComponentsData[0];
    if (!orderData || orderData.status !== 200) return tabComponentsData;

    if (storeData.endedAt !== undefined) {
        orderData.data.deliveryDate = storeData.endedAt;
    }
    orderData.data.status = storeData.status;

    tabComponentsData[0] = orderData;

    if (tabComponentsData[4] !== undefined && tabComponentsData[4].status === 200) {
        tabComponentsData[4] = undefined
    }

    return tabComponentsData;
}
