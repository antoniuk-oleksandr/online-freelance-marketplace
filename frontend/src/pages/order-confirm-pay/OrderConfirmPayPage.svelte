<script lang="ts">
    import Stepper from "@/common-components/Stepper/Stepper.svelte";
    import OrderConfirmPayPageLayout from "@/pages/order-confirm-pay/OrderConfirmPayPageLayout.svelte";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import PaymentForm from "@/pages/order-confirm-pay/components/PaymentForm/PaymentForm.svelte";
    import OrderOverview from "@/pages/order-confirm-pay/components/OrderOverview/OrderOverview.svelte";
    import {fetchServiceDetailsAndPackage} from "@/pages/order-request/helpers.ts";
    import type {Package} from "@/types/Package.ts";
    import type {Service} from "@/types/Service.ts";
    import type {StepperItem} from "@/types/StepperItem.ts";

    type OrderConfirmPayPageParams = {
        serviceId: string,
    };

    const {serviceId}: OrderConfirmPayPageParams = $props();

    let selectedPackage = $state<Package | undefined>();
    const setSelectedPackage = (newPackage: Package) => selectedPackage = newPackage;

    let serviceData = $state<Service | undefined>();
    const setServiceData = (newServiceData: Service) => serviceData = newServiceData;

    fetchServiceDetailsAndPackage(serviceId, setServiceData, setSelectedPackage);

    let steps: StepperItem[] = $derived(
        [{
            text: "Order details",
            link: `/order/request/${serviceId}?packageId=${selectedPackage?.id}`,
        }, {text: "Confirm & pay"}, {text: "Submit requirements"}]
    );
</script>

<OrderConfirmPayPageLayout>
    <Stepper styles="col-span-2 capitalize flex h-fit justify-center" activeStepIndex={1} steps={steps}/>
    <div class="flex flex-col lg:grid grid-cols-search-page gap-6 h-full">
        <PaperElement
                styles="flex flex-col gap-6  !shadow-none md:!shadow-md !p-0 md:!p-6 !bg-transparent !ring-0 md:dark:!ring-1 md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
        >
            <PaymentForm/>
        </PaperElement>
        <OrderOverview
                pkg={selectedPackage}
                serviceData={serviceData}
        />
    </div>
</OrderConfirmPayPageLayout>