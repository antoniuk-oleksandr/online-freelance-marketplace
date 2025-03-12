<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import OrderRequestSummaryItem
        from "@/pages/order-request/components/OrderRequestSummary/components/OrderRequestSummaryItem/OrderRequestSummaryItem.svelte";
    import Button from "@/common-components/Button/Button.svelte";
    import type {Package} from "@/types/Package.ts";
    import {round} from "@/utils/utils.ts";
    import {Link} from "svelte-routing";
    import Title from "@/common-components/Title/Title.svelte";

    type OrderRequestSummaryProps = {
        package: Package,
        serviceId: number,
    };

    const {package: pkg, serviceId}: OrderRequestSummaryProps = $props();

    let pkgData = $state({
        subtotal: 0,
        serviceFee: 0,
        total: 0,
    });

    const serviceFees = parseFloat(import.meta.env.VITE_SERVICE_FEES as string);

    $effect(() => {
        pkgData = {
            subtotal: pkg.price,
            serviceFee: pkg.price * serviceFees,
            total: pkg.price + pkg.price * serviceFees,
        }
    })
</script>

<PaperElement
        styles="flex flex-col gap-6 !h-fit !shadow-none md:!shadow-md !p-0 md:!p-6 !bg-transparent !ring-0 md:dark:!ring-1 md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
    <Title text="price summary"/>
    <OrderRequestSummaryItem
            title="subtotal"
            value="{round(pkgData.subtotal, 2)}$"
    />
    <OrderRequestSummaryItem
            title="service fee"
            value="{round(pkgData.serviceFee, 2)}$"
    />
    <DividerElement/>
    <OrderRequestSummaryItem
            styles="font-semibold"
            title="total"
            value="{round(pkgData.total, 2)}$"
    />
    <OrderRequestSummaryItem
            title="delivery time"
            value="{pkg.deliveryDays} {pkg.deliveryDays === 1 ? 'day' : 'days'}"
    />
    <Link to="/orders/confirm-pay?serviceId={serviceId}&packageId={pkg.id}">
        <Button styles="w-full">Continue to Checkout</Button>
    </Link>
</PaperElement>
