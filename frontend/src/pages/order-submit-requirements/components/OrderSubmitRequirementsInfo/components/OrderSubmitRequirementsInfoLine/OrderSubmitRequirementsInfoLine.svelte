<script lang="ts">
    import OrderSubmitRequirementsInfoLineLayout
        from "@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/components/OrderSubmitRequirementsInfoLine/OrderSubmitRequirementsInfoLineLayout.svelte";

    type OrderSubmitRequirementsInfoLineProps = {
        key: string,
        value: string | number
    }

    const {key, value}: OrderSubmitRequirementsInfoLineProps = $props();
</script>

<OrderSubmitRequirementsInfoLineLayout>
    <span>{key}</span>
    <span>{value}</span>
</OrderSubmitRequirementsInfoLineLayout>