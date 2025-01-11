<script lang="ts">
    import OrderRequestPackagesBlockItemList
        from "@/pages/order-request/components/OrderRequestPackagesBlock/components/OrderRequestPackagesBlockItemList/OrderRequestPackagesBlockItemList.svelte";
    import type {Package} from "@/types/Package.ts";
    import Title from "@/common-components/Title/Title.svelte";
    import OrderRequestPackagesBlockLayout
        from "@/pages/order-request/components/OrderRequestPackagesBlock/OrderRequestPackagesBlockLayout.svelte";
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type OrderRequestPackagesBlockProps = {
        serviceData: RestrictedService,
        selectedPackage: Package,
        setSelectedPackage: (newPackage: Package) => void,
    };

    const props: OrderRequestPackagesBlockProps = $props();
</script>

<OrderRequestPackagesBlockLayout>
    <Title text="Choose a package"/>
    <OrderRequestPackagesBlockItemList {...props}/>
</OrderRequestPackagesBlockLayout>
