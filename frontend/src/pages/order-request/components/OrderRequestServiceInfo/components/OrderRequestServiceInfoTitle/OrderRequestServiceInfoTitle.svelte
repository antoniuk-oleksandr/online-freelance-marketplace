<script lang="ts">
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type OrderRequestServiceInfoTitleProps = {
        serviceData: RestrictedService,
    };

    const {serviceData}: OrderRequestServiceInfoTitleProps = $props();
</script>

<p class="font-semibold line-clamp-2">{serviceData.title}</p>