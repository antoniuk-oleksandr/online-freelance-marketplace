<script lang="ts">
    import {makeSearchDropdownTitle} from "@/pages/search/helpers.ts";
    import {OrderType} from "@/types/OrderType.ts";
    import Icon from "@iconify/svelte";

    type SearchDropdownTriggerIconProps = {
        selectedAdditionalItem: string,
    }

    const {selectedAdditionalItem}: SearchDropdownTriggerIconProps = $props();
</script>

{#if selectedAdditionalItem === makeSearchDropdownTitle(OrderType.Descending, true)}
    <Icon icon="hugeicons:sort-by-down-01" width="24" height="24"/>
{:else}
    <Icon icon="hugeicons:sort-by-up-01" width="24" height="24"/>
{/if}