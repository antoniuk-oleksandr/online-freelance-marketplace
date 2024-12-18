<script lang="ts">
    import type {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
    import {searchFilterArrayInputDataStore} from "@/pages/search/stores/search-filter-array-input-data-store.ts";
    import {handleSearchFilterArrayInput} from "@/pages/search/handlers.ts";
    import type {SearchFilterArrayInputDataStore} from "@/types/SearchFilterArrayInputDataStore.ts";

    type SearchFilterArrayBlockDataInputProps = {
        title: string,
        attribute: SearchFilterArrayAttribute,
    }

    const {title, attribute}: SearchFilterArrayBlockDataInputProps = $props();
    let values = $state<SearchFilterArrayInputDataStore | undefined>();
    searchFilterArrayInputDataStore.subscribe((value) => values = value);
</script>

{#if values}
    <input
            value={values[attribute]}
            oninput={(e) => handleSearchFilterArrayInput(attribute, e)}
            placeholder="Filter by {title.toLowerCase()}"
            class="border-b placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary h-12 px-6 outline-none bg-transparent border-light-palette-divider dark:border-dark-palette-divider"
            type="text"
    />
{/if}
