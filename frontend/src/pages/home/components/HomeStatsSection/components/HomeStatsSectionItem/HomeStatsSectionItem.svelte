<script lang="ts">
    import StatsSectionItemLayout
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItem/HomeStatsSectionItemLayout.svelte";
    import StatsSectionItemTopText
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItemTopText/HomeStatsSectionItemTopText.svelte";
    import StatsSectionItemCenterText
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItemCenterText/HomeStatsSectionItemCenterText.svelte";
    import StatsSectionItemIcon
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItemIcon/HomeStatsSectionItemIcon.svelte";

    type StatsSectionItemProps = {
        topText: string,
        centerText: string,
        icon: string,
    };

    const props: StatsSectionItemProps = $props();
</script>

<StatsSectionItemLayout>
    <StatsSectionItemTopText {...props}/>
    <StatsSectionItemCenterText {...props}/>
    <StatsSectionItemIcon {...props}/>
</StatsSectionItemLayout>