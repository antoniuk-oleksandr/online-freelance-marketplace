<script lang="ts">
    import type {Package} from "@/types/Package.ts";
    import SelectedPackageLayout from "@/pages/services/components/SelectedPackage/SelectedPackageLayout.svelte";
    import MyButton from "@/common-components/Button/Button.svelte";
    import {Link} from "svelte-routing";

    type SelectedPackageProps = Package & {
        serviceId: number
    }

    const {title, price, description, deliveryDays, id, serviceId}: SelectedPackageProps = $props();

    const primaryTextStyle = "text-light-palette-text-primary dark:text-dark-palette-text-primary";
</script>

<SelectedPackageLayout>
    <h2 class="${primaryTextStyle} text-xl font-semibold mb-4">{title}</h2>
    <p class="mb-4 text-sm text-light-palette-text-secondary dark:text-dark-palette-text-secondary">{description}</p>
    <div class="flex justify-between text-lg items-center ${primaryTextStyle}">
    <span>
        {deliveryDays} {deliveryDays > 1 ? "Days Delivery" : "Day Delivery"}
    </span>
        <p class="font-bold">{price}$</p>
    </div>
    <Link to="/orders/request?serviceId={serviceId}&packageId={id}">
        <MyButton
                styles={"!mt-6 !w-full"}>
            Request Order
        </MyButton>
    </Link>
</SelectedPackageLayout>
