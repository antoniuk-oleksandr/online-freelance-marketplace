<script lang="ts">
    import {createForm} from "felte";
    import {submitSearchPage} from "@/pages/search/helpers.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type SearchInputFormLayoutProps = LayoutProps & {
        searchPageParams: SearchPageParams,
    }

    const {searchPageParams, children}: SearchInputFormLayoutProps = $props();

    const {form} = createForm({
        onSubmit: () => submitSearchPage(searchPageParams)
    });
</script>

<form
        use:form
        class="relative col-span-2"
    >
    {@render children()}
</form>