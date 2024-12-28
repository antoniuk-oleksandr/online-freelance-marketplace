<script lang="ts">
    import NotFound from "@/common-components/NotFound/NotFound.svelte";
    import ServicePageLayout from "@/pages/services/ServicePageLayout.svelte";
    import ServiceBreadcrumbs from "@/pages/services/components/ServiceBreadcrumbs/ServiceBreadcrumbs.svelte";
    import ServiceContent from "@/pages/services/components/ServiceContent/ServiceContent.svelte";
    import {request} from "@/api/request.ts";
    import type {GetUserByIdRequestResponse} from "@/types/GetServiceByIdRequestResponse.ts";
    import type {UpdateFunc} from "@/types/UpdateFunc.ts";

    type ServicePageProps = { id: string; }

    const {id}: ServicePageProps = $props();

    let serviceResponse = $state<GetUserByIdRequestResponse | undefined>();
    let setServiceResponse: UpdateFunc<GetUserByIdRequestResponse | undefined> = (updater) => {
        (async () => {
            serviceResponse = await updater(serviceResponse);
        })();
    }

    request<GetUserByIdRequestResponse>(`/freelances/${id}`, "GET").then((response) => {
        serviceResponse = response;
    });
</script>

{#if serviceResponse && serviceResponse.status !== 200}
    <NotFound/>
{:else if serviceResponse && serviceResponse.status === 200}
    <ServicePageLayout>
        <ServiceBreadcrumbs service={serviceResponse.data.service}/>
        <ServiceContent
                serviceResponse={serviceResponse}
                setServiceResponse={setServiceResponse}
        />
    </ServicePageLayout>
{/if}
