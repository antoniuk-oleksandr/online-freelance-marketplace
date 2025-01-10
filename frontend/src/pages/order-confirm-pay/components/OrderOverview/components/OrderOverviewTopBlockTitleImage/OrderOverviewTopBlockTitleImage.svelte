<script lang="ts">
    import OrderOverviewTopBlockTitleImageLayout
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockTitleImage/OrderOverviewTopBlockTitleImageLayout.svelte";
    import type {Service} from "@/types/Service.ts";

    type OrderOverviewTopBlockTitleImageProps = {
        serviceData: Service,
    };

    const {serviceData}: OrderOverviewTopBlockTitleImageProps = $props();
</script>

<OrderOverviewTopBlockTitleImageLayout>
    <img
            class="aspect-video object-center object-cover rounded-sm w-32"
            alt="service"
            src={serviceData.images[0]}
    />
    <p class="line-clamp-3">{serviceData.title}</p>
</OrderOverviewTopBlockTitleImageLayout>