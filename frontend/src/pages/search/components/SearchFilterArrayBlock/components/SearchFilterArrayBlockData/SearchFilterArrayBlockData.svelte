<script lang="ts">
    import {searchFilterArrayStore} from "@/pages/search/stores/search-filter-array-store.ts";
    import SearchFilterArrayBlockDataLayout
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockData/SearchFilterArrayBlockDataLayout.svelte";
    import SearchFilterArrayBlockDataElementList
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataElementList/SearchFilterArrayBlockDataElementList.svelte";
    import type {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";
    import SearchFilterArrayBlockDataTopLine
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataTopLine/SearchFilterArrayBlockDataTopLine.svelte";
    import SearchFilterArrayBlockDataSelectButton
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataSelectButton/SearchFilterArrayBlockDataSelectButton.svelte";
    import SearchFilterArrayBlockDataInput
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataInput/SearchFilterArrayBlockDataInput.svelte";
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";

    type SearchFilterArrayBlockDataProps = {
        defaultFilterParams: GetFilterParamsRequestResponse["data"],
    }

    const {defaultFilterParams}: SearchFilterArrayBlockDataProps = $props();

    let searchFilterArrayStoreData = $state<SearchFilterArrayStore | undefined>();
    searchFilterArrayStore.subscribe((value) => searchFilterArrayStoreData = value)
</script>

{#if searchFilterArrayStoreData}
    <SearchFilterArrayBlockDataLayout>
        <SearchFilterArrayBlockDataTopLine
                storeData={searchFilterArrayStoreData}
        />
        <SearchFilterArrayBlockDataInput
                attribute={searchFilterArrayStoreData.attribute}
                title={searchFilterArrayStoreData.title}
        />
        <SearchFilterArrayBlockDataElementList
                defaultFilterParams={defaultFilterParams}
                storeData={searchFilterArrayStoreData}
        />
        <SearchFilterArrayBlockDataSelectButton/>
    </SearchFilterArrayBlockDataLayout>
{/if}
