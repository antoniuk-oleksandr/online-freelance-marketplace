<script lang="ts">
    import {capitalize} from "@/utils/utils.ts";
    import {getFormFromInputValue, setFormFromInputValue} from "@/pages/search/helpers.ts";
    import SearchFromToBlock from "@/pages/search/components/SearchFromToBlock/SearchFromToBlock.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";

    type SearchFromToBlockListProps = {
        searchPageParams: SearchPageParams
    }

    const {searchPageParams}: SearchFromToBlockListProps = $props();

    const data = ["rating", "deliveryTime", "level", "price"];

</script>

{#each data as item}
    <SearchFromToBlock
            title={capitalize(item)}
            from={getFormFromInputValue(searchPageParams, item, "From")}
            setFrom={(value) => setFormFromInputValue(value, item, "From")}
            to={getFormFromInputValue(searchPageParams, item, "To")}
            setTo={(value) => setFormFromInputValue(value, item, "To")}
    />
{/each}