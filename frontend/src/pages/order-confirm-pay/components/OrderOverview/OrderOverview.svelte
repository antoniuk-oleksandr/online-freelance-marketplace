<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import Button from "@/common-components/Button/Button.svelte";
    import type {Service} from "@/types/Service.ts";
    import Spinner from "@/common-components/Spinner/Spinner.svelte";
    import type {Package} from "@/types/Package.ts";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import {round} from "@/utils/utils.ts";
    import OrderRequestSummaryItem
        from "@/pages/order-request/components/OrderRequestSummary/components/OrderRequestSummaryItem/OrderRequestSummaryItem.svelte";
    import OrderOverviewTopBlock
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlock/OrderOverviewTopBlock.svelte";
    import OrderOverviewBottomBlock
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewBottomBlock/OrderOverviewBottomBlock.svelte";

    type OrderOverviewProps = {
        serviceData?: Service,
        pkg?: Package,
    };

    const props: OrderOverviewProps = $props();
</script>

<PaperElement
        styles="flex flex-col {props.serviceData ? '!h-fit' : 'justify-center items-center !h-order-overview'} gap-6 !shadow-none md:!shadow-md !p-0 md:!p-6 !bg-transparent !ring-0 md:dark:!ring-1 md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
    {#if props.serviceData && props.pkg}
        <OrderOverviewTopBlock
                serviceData={props.serviceData}
                pkg={props.pkg}
        />
        <DividerElement/>
        <OrderOverviewBottomBlock pkg={props.pkg}/>
        <Button type="submit">Confirm & Pay</Button>
    {:else}
        <Spinner/>
    {/if}
</PaperElement>