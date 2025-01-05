<script>
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import CategoriesSectionList
        from "@/pages/home/components/HomeCategoriesSection/components/HomeCategoriesSectionList/HomeCategoriesSectionList.svelte";
</script>

<HomeSectionLayout>
    <HomeSectionHeader
            title="Popular Categories"
            description="Explore services across various categories"
    />
    <CategoriesSectionList/>
</HomeSectionLayout>
