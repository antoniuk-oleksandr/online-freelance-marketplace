<script lang="ts">
    import SearchRightSideLayout from "@/pages/search/components/SearchRightSide/SearchRightSideLayout.svelte";
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
    import SearchRightSideContent
        from "@/pages/search/components/SearchRightSide/components/SearchRightSideContent/SearchRightSideContent.svelte";
    import SearchFilterArrayBlockData
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockData/SearchFilterArrayBlockData.svelte";

    type SearchRightSideProps = {
        isFiltersModalOpen: boolean,
        defaultFilterParams: GetFilterParamsRequestResponse["data"],
    }

    const props: SearchRightSideProps = $props();
</script>

<SearchRightSideLayout {...props}>
    <SearchRightSideContent {...props}/>
    <SearchFilterArrayBlockData {...props}/>
</SearchRightSideLayout>
