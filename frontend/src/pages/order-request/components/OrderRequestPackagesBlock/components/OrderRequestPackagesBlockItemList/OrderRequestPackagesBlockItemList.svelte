<script lang="ts">
    import OrderRequestPackagesBlockItem
        from "@/pages/order-request/components/OrderRequestPackagesBlock/components/OrderRequestPackagesBlockItem/OrderRequestPackagesBlockItem.svelte";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import type {Service} from "@/types/Service.ts";
    import type {Package} from "@/types/Package.ts";

    type OrderRequestPackagesBlockItemListProps = {
        serviceData: Service,
        selectedPackage: Package,
        setSelectedPackage: (newPackage: Package) => void,
    };

    const props: OrderRequestPackagesBlockItemListProps = $props();
</script>

<div class="flex flex-col gap-y-3">
    {#each props.serviceData.packages as pkg, index}
        <OrderRequestPackagesBlockItem
                {...props}
                pkg={pkg}
        />
        {#if index !== props.serviceData.packages.length - 1}
            <DividerElement/>
        {/if}
    {/each}
</div>