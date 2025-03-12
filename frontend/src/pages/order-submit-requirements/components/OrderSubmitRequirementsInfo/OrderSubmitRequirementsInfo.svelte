<script lang="ts">
    import type {OrderSubmitRequirementsData} from "@/types/OrderSubmitRequirementsData.ts";
    import OrderSubmitRequirementsInfoLayout
        from "@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/OrderSubmitRequirementsInfoLayout.svelte";
    import OrderSubmitRequirementsInfoImage
        from "@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoImage/OrderSubmitRequirementsInfoImage.svelte";
    import OrderSubmitRequirementsInfoLineList
        from "@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoLineList/OrderSubmitRequirementsInfoLineList.svelte";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import OrderSubmitRequirementsInfoTitle
        from "@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoTitle/OrderSubmitRequirementsInfoTitle.svelte";

    type OrderSubmitRequirementsInfoProps = {
        orderData: OrderSubmitRequirementsData;
    }

    const {orderData}: OrderSubmitRequirementsInfoProps = $props();
</script>

<OrderSubmitRequirementsInfoLayout>
    <OrderSubmitRequirementsInfoImage
            title={orderData.service.title}
            image={orderData.service.image}
    />
    <div class="pt-6 flex flex-col gap-6">
        <OrderSubmitRequirementsInfoTitle title={orderData.service.title}/>
        <DividerElement/>
        <OrderSubmitRequirementsInfoLineList orderData={orderData}/>
    </div>
</OrderSubmitRequirementsInfoLayout>