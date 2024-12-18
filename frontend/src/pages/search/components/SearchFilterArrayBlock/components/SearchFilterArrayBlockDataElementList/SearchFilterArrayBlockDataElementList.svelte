<script lang="ts">
    import SearchFilterArrayBlockDataElementListLayout
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataElementList/SearchFilterArrayBlockDataElementListLayout.svelte";
    import SearchFilterArrayBlockDataElement
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataElement/SearchFilterArrayBlockDataElement.svelte";
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import type {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import type {SearchFilterArrayInputDataStore} from "@/types/SearchFilterArrayInputDataStore.ts";
    import {searchFilterArrayInputDataStore} from "@/pages/search/stores/search-filter-array-input-data-store.ts";
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
    import type {Language} from "@/types/Language.ts";
    import type {Skill} from "@/types/Skill.ts";
    import type {Category} from "@/types/Category.ts";
    import {getAllFilterArrayBlockDataElements} from "@/pages/search/helpers.ts";

    type SearchFilterArrayBlockDataElementListProps = {
        storeData: SearchFilterArrayStore,
        defaultFilterParams: GetFilterParamsRequestResponse["data"],
    }

    const {storeData, defaultFilterParams}: SearchFilterArrayBlockDataElementListProps = $props();

    let searchStoreData = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchStoreData = value);

    let inputValues = $state<SearchFilterArrayInputDataStore | undefined>();
    searchFilterArrayInputDataStore.subscribe((value) => inputValues = value);

    let allValues = $state<Language[] | Skill[] | Category[] | undefined>();
    $effect(() => {
        allValues = getAllFilterArrayBlockDataElements(defaultFilterParams, storeData, inputValues)
    })

    let searchValues = $state<string[] | undefined>();
    $effect(() => {
        searchValues = searchStoreData ? searchStoreData[storeData.attribute] : undefined
    });
</script>

{#if searchValues && allValues}
    <SearchFilterArrayBlockDataElementListLayout>
        {#each allValues as item}
            <SearchFilterArrayBlockDataElement
                    data={item}
                    attribute={storeData.attribute}
                    selected={searchValues.includes(item.id.toString())}
            />
        {/each}
    </SearchFilterArrayBlockDataElementListLayout>
{/if}
