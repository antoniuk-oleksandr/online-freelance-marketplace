<script lang="ts">
    import Stepper from "@/common-components/Stepper/Stepper.svelte";
    import type {Package} from "@/types/Package.ts";
    import OrderRequestPageLayout from "@/pages/order-request/OrderRequestPageLayout.svelte";
    import OrderRequestPageContent
        from "@/pages/order-request/components/OrderRequestPageContent/OrderRequestPageContent.svelte";
    import {fetchServiceDetailsAndPackage} from "@/pages/order-request/helpers.ts";
    import type {StepperItem} from "@/types/StepperItem.ts";
    import type {RestrictedService} from "@/types/RestrictedService.ts";

    type RequestPageParams = {
        serviceId: string;
    };

    const {serviceId}: RequestPageParams = $props();

    let serviceData = $state<RestrictedService | undefined>();
    const setServiceData = (newServiceData: RestrictedService) => serviceData = newServiceData;

    let selectedPackage = $state<Package | undefined>();
    const setSelectedPackage = (newPackage: Package) => selectedPackage = newPackage;

    fetchServiceDetailsAndPackage(serviceId, setServiceData, setSelectedPackage);


    const steps: StepperItem[] = [{text: "Order details",}, {text: "Confirm & pay"}, {text: "Submit requirements"}];
</script>

{#if serviceData && selectedPackage !== undefined}
    <OrderRequestPageLayout>
        <Stepper
                styles="col-span-2 capitalize flex h-fit justify-center"
                activeStepIndex={0} steps={steps}
        />
        <OrderRequestPageContent
                serviceData={serviceData}
                selectedPackage={selectedPackage}
                setSelectedPackage={setSelectedPackage}
        />
    </OrderRequestPageLayout>
{/if}
