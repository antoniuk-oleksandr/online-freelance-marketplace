import type {TabType} from "@/types/TabType"
import {navigate} from "svelte-routing"
// @ts-ignore
import {orderByIdChatTabComponent} from './components/OrderByIdChatTab/OrderByIdChatTab.svelte'
// @ts-ignore
import {orderByIdDiaryTabComponent} from './components/OrderByIdDiaryTab/OrderByIdDiaryTab.svelte'
// @ts-ignore
import {orderByIdOverviewTabComponent} from './components/OrderByIdOverviewTab/OrderByIdOverviewTab.svelte'
// @ts-ignore
import {orderByIdRequirementsTabComponent} from './components/OrderByIdRequirementsTab/OrderByIdRequirementsTab.svelte'
// @ts-ignore
import {orderByIdDeliveryTabComponent} from './components/OrderByIdDeliveryTab/OrderByIdDeliveryTab.svelte'
import {StatusEnum} from "@/types/StatusEnum"
import {round} from "@/utils/utils"
import type {MyProfileOverviewRequestResponse} from "@/types/MyProfileOverviewRequestResponse"
import type {MyProfileOrderByIdData} from "@/types/MyProfileOrderByIdData"
import type {RouterBase} from "svelte-routing/types/RouterContext"
import {SvelteURLSearchParams} from "svelte/reactivity"
import {makeMyProfileChatRequest} from "./components/OrderByIdChatTab/helpers"
import {makeMyProfileDiaryRequest} from "./components/OrderByIdDiaryTab/helpers"
import {makeMyProfileRequirementsRequest} from "./components/OrderByIdRequirementsTab/helprers"
import {makeMyProfileDeliveryRequest} from "./components/OrderByIdDeliveryTab/helpers"

export const getInitialTabComponentsData = (): MyProfileOrderByIdData => [
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
]

export const makeMyProfileOverviewRequest = async (id: string): Promise<MyProfileOverviewRequestResponse> => {

    return {
        data: {
            id: 73,
            deliveryDate: 1741118830000,
            createdAt: 1740859630000,
            subtotal: 49.99,
            serviceFee: round(49.99 * 0.05, 2),
            totalPrice: round(49.99 + 49.99 * 0.05, 2),
            status: StatusEnum.InProgress,
            service: {
                image: 'http://localhost:8030/files/image_2.jpg',
                title: 'Professional Logo Design',
                package: {
                    description: 'Get a professional logo design design design design design design design design design design',
                    name: 'Premium',
                    deliveryTime: 3,
                },
            },
            freelancer: {
                username: 'name',
                id: 1,
                avatar: 'http://localhost:8030/files/avatar_2.jpg',
            }
        },
        status: 200
    }
}

export const fetchOrderTabData = (
    route: RouterBase,
    orderId: string,
    setTabIndex: (index: number) => void,
    tabComponentsData: MyProfileOrderByIdData,
    setTabComponentsData: (value: MyProfileOrderByIdData) => void,
) => {
    if (!route.path.includes('my-profile/orders/:orderId')) return

    const params = new SvelteURLSearchParams(window.location.search)
    const tabIndex = params.get('tabIndex') ? parseInt(params.get('tabIndex')!) : 0
    setTabIndex(tabIndex)

    const requestFunctions = [
        makeMyProfileOverviewRequest,
        makeMyProfileRequirementsRequest,
        makeMyProfileChatRequest,
        makeMyProfileDiaryRequest,
        makeMyProfileDeliveryRequest,
    ]

    const fetchData = async () => {
        const requestFn = requestFunctions[tabIndex]
        if (!requestFn || tabComponentsData[tabIndex] !== undefined) return

        tabComponentsData[tabIndex] = await requestFn(orderId)
        setTabComponentsData(tabComponentsData)
    }

    fetchData()
}
