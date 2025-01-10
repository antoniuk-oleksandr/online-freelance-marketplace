<script lang="ts">
    import type {Service} from "@/types/Service.ts";

    type OrderRequestServiceInfoTitleProps = {
        serviceData: Service,
    };

    const {serviceData}: OrderRequestServiceInfoTitleProps = $props();
</script>

<p class="font-semibold line-clamp-2">{serviceData.title}</p>