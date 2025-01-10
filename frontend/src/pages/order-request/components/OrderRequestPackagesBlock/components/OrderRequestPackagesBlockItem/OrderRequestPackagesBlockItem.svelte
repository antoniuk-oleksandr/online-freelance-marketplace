<script lang="ts">
    import RadioButton from "@/common-components/RadioButton/RadioButton.svelte";
    import type {Package} from "@/types/Package.ts";
    import OrderRequestPackagesBlockItemLayout
        from "@/pages/order-request/components/OrderRequestPackagesBlock/components/OrderRequestPackagesBlockItem/OrderRequestPackagesBlockItemLayout.svelte";
    import OrderRequestPackagesBlockItemTitle
        from "@/pages/order-request/components/OrderRequestPackagesBlock/components/OrderRequestPackagesBlockItemTitle/OrderRequestPackagesBlockItemTitle.svelte";
    import OrderRequestPackagesBlockItemInfo
        from "@/pages/order-request/components/OrderRequestPackagesBlock/components/OrderRequestPackagesBlockItemInfo/OrderRequestPackagesBlockItemInfo.svelte";

    type OrderRequestPackagesBlockProps = {
        pkg: Package,
        selectedPackage: Package,
        setSelectedPackage: (newPackage: Package) => void,
    };

    const {pkg, selectedPackage, setSelectedPackage}: OrderRequestPackagesBlockProps = $props();
</script>

<OrderRequestPackagesBlockItemLayout
        package={pkg}
        setSelectedPackage={setSelectedPackage}
>
    <RadioButton
            setValue={() => setSelectedPackage(pkg)}
            value={pkg.id === selectedPackage.id}
    />
    <div>
        <OrderRequestPackagesBlockItemTitle title={pkg.title}/>
        <OrderRequestPackagesBlockItemInfo
                price={pkg.price}
                deliveryDays={pkg.deliveryDays}
        />
    </div>
</OrderRequestPackagesBlockItemLayout>