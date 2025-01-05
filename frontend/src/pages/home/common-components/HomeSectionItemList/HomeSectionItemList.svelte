<script lang="ts">
    import HomeSectionItemListLayout
        from "@/pages/home/common-components/HomeSectionItemList/HomeSectionItemListLayout.svelte";
    import HomeSectionItem from "@/pages/home/common-components/HomeSectionItem/HomeSectionItem.svelte";
    import type {HomeSectionItemType} from "@/types/HomeSectionItemType.ts";

    type HomeSectionItemListProps = {
        arr: HomeSectionItemType[]
    };

    const {arr} :HomeSectionItemListProps = $props();
</script>

<HomeSectionItemListLayout>
    {#each arr as item}
        <HomeSectionItem {...item}/>
    {/each}
</HomeSectionItemListLayout>