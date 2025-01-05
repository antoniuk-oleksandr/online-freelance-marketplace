<script>
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeFAQSectionItemList
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemList/HomeFAQSectionItemList.svelte";
</script>

<HomeSectionLayout fullHeight={true}>
    <HomeSectionHeader
            title="Frequently Asked Questions"
            description="Here are some common questions about our platform."
    />
    <HomeFAQSectionItemList/>
</HomeSectionLayout>