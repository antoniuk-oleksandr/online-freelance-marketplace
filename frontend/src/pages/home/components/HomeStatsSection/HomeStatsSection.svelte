<script lang="ts">
    import StatsSectionItemList
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItemList/HomeStatsSectionItemList.svelte";
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import type {KeyMetrics} from "@/types/KeyMetrics.ts";

    type StatsSectionProps = {
        keyMetrics: KeyMetrics,
    };

    const props: StatsSectionProps = $props();
</script>

<HomeSectionLayout>
    <HomeSectionHeader
            title="Key Metrics"
            description="Our Platform at a Glance"
    />
    <StatsSectionItemList {...props}/>
</HomeSectionLayout>