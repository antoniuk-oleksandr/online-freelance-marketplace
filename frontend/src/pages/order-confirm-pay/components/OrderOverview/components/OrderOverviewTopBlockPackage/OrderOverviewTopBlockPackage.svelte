<script lang="ts">
    import type {Package} from "@/types/Package.ts";
    import OrderOverviewTopBlockPackageLayout
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockPackage/OrderOverviewTopBlockPackageLayout.svelte";

    type OrderOverviewTopBlockPackageProps = {
        pkg: Package,
    };

    const {pkg}: OrderOverviewTopBlockPackageProps = $props();
</script>

<OrderOverviewTopBlockPackageLayout>
    <span class="line-clamp-1">{pkg.title}</span>
    <span>{pkg.price}$</span>
</OrderOverviewTopBlockPackageLayout>