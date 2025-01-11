<script lang="ts">
    import Rating from "@/common-components/Rating/Rating.svelte";
    import OrderRequestServiceInfoBottomLineLayout
        from "@/pages/order-request/components/OrderRequestServiceInfo/components/OrderRequestServiceInfoBottomLine/OrderRequestServiceInfoBottomLineLayout.svelte";
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type OrderRequestServiceInfoBottomLineProps = {
        serviceData: RestrictedService,
    };

    const {serviceData}: OrderRequestServiceInfoBottomLineProps = $props();
</script>

<OrderRequestServiceInfoBottomLineLayout>
    <Rating value={serviceData.rating}/>
    <span class="font-semibold ml-1">{serviceData.rating}</span>
    <span>({123} reviews)</span>
</OrderRequestServiceInfoBottomLineLayout>