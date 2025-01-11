<script lang="ts">
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type OrderRequestServiceInfoImageProps = {
        serviceData: RestrictedService,
    };

    const {serviceData}: OrderRequestServiceInfoImageProps = $props();
</script>

<img
        src={serviceData.image}
        alt={serviceData.title}
        class="h-32 aspect-video object-cover"
/>