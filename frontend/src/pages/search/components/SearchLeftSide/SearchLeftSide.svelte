<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import SearchTopBar from "@/pages/search/components/SearchTopBar/SearchTopBar.svelte";
    import SearchServices from "@/pages/search/components/SearchServices/SearchServices.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import SearchLeftSideLayout from "@/pages/search/components/SearchLeftSide/SearchLeftSideLayout.svelte";

    type SearchLeftSideProps = {
        searchRequestResponse: any;
    }

    const {searchRequestResponse}: SearchLeftSideProps = $props();
</script>

<SearchLeftSideLayout searchRequestResponse={searchRequestResponse}>
    <PaperElement styles="p-3 flex flex-col w-full gap-y-3">
        <SearchTopBar/>
        <SearchServices
                searchRequestResponse={searchRequestResponse}
        />
    </PaperElement>
</SearchLeftSideLayout>
