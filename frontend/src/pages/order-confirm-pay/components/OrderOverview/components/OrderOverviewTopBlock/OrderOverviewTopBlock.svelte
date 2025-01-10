<script lang="ts">
    import type {Service} from "@/types/Service.ts";
    import type {Package} from "@/types/Package.ts";
    import OrderOverviewTopBlockLayout
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlock/OrderOverviewTopBlockLayout.svelte";
    import OrderOverviewTopBlockPackage
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockPackage/OrderOverviewTopBlockPackage.svelte";
    import OrderOverviewTopBlockTitleImage
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockTitleImage/OrderOverviewTopBlockTitleImage.svelte";

    type OrderOverviewTopBlockProps = {
        serviceData: Service,
        pkg: Package,
    };

    const props: OrderOverviewTopBlockProps = $props();
</script>

<OrderOverviewTopBlockLayout>
    <OrderOverviewTopBlockTitleImage {...props}/>
    <OrderOverviewTopBlockPackage {...props}/>
</OrderOverviewTopBlockLayout>