<script lang="ts">
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
    import SearchFilterArrayBlock from "@/pages/search/components/SearchFilterArrayBlock/SearchFilterArrayBlock.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";

    type SearchFilterArrayBlockListProps = {
        defaultFilterParams: GetFilterParamsRequestResponse["data"],
        searchPageParams: SearchPageParams,
    }

    const {defaultFilterParams, searchPageParams}: SearchFilterArrayBlockListProps = $props();
</script>

{#if defaultFilterParams}
    <SearchFilterArrayBlock
            searchPageParams={searchPageParams}
            searchParamsAttribute={"skill"}
            data={defaultFilterParams.skill}
            title={"Skills"}
    />
    <SearchFilterArrayBlock
            searchPageParams={searchPageParams}
            searchParamsAttribute={"category"}
            data={defaultFilterParams.category}
            title={"Categories"}
    />
    <SearchFilterArrayBlock
            searchPageParams={searchPageParams}
            searchParamsAttribute={"language"}
            data={defaultFilterParams.language}
            title={"Languages"}
    />
{/if}
