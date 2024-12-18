<script lang="ts">
    import Button from "@/common-components/Button/Button.svelte";
    import SearchTopBarLayout from "@/pages/search/components/SearchTopBar/SearchTopBarLayout.svelte";
    import SearchDropdown from "@/pages/search/components/SearchDropdown/SearchDropdown.svelte";
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";
    import SearchInputForm from "@/pages/search/components/SearchInputForm/SearchInputForm.svelte";
    import SearchLeftSideTitle from "@/pages/search/components/SearchLeftSideTitle/SearchLeftSideTitle.svelte";

</script>

<SearchTopBarLayout>
    <SearchLeftSideTitle/>
    <SearchDropdown/>
    <SearchInputForm/>
    <Button
            clickAction={() => searchFilterDrawerStore.set(true)}
            styles="lg:!hidden w-fit row-start-1 col-start-2"
    >Filters
    </Button>
</SearchTopBarLayout>