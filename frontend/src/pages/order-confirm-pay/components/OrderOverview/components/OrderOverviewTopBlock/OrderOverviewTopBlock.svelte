<script lang="ts">
    import type {Package} from "@/types/Package.ts";
    import OrderOverviewTopBlockLayout
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlock/OrderOverviewTopBlockLayout.svelte";
    import OrderOverviewTopBlockPackage
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockPackage/OrderOverviewTopBlockPackage.svelte";
    import OrderOverviewTopBlockTitleImage
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockTitleImage/OrderOverviewTopBlockTitleImage.svelte";
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type OrderOverviewTopBlockProps = {
        serviceData: RestrictedService,
        pkg: Package,
    };

    const props: OrderOverviewTopBlockProps = $props();
</script>

<OrderOverviewTopBlockLayout>
    <OrderOverviewTopBlockTitleImage {...props}/>
    <OrderOverviewTopBlockPackage {...props}/>
</OrderOverviewTopBlockLayout>