<script lang="ts">
    import ServiceContentLayout from "@/pages/services/components/ServiceContent/ServiceContentLayout.svelte";
    import ServiceBottomLeftSideContent
        from "@/pages/services/components/ServiceBottomLeftSideContent/ServiceBottomLeftSideContent.svelte";
    import ServiceTopLeftSideContent
        from "@/pages/services/components/ServiceTopLeftSideContent/ServiceTopLeftSideContent.svelte";
    import ServicePackagesBlock from "@/pages/services/components/ServicePackagesBlock/ServicePackagesBlock.svelte";
    import type {UpdateFunc} from "@/types/UpdateFunc.ts";
    import type {GetServiceByUserIdRequestResponse} from "@/types/GetServiceByIdRequestResponse.ts";

    type ServiceContentProps = {
        setServiceResponse: UpdateFunc<GetServiceByUserIdRequestResponse | undefined>,
        serviceResponse: GetServiceByUserIdRequestResponse
    };

    const serviceProps: ServiceContentProps = $props();
</script>

{#if serviceProps.serviceResponse.status === 200}
    <ServiceContentLayout>
        <ServiceTopLeftSideContent {...serviceProps.serviceResponse.data.service}/>
        <ServiceBottomLeftSideContent {...serviceProps}/>
        <ServicePackagesBlock {...serviceProps.serviceResponse.data.service}/>
    </ServiceContentLayout>
{/if}
