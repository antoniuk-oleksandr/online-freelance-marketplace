<script lang="ts">
    import StatusElement from "@/common-components/StatusElement/StatusElement.svelte";
    import OrderSubmitRequirementsInfoStatusLineLayout
        from "@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoStatusLine/OrderSubmitRequirementsInfoStatusLineLayout.svelte";
    import type {OrderSubmitRequirementsData} from "@/types/OrderSubmitRequirementsData.ts";

    type OrderSubmitRequirementsInfoStatusLineProps = {
        orderData: OrderSubmitRequirementsData
    }

    const {orderData}: OrderSubmitRequirementsInfoStatusLineProps = $props();
</script>

<OrderSubmitRequirementsInfoStatusLineLayout>
    <span>Status</span>
    <StatusElement status={orderData.order.status}/>
</OrderSubmitRequirementsInfoStatusLineLayout>