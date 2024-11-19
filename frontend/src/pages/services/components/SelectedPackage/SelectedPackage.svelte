<script lang="ts">
    import type {Package} from "@/types/Package.ts";
    import SelectedPackageLayout from "@/pages/services/components/SelectedPackage/SelectedPackageLayout.svelte";
    import MyButton from "@/common-components/Button/Button.svelte";

    const {title, price, description, deliveryDays}: Package = $props();
    const primaryTextStyle = "text-light-palette-text-primary dark:text-dark-palette-text-primary";
</script>

<SelectedPackageLayout>
    <h2 class="${primaryTextStyle} text-xl font-semibold mb-4">{title}</h2>
    <p class="mb-4 text-sm">{description}</p>
    <div class="flex justify-between text-lg items-center ${primaryTextStyle}">
    <span>
        {deliveryDays} {deliveryDays > 1 ? "Days Delivery" : "Day Delivery"}
    </span>
        <p class="font-bold">{price}$</p>
    </div>
    <MyButton
            styles={"!mt-8 !w-full"}>
        Order Now
    </MyButton>
</SelectedPackageLayout>
