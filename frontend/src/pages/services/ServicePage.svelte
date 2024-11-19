<script lang="ts">
    import {onMount} from "svelte";
    import type {Service} from "@/types/Service.ts";
    import {getServiceByIdRequest} from "@/api/get-service-by-id-request.ts";
    import {getFile} from "@/utils/utils.ts";
    import {tryToGetServiceById} from "@/pages/services/helpers.ts";
    import NotFound from "@/common-components/NotFound/NotFound.svelte";
    import ServicePageLayout from "@/pages/services/ServicePageLayout.svelte";
    import ServiceBreadcrumbs from "@/pages/services/components/ServiceBreadcrumbs/ServiceBreadcrumbs.svelte";
    import ServiceContent from "@/pages/services/components/ServiceContent/ServiceContent.svelte";

    type ServicePageProps = { id: string; }

    const {id}: ServicePageProps = $props();

    let service = $state<Service | null>(null);
    const setService = (value: Service) => service = value;

    let status = $state<number | null>(null);
    const setStatus = (value: number) => status = value;

    onMount(() => tryToGetServiceById(id, setService, setStatus))
</script>

{#if status === 404}
    <NotFound/>
{:else if service}
    <ServicePageLayout>
        <ServiceBreadcrumbs service={service}/>
        <ServiceContent {...service}/>
    </ServicePageLayout>
{/if}
