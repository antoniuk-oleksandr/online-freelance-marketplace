<script lang="ts">
    import ServiceContentLayout from "@/pages/services/components/ServiceContent/ServiceContentLayout.svelte";
    import ServiceBottomLeftSideContent
        from "@/pages/services/components/ServiceBottomLeftSideContent/ServiceBottomLeftSideContent.svelte";
    import ServiceTopLeftSideContent
        from "@/pages/services/components/ServiceTopLeftSideContent/ServiceTopLeftSideContent.svelte";
    import ServicePackagesBlock from "@/pages/services/components/ServicePackagesBlock/ServicePackagesBlock.svelte";
    import type {Service} from "@/types/Service.ts";

    const serviceProps: Service = $props();
</script>

<ServiceContentLayout>
    <ServiceTopLeftSideContent {...serviceProps}/>
    <ServiceBottomLeftSideContent {...serviceProps}/>
    <ServicePackagesBlock {...serviceProps}/>
</ServiceContentLayout>