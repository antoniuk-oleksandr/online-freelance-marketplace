<script lang="ts">
    import type {Service} from "@/types/Service.ts";

    type OrderRequestServiceInfoImageProps = {
        serviceData: Service,
    };

    const {serviceData}: OrderRequestServiceInfoImageProps = $props();
</script>

<img
        src={serviceData.images[0]}
        alt={serviceData.title}
        class="h-32 aspect-video object-cover"
/>