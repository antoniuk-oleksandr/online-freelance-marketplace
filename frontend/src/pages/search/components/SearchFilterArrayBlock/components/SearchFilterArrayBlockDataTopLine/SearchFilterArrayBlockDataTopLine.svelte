<script lang="ts">
    import SearchFilterArrayBlockDataTopLineLayout
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataTopLine/SearchFilterArrayBlockDataTopLineLayout.svelte";
    import SearchFilterArrayBlockDataTopLineBackButton
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataTopLineBackButton/SearchFilterArrayBlockDataTopLineBackButton.svelte";
    import type {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";
    import SearchFilterArrayBlockDataTopLineResetButton
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataTopLineResetButton/SearchFilterArrayBlockDataTopLineResetButton.svelte";

    type SearchFilterArrayBlockDataTopLine = {
        storeData: SearchFilterArrayStore
    }

    const props: SearchFilterArrayBlockDataTopLine = $props();
</script>

<SearchFilterArrayBlockDataTopLineLayout>
    <SearchFilterArrayBlockDataTopLineBackButton/>
    <SearchFilterArrayBlockDataTopLineResetButton {...props}/>
</SearchFilterArrayBlockDataTopLineLayout>