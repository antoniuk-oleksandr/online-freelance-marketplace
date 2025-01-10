<script lang="ts">
    import BulletElement from "@/common-components/BulletElement/BulletElement.svelte";
    import OrderRequestPackagesBlockItemInfoLayout
        from "@/pages/order-request/components/OrderRequestPackagesBlock/components/OrderRequestPackagesBlockItemInfo/OrderRequestPackagesBlockItemInfoLayout.svelte";

    type OrderRequestPackagesBlockItemInfoProps = {
        price: number;
        deliveryDays: number;
    }

    const {price, deliveryDays}: OrderRequestPackagesBlockItemInfoProps = $props();
</script>

<OrderRequestPackagesBlockItemInfoLayout>
    <span>{price}$</span>
    <BulletElement/>
    <span>{deliveryDays} days</span>
</OrderRequestPackagesBlockItemInfoLayout>