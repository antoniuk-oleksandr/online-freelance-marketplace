<script lang="ts">
    import OrderOverviewTopBlockTitleImageLayout
        from "@/pages/order-confirm-pay/components/OrderOverview/components/OrderOverviewTopBlockTitleImage/OrderOverviewTopBlockTitleImageLayout.svelte";
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type OrderOverviewTopBlockTitleImageProps = {
        serviceData: RestrictedService,
    };

    const {serviceData}: OrderOverviewTopBlockTitleImageProps = $props();
</script>

<OrderOverviewTopBlockTitleImageLayout>
    <img
            class="aspect-video object-center object-cover rounded-sm w-32"
            alt="service"
            src={serviceData.image}
    />
    <p class="line-clamp-3">{serviceData.title}</p>
</OrderOverviewTopBlockTitleImageLayout>