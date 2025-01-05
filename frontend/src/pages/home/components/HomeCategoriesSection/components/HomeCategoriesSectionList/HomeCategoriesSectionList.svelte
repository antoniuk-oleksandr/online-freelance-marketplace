<script lang="ts">
    import CategoriesSectionItem
        from "@/pages/home/components/HomeCategoriesSection/components/HomeCategoriesSectionItem/HomeCategoriesSectionItem.svelte";
    import CategoriesSectionListLayout
        from "@/pages/home/components/HomeCategoriesSection/components/HomeCategoriesSectionList/HomeCategoriesSectionListLayout.svelte";
    import {homeCategories, sliceHomeCategories} from "@/pages/home/helpers.ts";

    let categoryItems = $state(sliceHomeCategories(homeCategories));

    const handleResize = () =>
        categoryItems = sliceHomeCategories(homeCategories);

    $effect(() => {
        window.addEventListener("resize", handleResize)
        return () => window.removeEventListener("resize", handleResize)

    })
</script>

<CategoriesSectionListLayout>
    {#each categoryItems as categoryItem}
        <CategoriesSectionItem {...categoryItem}/>
    {/each}
</CategoriesSectionListLayout>
