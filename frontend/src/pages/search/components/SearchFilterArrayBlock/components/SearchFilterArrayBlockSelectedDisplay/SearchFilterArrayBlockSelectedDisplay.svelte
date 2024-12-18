<script lang="ts">
    import type {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";

    type SearchFilterArrayBlockProps = {
        searchParamsAttribute: SearchFilterArrayAttribute,
        searchPageParams: SearchPageParams,
    }

    const props: SearchFilterArrayBlockProps = $props();

    let selectedNum = $state(0);

    $effect(() => {
        selectedNum = props.searchPageParams[props.searchParamsAttribute].length
    })
</script>

<span>{selectedNum === 0 ? "Any" : "Selected " + selectedNum}</span>
